package crontab

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/jetlwx/k8s_admin/common"
	"github.com/jetlwx/k8s_admin/models"
	"github.com/jetlwx/websocket/websocketclient"
	"github.com/robfig/cron"
	"log"
	"strings"
)

func StartCrontab() {
	c_cleartokenfile := beego.AppConfig.String("clear_host_vnc_token")
	c_updatevmmac := beego.AppConfig.String("Check_vm_state")
	c_updatedomainstate := beego.AppConfig.String("Scan_vm_mac")
	cr := cron.New()
	cr2 := cron.New()
	err1 := cr.AddFunc(c_cleartokenfile, CrontabClearTokenFiles)
	if err1 != nil {
		common.Writelog("E", "[Crontab] clear token file cron err:", common.CustomerErr(err1))
	}
	err2 := cr.AddFunc(c_updatevmmac, CrontabUpdateVMMac)
	if err2 != nil {
		common.Writelog("E", "[Crontab] update vm mac cron err:", common.CustomerErr(err2))
	}
	err3 := cr2.AddFunc(c_updatedomainstate, CrontabUpdateDomainState)
	if err3 != nil {
		common.Writelog("E", "[Crontab] update domain state cron err:", common.CustomerErr(err3))
	}

	cr.Start()
	cr2.Start()

}

//UpdateDomainState
func CrontabUpdateDomainState() {
	vmlist := models.Getvirtualmachinelist()
	for _, v := range vmlist {
		domainname := v.VmName
		if domainname == "" {
			continue
		}
		hostip := v.HostsIp
		host, err := models.GetHostInfoByIp(hostip)
		if err != nil {
			common.Writelog("E", "[Crontab] GetHostInfoByIp"+common.CustomerErr(err))
			return
		}
		if v.Useing == false {
			continue
		}
		websocketport := host.WebsocketAgentPort
		if websocketport == "" {
			common.Writelog("C", "The host "+hostip+" WebsocketAgentPort is not found")
		}
		commstr := "virsh domstate" + " " + domainname

		str, e := CallBackServer(hostip, "CommonCommand", commstr)

		str = strings.Replace(str, "<br>", "", -1)
		if e != nil {
			common.Writelog("E", "[Crontab] "+common.CustomerErr(e))
			str = "Unknow"
		}
		if str == "" || len(str) > 8 {
			str = "Unknow"
		}
		vm := models.IpPool{}
		vm.VmName = v.VmName
		vm.Id = v.Id
		vm.VmState = str
		count, e := models.UpdateVmvmstate(vm)
		if count == 0 && e == nil {
			common.Writelog("I", "[Crontab] "+"vm host:"+domainname+" state is newleast state: "+str+" ignor to update")
		}
		if count == 1 {
			common.Writelog("I", "[Crontab] vm host"+domainname+" state has update to:"+str)
		}
		if e != nil {
			common.Writelog("E", "[Crontab]  update vm host "+domainname+"faild->"+common.CustomerErr(e))
		}
	}
}

//ClearTokenFiles
func CrontabClearTokenFiles() {
	tokendir := beego.AppConfig.String("vncTokenDir")
	if tokendir == "" {
		common.Writelog("E", "[Crontab] "+"beego.AppConfig:vncTokenDir is not found")
	}
	hostlist, err := models.ShowAllHost()
	if err != nil {
		common.Writelog("E", common.CustomerErr(err))
	}
	for _, v := range hostlist {
		hostip := v.Ipaddr
		websocktport := v.WebsocketAgentPort
		if websocktport == "" {
			common.Writelog("C", "The host "+hostip+" WebsocketAgentPort is not found")
		}
		str, e := CallBackServer(hostip, "ClearVNCToken", tokendir)
		if e != nil {
			common.Writelog("E", "[Crontab] "+common.CustomerErr(e))
			continue
		}
		common.Writelog("D", str)
	}
}

//update vm  mac address
func CrontabUpdateVMMac() {
	macSqlCondition := "useing = 1"
	vmHostList, err := models.GetvmIPlistByConditions(macSqlCondition)
	if err != nil {
		common.Writelog("E", common.CustomerErr(err))
	}
	for _, v := range vmHostList {
		vmname := v.VmName
		hostip := v.HostsIp
		vmid := v.Id
		mac, e := CallBackServer(hostip, "GetVMMac", vmname)
		if e != nil {
			common.Writelog("E", common.CustomerErr(e))
			continue
		}
		mac = strings.Replace(mac, "<br>", "", -1)
		common.Writelog("I", mac)
		if mac != "" {
			h := models.IpPool{}
			h.Id = vmid
			h.VmName = vmname
			h.HostsIp = hostip
			h.VmMac = mac
			count, er := models.UpdateVmIPMacaddr(h)
			if er != nil {
				common.Writelog("E", common.CustomerErr(er))
			}
			if count != 0 {
				common.Writelog("I", "[Crontab] "+vmname+" MAC address"+mac+" has been Update")
			}
			if count == 0 && err == nil {
				common.Writelog("I", "[Crontab] "+"The mac address "+mac+" is the newleast ,ignor update")
			}
		}
	}

}

