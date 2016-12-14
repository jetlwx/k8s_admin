package main

import (
	"flag"
	"log"
	"net/url"
	 "time"
	"github.com/gorilla/websocket"
    
)

func main(){
    sendmsg:="Hello"
 var C1=make(chan bool)
 var msg=make(chan string)
var addr = flag.String("addr", "localhost:10009", "http service address")
var u = url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
    flag.Parse()
    log.SetFlags(0)
    c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

 if err !=nil{
     log.Println("err for Open -> ",err)
 }
 log.Println("Open Ok!")


// read
go func(){
    defer c.Close()
    for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
             m:=string(message)
			 if m == "Over" {
                  msg <- m
			 	C1 <-true 
			 }else{
                 if m !="" {
                     msg <- m
                 }
             }
		}
        
       
}()


//write
go func(){
     err := c.WriteMessage(websocket.TextMessage, []byte(string(sendmsg)))
          		if err != nil {
				log.Println("write:", err)
				return
			}

}()

//call fun
L:
for {
select {
     case t,_:=<-msg:
     log.Println("recv:",t)
     case v,_:= <-C1:
      log.Println("C1",v)
     break L
    case <-time.After(3*time.Second):
   log.Println("timout,exit 1")
   break L
}
}
}