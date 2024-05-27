package rooms

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/kbinani/screenshot"
	"image"
	"image/png"
	"net/http"
)

func JoinRoom(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	//UPDATE logic when you figure out client side
	//roomId, _ := strconv.Atoi(request.URL.Path[len("/api/rooms/join/"):])
	ws, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
	}
	for true {
		wsErr := ws.WriteMessage(websocket.BinaryMessage, encodeImage())
		if wsErr != nil {
			fmt.Println(err)
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
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