// func Updatedomainstate2() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			common.Writelog("recover information:", common.CustomerErr(r.(error)))
// 		}
// 	}()
// 	hostlist, err := models.ShowAllHost()
// 	if err != nil {
// 		common.Writelog("E", "Get an occour when get the host list:"+common.CustomerErr(err))
// 		return
// 	}
// 	for _, h := range hostlist {
// 		hostip := h.Ipaddr
// 		SQL := "hosts_ip='" + hostip + "' and useing = 1"
// 		vmhostlist, err := models.GetvmIPlistByConditions(SQL)
// 		if err != nil {
// 			common.Writelog("E", "get host"+hostip+" vmlist error"+common.CustomerErr(err))
// 			continue
// 		}
// 		C2, Cok, c := CallBackServer2(hostip)
// 		if C2 == nil || Cok == nil || c == nil {
// 			continue
// 		}
// 		for _, vm := range vmhostlist {

// 			commstr := "virsh domstate" + " " + vm.VmName
// 			WriteMsgToBackServer(c, "CommonCommand", commstr)
// 			msg, err := ReadMsgFromBackServer(c, C2, Cok)
// 			log.Println("e---->", err)
// 			log.Println("msg---->", msg)
// 		}
// 		WriteMsgToBackServer(c, "Over", "")
// 		go CheckChanState(c, C2, Cok)
// 	}

// }

// func CallBackServer2(hostip string) (chan bool, chan bool, *websocket.Conn) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			common.Writelog("recover information:", common.CustomerErr(r.(error)))
// 		}
// 	}()

// 	host, err := models.GetHostInfoByIp(hostip)
// 	if err != nil {
// 		log.Println("err:", err)
// 	}

// 	//get the server socket info from DB
// 	WebsocketAgentPort := host.WebsocketAgentPort
// 	if WebsocketAgentPort == "" || WebsocketAgentPort == "0" {
// 		log.Println("The socket port is null ,Add firt")
// 		common.Writelog("E", "The host "+hostip+" socket port is null ,Add firt")
// 		return nil, nil, nil
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
// 		common.Writelog("E", "err for Open remote -> "+common.CustomerErr(err))

// 	}
// 	return C2, Cok, c

// }

// //WriteMsgToBackServer must call  CallBackServer mothed
// func WriteMsgToBackServer(c *websocket.Conn, commandType string, commandStr string) {
// 	websocketclient.WriteMsg(c, commandType, commandStr)
// }
// func ReadMsgFromBackServer(c *websocket.Conn, C2 chan bool, Cok chan bool) (string, error) {
// 	m, er := websocketclient.ForReadToReturn(c, C2, Cok)
// 	if er != nil {
// 		return "", errors.New("an occour found at read message from remote server:" + common.CustomerErr(er))
// 	}

// 	return m, nil

// }
// func CheckChanState(c *websocket.Conn, C2 chan bool, Cok chan bool) {
// 	go websocketclient.ReturnCloseForRead(c, C2)

// 	status := <-Cok
// 	log.Println("status:", status)
// }

func CallBackServer(hostip string, commandType string, commandStr string) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			common.Writelog("C" + "recover information:" + common.CustomerErr(r.(error)))
		}
	}()

	host, err := models.GetHostInfoByIp(hostip)
	if err != nil {
		log.Println("err:", err)
	}

	//get the server socket info from DB
	WebsocketAgentPort := host.WebsocketAgentPort
	if WebsocketAgentPort == "" || WebsocketAgentPort == "0" {
		log.Println("The socket port is null ,Add firt")
		return "", errors.New("The host " + hostip + " socket port is null ,Add firt")
	}
	addr := hostip + ":" + WebsocketAgentPort
	//define the over channel
	var C2 = make(chan bool, 1)
	//define the status of command ,success or faild
	var Cok = make(chan bool, 1)
	//open the socket
	c, err := websocketclient.Open(addr)
	//defer c.Close()
	if err != nil {
		log.Println("err for Open -> ", err)
		return "", errors.New("err for Open remote -> " + common.CustomerErr(err))

	}
	//write the messages
	websocketclient.WriteMsg(c, commandType, commandStr)
	//at the end of write message ,must give the end of single action,so must write the Over type to server
	websocketclient.WriteMsg(c, "Over", "")
	//read the message that return from server
	m, er := websocketclient.ForReadToReturn(c, C2, Cok)
	if er != nil {
		return "", errors.New("an occour found at read message from remote server:" + common.CustomerErr(er))
	}
	//close the  read channel
	go websocketclient.ReturnCloseForRead(c, C2)

	status := <-Cok
	//then you can follow the status do something
	log.Println("status:", status)
	return m, nil
}
