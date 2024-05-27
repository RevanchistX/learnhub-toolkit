package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/kbinani/screenshot"
	"image"
	"image/png"
	"learnhub-toolkit/database/models"
	"learnhub-toolkit/database/services"
	"net/http"
	"strconv"
)

func CreateRoom(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")

	_, room := new(services.Rooms).Create(new(models.Room))
	roomJson, err := json.Marshal(room)
	if err != nil {
		fmt.Print("cannot generate room json")
	}
	writer.Write(roomJson)
}

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

func DeleteRoom(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")
	roomId, _ := strconv.Atoi(request.URL.Path[len("/api/rooms/delete/"):])
	_, room := new(services.Rooms).Delete(roomId)
	roomJson, err := json.Marshal(room)
	if err != nil {
		fmt.Print("cannot generate room json")
	}
	writer.Write(roomJson)
}

func AllRooms(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")

	_, room := new(services.Rooms).All(new(models.Room))
	roomJson, err := json.Marshal(room)
	if err != nil {
		fmt.Print("cannot generate room json")
	}
	writer.Write(roomJson)
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
