package websocket_learn

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("client")
	// dl := websocket.Dialer{}
	// conn, _, err := dl.Dial("ws://127.0.0.1:8888/aaa", nil)

	u := url.URL{Scheme: "ws", Host: "localhost:8888", Path: "/aaa"}
	log.Printf("connecting to %s.", u.String())
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Println(err)
		return
	}
	conn.WriteMessage(websocket.TextMessage, []byte("我来了"))

	go send(conn)

	for {
		m, p, e := conn.ReadMessage()
		if e != nil {
			break
		}
		fmt.Println(m, string(p))
	}
}

func send(conn *websocket.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		l, _, _ := reader.ReadLine()
		conn.WriteMessage(websocket.TextMessage, l)
	}
}
