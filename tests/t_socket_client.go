package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:10009", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)
log.Println("ding ...... 1")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
log.Println("ding ...... 2")
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Println("ding ...... 3")
	log.Printf("connecting to %s", u.String())


	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	log.Println("ding ...... 4")
	if err != nil {
		log.Fatal("dial:", err)
	}
	log.Println("ding ...... 5")
	defer c.Close()
log.Println("ding ...... 6")
	done := make(chan struct{ })
	over :=make(chan bool)

	go func() {
		log.Println("ding ...... 7")
		defer c.Close()
		defer close(done)
		for {
			log.Println("ding ...... 8")
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			if string(message) == "Over" {
				//close(done)
				over <- true
			}
			log.Printf("recv: %s", message)
		}
	}()
log.Println("ding ...... 9")
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
log.Println("ding ...... 10")
Out:
	for {
		log.Println("ding ...... 11")
		select {
			case <-over:
            break Out
		case t := <-ticker.C:
		log.Println("ding ...... 12")
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
		log.Println("ding ...... 13")
			log.Println("interrupt")
			// To cleanly close a connection, a client should send a close
			// frame and wait for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			log.Println("ding ...... 15")
			case <-time.After(time.Second):
			log.Println("ding ...... 16")
			}
			log.Println("ding ...... 17")
			c.Close()
			log.Println("ding ...... 18")
			return
		}
	}
}

// func Write(c *websocket.Conn) {

	
// }