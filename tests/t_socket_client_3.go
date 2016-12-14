package virtualization

import (
	"encoding/json"
	"flag"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/jetlwx/k8s_admin/common"
)

type VtController struct {
	beego.Controller
}

// type is interface:means libvirtd interface , command :means the linux system command
type SendMsg struct {
	Type   string
	Conten string
}

var Addr = flag.String("addr", "localhost:10009", "http service address")

// make the msg to json
func MakeMsgToJson(mtype string, msg ...string) ([]byte, error) {
	m := &SendMsg{}
	m.Type = mtype
	for _, v := range msg {
		m.Conten = m.Conten + " " + v
	}
	js, err := json.Marshal(m)
	return js, err
}

//return recved command log temp log filename and file path
func CreateLogFile() (filename string, filepath string) {
	filename = common.RandomString()
	fpath := beego.AppConfig.String("beegologdir") + filename
	return filename, fpath
}
func (w *VtController) Get() {
	w.TplName = "test/t1.html"
	var C1 = make(chan bool)
	var msg = make(chan string)
	c, err := Open()
	if err != nil {
		log.Println("err for Open -> ", err)
	}
	log.Println("Open Ok!")

	//go Write(c, "/root/scripts/auto_create_vm.sh  testvm 72.16.6.135 255.255.255.0 172.16.6.1 8.8.8.8")
	go Read(c, C1, msg)

	//data := ""
	filename, filepath := CreateLogFile()
	log.Println("filepath", filepath, "filename", filename)
	go func() {
	L:
		for {
			select {
			case t, _ := <-msg:
				log.Println("recv1:", t)
				file, err1 := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
				if err1 != nil {
					w.Ctx.WriteString(common.CustomerErr(err))
				}
				_, err := file.WriteString(t + "<br>")
				if err != nil {
					log.Println("eeeee--->", err)
				}

				if t == "OverOk" {
					file.Close()
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

	}()
	w.Data["Filename"] = "/logs/" + filename

}

func Open() (*websocket.Conn, error) {
	var U = url.URL{Scheme: "ws", Host: *Addr, Path: "/echo"}
	c, _, err := websocket.DefaultDialer.Dial(U.String(), nil)
	flag.Parse()
	log.SetFlags(0)
	return c, err
}

func Write(c *websocket.Conn, t2 string) {
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
func Read(c *websocket.Conn, C1 chan bool, msg chan string) {
	for {
		_, message, err2 := c.ReadMessage()
		if err2 != nil {
			log.Println("read:", err2)
			return
		}
		m := string(message)
		log.Println("msg:", string(message))
		if m == "OverOk" {
			break
		} else {
			if m != "" {
				msg <- m
			}
		}
	}
	C1 <- true

}
