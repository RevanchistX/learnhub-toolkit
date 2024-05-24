package server

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/kbinani/screenshot"
	"image"
	"image/png"
	"net"
	"net/http"
)

func AppServer() {
	setupRoutes()
	ip := GetOutboundIP()
	fmt.Println(ip)
	http.ListenAndServe(":8080", nil)
}

func GetOutboundIP() net.IP {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP
			}
		}
	}
	return nil
}

func setupRoutes() {
	http.HandleFunc("/streaming", streaming)
	http.HandleFunc("/wasm", wasm)
}

func wasm(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	http.ServeFile(writer, request, "./assets/simple.wasm")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func streaming(writer http.ResponseWriter, request *http.Request) {
	ws, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
	}
	for true {
		wsErr := ws.WriteMessage(websocket.BinaryMessage, encodeImage())
		if wsErr != nil {
			fmt.Println("am i throwing errors")
			fmt.Println(err)
		}
	}
}

func generateImage(index int) image.Image {
	bounds := screenshot.GetDisplayBounds(index)
	capturedImg, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	return image.Image(capturedImg)
}

func encodeImage() []byte {
	buffer := new(bytes.Buffer)
	png.Encode(buffer, generateImage(1))
	return buffer.Bytes()
}
