package main

import (
	"learnhub-toolkit/api"
	"learnhub-toolkit/database"
	"net/http"
)

func main() {
	database.Migrate()
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}

func setupRoutes() {
	http.HandleFunc("/api/rooms/create", api.CreateRoom)
	http.HandleFunc("/api/rooms/delete/{id}", api.DeleteRoom)
	http.HandleFunc("/api/rooms/all", api.AllRooms)
	http.HandleFunc("/api/rooms/join/{id}", api.JoinRoom)
}
