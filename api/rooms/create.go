package rooms

import (
	"encoding/json"
	"fmt"
	"learnhub-toolkit/database/models"
	"learnhub-toolkit/database/services"
	"net/http"
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
