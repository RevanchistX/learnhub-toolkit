package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func main() {
	address := ":8080"
	prefix := "/"
	root := "./wasm/"

	var err error
	root, err = filepath.Abs(root)

	currentTime := time.Now()
	fmt.Println("Share screen is starting", currentTime.Format("Mon 02 Jan 2006 03:04pm"))
	log.Printf("serving %s as %s on %s", root, prefix, address)
	http.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir(root))))

	//routes
	http.HandleFunc("/start-sharing", takeScreenshot)
	http.HandleFunc("/shared-screen", fetchScreenshot)
	http.HandleFunc("/stop-sharing", stopSharing)
	http.HandleFunc("/fetch-png", fetchPNG)

	mux := http.DefaultServeMux.ServeHTTP
	logger := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr + " " + r.Method + " " + r.URL.String())
	})
	err = http.ListenAndServe(address, logger)
	if err != nil {
		log.Fatalln(err)
	}

	gracefulTerminateSystem()
}

func takeScreenshot(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Sharing screen Started")
	ticker := time.NewTicker(500)
	defer ticker.Stop()
	buffer := new(bytes.Buffer)
	//TODO finish watching video  https://www.youtube.com/watch?v=aKFD5UmdzQQ&t=6m5s
	//screenRegion := image.Rect(0, 0, 800, 600)  - portion of screen

}
func fetchScreenshot(w http.ResponseWriter, r *http.Request) {}
func stopSharing(w http.ResponseWriter, r *http.Request)     {}
func fetchPNG(w http.ResponseWriter, r *http.Request)        {}

func gracefulTerminateSystem() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		interrupt <- true
		fmt.Println("Ctrl+C was pressed in terminal")
		os.Exit(0)
	}()
}
