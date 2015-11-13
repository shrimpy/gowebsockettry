package main

import (
	"code.google.com/p/go.net/websocket"
	"io"
	"net/http"
	"os"
)

func echoHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func main() {
	port := "8080"
    if os.Getenv("HTTP_PLATFORM_PORT") != "" {
        port = os.Getenv("HTTP_PLATFORM_PORT")
    }
	
	http.Handle("/echo", websocket.Handler(echoHandler))
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
