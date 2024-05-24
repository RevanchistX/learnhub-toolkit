package main

import (
	"lh-toolkit/internal/client"
	"lh-toolkit/internal/server"
	"os"
)

func main() {
	args := os.Args
	isServer := args[1]
	if isServer == "server" {
		server.AppServer()
	} else {
		client.AppClient()
	}
}
