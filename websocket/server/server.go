package websocket_server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var Up = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := Up.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(messageType, string(p))
	}
	defer conn.Close()
	log.Println("服务关闭")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}
