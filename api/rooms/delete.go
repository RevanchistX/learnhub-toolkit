package rooms

import (
	"encoding/json"
	"fmt"
	"learnhub-toolkit/database/services"
	"net/http"
	"strconv"
)

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
