package main

import (
	//	"github.com/codeskyblue/go-sh"
	//	"log"
	// "errors"
	// "github.com/jetlwx/k8s_admin/common"
	// "github.com/jetlwx/k8s_admin/models"
	// "github.com/jetlwx/websocket/websocketclient"
	//"log"

	// "os/exec"
	// "strings"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	dir := "/tmp/token"
	f, _ := ioutil.ReadDir(dir)
	for _, v := range f {
		n := v.Name()
		e := os.Remove(dir + "/" + n)
		log.Println(e)
	}
	//domain := "test35"
	//	c1 := "/usr/bin/virsh domblklist " + domain
	// disk1, err := sh.Command(c1).Output()
	// // b_buf := bytes.NewBuffer(disk1)

	// // disk := int(disk1) - 2
	// log.Println(err)
	// // log.Println("disk", disk1)
	// comm := "/usr/bin/virsh domblklist test35 | grep -v '^$'|wc -l"
	// cmd, err := exec.Command("/bin/sh", "-c", comm).Output()
	// log.Println(err)
	// log.Println("disk:", string(cmd))
	// diskcount := 5
	// diskcount = diskcount - 2
	// diskname := rune(97 + diskcount)
	// log.Println("diskname", string(diskname))
	// devname, imgname, err := GetDiskDeviceName("test35", "IDE")
	// log.Println("devname", devname, "e", err)
	// s, _ := GetDiskPath("test35", imgname)
	// log.Println("s", s)
	// l, _ := ListDomainsTargetDisk("test33")
	// log.Println(l)
	// str := "virsh vncdisplay test33 |172.16.6.26|/tmp/token"
	// // GentVNCTokenfile(str)
	// a, _ := os.Getwd()
	// log.Println(a)
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		log.Println("panic <--->", r.(error))
	// 	}
	// }()
	// hostip := "172.16.6.26"
	// com := "CommonCommand"
	// str := "virsh domiflist test33 | grep ':' |awk -F' ' '{print $NF}'"
	// a := CallBackServer(hostip, com, str)
	// log.Println("a:", a)

}

// func CallBackServer(hostip string, commandType string, commandStr string) string {
// 	host, err := models.GetHostInfoByIp(hostip)
// 	if err != nil {
// 		log.Println("err:", err)
// 	}

// 	//get the server socket info from DB
// 	WebsocketAgentPort := host.WebsocketAgentPort
// 	if WebsocketAgentPort == "" {
// 		log.Println("The socket port is null ,Add firt")
// 	}
// 	addr := hostip + ":" + WebsocketAgentPort
// 	//define the over channel
// 	var C2 = make(chan bool, 1)
// 	//define the status of command ,success or faild
// 	var Cok = make(chan bool, 1)
// 	//open the socket
// 	c, err := websocketclient.Open(addr)
// 	//defer c.Close()
// 	if err != nil {
// 		log.Println("err for Open -> ", err)

// 	}
// 	//write the messages
// 	websocketclient.WriteMsg(c, commandType, commandStr)
// 	//at the end of write message ,must give the end of single action,so must write the Over type to server
// 	websocketclient.WriteMsg(c, "Over", "")
// 	//read the message that return from server
// 	m, er := websocketclient.ForReadToReturn(c, C2, Cok)
// 	if er != nil {

// 	}
// 	//close the  read channel
// 	go websocketclient.ReturnCloseForRead(c, C2)

// 	status := <-Cok
// 	//then you can follow the status do something
// 	log.Println("status:", status)
// 	return m
// }

// func GentVNCTokenfile(str string) (string, error) {
// 	s := strings.Split(str, "|")
// 	displayportcmd := s[0]
// 	hostip := s[1]
// 	tokenpath := s[2]
// 	if displayportcmd == "" || hostip == "" || tokenpath == "" {
// 		return "", errors.New("Some information null")
// 	}
// 	cmd, err := exec.Command("/bin/sh", "-c", displayportcmd).Output()
// 	if string(cmd) == "" {
// 		return "", errors.New("The domain have no vnc config")
// 	}
// 	if err != nil {
// 		log.Println("err", err)
// 		return "", err
// 	}
// 	t := common.RandomString()
// 	log.Println("token:", t)
// 	log.Println("displayport", string(cmd))
// 	log.Println("hostip", hostip)
// 	log.Println("tokendir", tokenpath)

