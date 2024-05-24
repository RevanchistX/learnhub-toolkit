package socket

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/kbinani/screenshot"
	"image"
	"image/png"
)

func Reader(conn *websocket.Conn) {
	fmt.Println("here i am")
	err := conn.WriteMessage(websocket.BinaryMessage, encodeImage())
	if err != nil {
		fmt.Println(err)
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
	png.Encode(buffer, generateImage(0))
	return buffer.Bytes()
}
