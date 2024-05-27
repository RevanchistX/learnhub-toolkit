package main

import "fmt"

func main() {
	wait := make(chan struct{})
	<-wait
}

//export add
func add(x, y int) int {
	fmt.Print("im adding two numbers")
	return x + y
}