// 	p := strings.Split(string(cmd), ":")
// 	remoteport := "590" + p[1]

// 	writetokenstr := t + ":" + " " + hostip + ":" + remoteport
// 	log.Println("writetokenstr", writetokenstr)

// 	dstFile, err3 := os.Create(tokenpath + "/" + t)
// 	if err3 != nil {
// 		log.Println("err3:", err3)
// 		return "", err3
// 	}
// 	defer dstFile.Close()
// 	n, err2 := dstFile.WriteString(writetokenstr)
// 	if err2 != nil {
// 		return "", err2
// 	}
// 	log.Println("byte-n", n)
// 	return t, nil
// }

// func gentNextTargetBus(domainname string, driverType string) (bus string, err error) {

// 	var busType string
// 	switch driverType {
// 	case "VirtIO":
// 		busType = "v"
// 	case "IDE":
// 		busType = "h"
// 	}
// 	comm := "virsh domblklist " + domainname + " | sed -n '3,$'p | grep -v '^$' | awk -F' ' '{print $1}'|grep" + " " + "'^" + busType + "'"
// 	log.Println("comm", comm)

// 	cmd, err := exec.Command("/bin/sh", "-c", comm).Output()
// 	if string(cmd) == "" {
// 		return busType + "da", nil
// 	}
// 	if err != nil {
// 		log.Println("err", err)
// 		return "", err
// 	}
// 	log.Println("cmd", string(cmd))

// 	str := strings.Split(string(cmd), "\n")
// 	s := str[0]
// 	if s == "" {
// 		return "", err
// 	}
// 	t := strings.Split(s, "")
// 	third := t[len(t)-1]
// 	second := t[len(t)-2]
// 	for _, v := range str {
// 		if v != "" {
// 			t2 := strings.Split(v, "")
// 			t2Third := t2[len(t2)-1]
// 			t2Second := t2[len(t2)-2]
// 			if t2Third > third {
// 				third = t2Third
// 			}
// 			if t2Second > second {
// 				second = t2Second
// 			}
// 		}

// 	}
// 	log.Println("max second:", second, "max third:", third)
// 	//log.Println([]int(third))
// 	a := string(third[0] + 1)
// 	log.Println("ffff->", a)
// 	//convert to ascii
// 	i := third[0]
// 	i2 := second[0]

// 	if i >= 122 {
// 		second = string(i2 + 1)
// 		log.Println("second", second)
// 		third = "a"
// 	} else {
// 		third = string(i + 1)
// 		log.Println("third", third)
// 	}

// 	switch driverType {
// 	case "VirtIO":
// 		bus = "v" + string(second) + string(third)
// 	case "IDE":
// 		bus = "h" + string(second) + string(third)
// 	}
// 	log.Println("bus->", bus)
// 	return bus, nil

// }

// //list the domain of target disk
// func ListDomainsTargetDisk(domainname string) (string, error) {
// 	comm := "virsh domblklist" + " " + domainname + " " + "| sed -n '3,$'p | grep -v '^$' | awk -F' ' '{print $1}'"
// 	cmd, err := exec.Command("/bin/sh", "-c", comm).Output()
// 	if err != nil {
// 		log.Println("err", err)
// 		return "", err
// 	}
// 	log.Println("cmd", string(cmd))

// 	str := strings.Split(string(cmd), "\n")
// 	var list string
// 	for _, v := range str {
// 		if v != "" {
// 			list = list + "|" + v
// 		}
// 	}
// 	log.Println("list-->", list)
// 	return list, nil
// }
