package main

import (
	"learnhub-toolkit/api/rooms"
	"learnhub-toolkit/database"
	"net/http"
)

func main() {
	database.Migrate()
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}

func setupRoutes() {
	http.HandleFunc("/api/rooms/create", rooms.CreateRoom)
	http.HandleFunc("/api/rooms/delete/{id}", rooms.DeleteRoom)
	http.HandleFunc("/api/rooms/all", rooms.AllRooms)
	http.HandleFunc("/api/rooms/join/{id}", rooms.JoinRoom)
}
