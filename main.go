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
	wait := make(chan struct{})
	<-wait
}

//func main() {
//	wait := make(chan struct{})
//	<-wait
//}
//
////export add
//func add(x, y int) int {
//	fmt.Print("im adding two numbers")
//	return x + y
//}
