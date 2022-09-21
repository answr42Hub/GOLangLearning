package main

import (
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/", gererConnection)
	http.HandleFunc("/ws", ws)

	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	http.ListenAndServe(":8080", nil)
}

func gererConnection(w http.ResponseWriter, r *http.Request) {

	vueRaw, _ := os.ReadFile("views/index.html")
	vue := string(vueRaw)
	vue = strings.Replace(vue, "###TITRE###", "Go Funk yourself yaooooooo", 1)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, vue)
}

func ws(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
	}
	defer conn.Close()
	for {
		mt, msg, _ := conn.ReadMessage()
		log.Printf("recv: %s", msg)
		conn.WriteMessage(mt, msg)
	}
}
