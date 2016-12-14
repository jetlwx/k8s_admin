package main

import (
	"flag"
	"log"
	"net/url"
	// "os"
	// "os/signal"
	 "time"
// "strconv"
	"github.com/gorilla/websocket"
    
)


var Addr = flag.String("addr", "localhost:10009", "http service address")
var U = url.URL{Scheme: "ws", Host: *Addr, Path: "/echo"}
 var C1=make(chan bool)
 var msg=make(chan string)

func Open() (*websocket.Conn,error){
    c, _, err := websocket.DefaultDialer.Dial(U.String(), nil)
    flag.Parse()
    log.SetFlags(0)
    return c, err
}


func Write(c *websocket.Conn,t2 string) {
              err := c.WriteMessage(websocket.TextMessage, []byte(string(t2)))
          		if err != nil {
				log.Println("write:", err)
				return
			}
               err1 := c.WriteMessage(websocket.TextMessage, []byte("Over"))
          		if err1 != nil {
				log.Println("write:", err1)
				return
			}
}
func Read(c *websocket.Conn,C1 chan bool,msg chan string)  {
    for{
     _, message, err2 := c.ReadMessage()
			if err2 != nil {
				log.Println("read:", err2)
				return
			}
            m:=string(message)
            log.Println("msg:",string(message))
			 if m == "OverOk" {
			 	break
			 }else{
                 if m !="" {
                     msg <- m
                 }
             }
    }
    C1 <-true  
   
}

func Check(c *websocket.Conn,logfilepath string){
		file, _ := os.OpenFile(logfilepath, os.O_APPEND, 0666)
		defer file.Close()
		defer c.Close()
	L:
		for {
			select {
			case t, _ := <-msg:
				log.Println("recv1:", t)
				file.WriteString(t)

				if t == "OverOk" {
					break L
				}
			//	data = data + t
			case v, _ := <-C1:
				log.Println("C1", v)
				break L
			case <-time.After(600 * time.Second):
				log.Println("timout,exit 1")
				break L
			}
		}
}

func main() {  

c,err:=Open()
defer c.Close()
 if err !=nil{
     log.Println("err for Open -> ",err)
 }
 log.Println("Open Ok!")

go Write(c,"ttttttttttttttt")
go Read(c,C1,msg)

logfilepath, _ := CreateLogFile()
go Check(c,logfilepath)
	// for {
	// 	select {
	// 	case t, _ := <-msg:
	// 		log.Println("recv1:", t)
	// 		if t == "OverOk" {
	// 			break L
	// 		}

		
	// 	case v, _ := <-C1:
	// 		log.Println("C1", v)
	// 		break L
	// 	case <-time.After(600 * time.Second):
	// 		log.Println("timout,exit 1")
	// 		break L
	// 	}
	// }
	c.Close()

}