package virtualization

import (
	"log"
	//	"strings"

	"github.com/astaxie/beego"
	"github.com/jetlwx/k8s_admin/common"
	"github.com/jetlwx/k8s_admin/models"
	"github.com/jetlwx/websocket/websocketclient"
	"os"
	"strconv"
	"strings"
)

//VirtualOperationController auto route
type VirtualOperationController struct {
	beego.Controller
}

// type Msg struct {
// 	Type   string
// 	Conten string
// }

func (v *VirtualOperationController) DomainCommand() {
	v.TplName = "virtualization/vm/commanddetailshow.html"
	if !common.IsLogin(v.GetSession("LoginName")) {
		v.Ctx.WriteString("login firt")
		return
	}

	var commandStr string
	var com1 string
	var args string
	var commandType string
	defer func() {
		if r := recover(); r != nil {
			v.Ctx.WriteString(common.CustomerErr(r.(error)))
		}
	}()
	hostip := v.GetString("hostip")
	//vmid := v.Input().Get("Vmid")
	vmname := v.GetString("vmname")
	//vmip := v.Input().Get("Vmip")
	command := v.GetString("type")
	log.Println("vmname-->", vmname)
	log.Println("command-->", command)
	if command == "" {
		v.Ctx.WriteString("Command not found")
		v.StopRun()
	}
	switch command {
	case "hostname":
		if vmname == "" {
			v.Ctx.WriteString("vm Host name not found")
			v.StopRun()
		}
		com1 = "virsh domhostname"
		args = vmname
		commandType = "CommonCommand"
	case "DomainState":
		com1 = "virsh domstats"
		args = vmname
		commandType = "CommonCommand"

	case "DomainBlockError":
		com1 = "virsh domblkerror"
		args = vmname
		commandType = "CommonCommand"

	case "DomaminControl":
		com1 = "virsh domcontrol"
		args = vmname
		commandType = "CommonCommand"

	case "DomainBlockState":
		com1 = "virsh domblkstat"
		args = vmname
		commandType = "CommonCommand"

	case "DomainBlockList":
		com1 = "virsh domblklist"
		args = vmname
		commandType = "CommonCommand"

	case "DomainInfo":
		com1 = "virsh dominfo"
		args = vmname
		commandType = "CommonCommand"

	case "DomainInterfaceList":
		com1 = "virsh domiflist"
		args = vmname
		commandType = "CommonCommand"

	case "RunningDomains":
		com1 = "virsh list"
		args = ""
		commandType = "CommonCommand"

	case "AllDomains":
		com1 = "virsh list"
		args = "--all"
		commandType = "CommonCommand"

	case "startDomain":
		com1 = "virsh start"
		args = vmname
		commandType = "CommonCommand"

	case "ShutDown":
		com1 = "virsh shutdown"
		args = vmname
		commandType = "CommonCommand"

	case "Suspend":
		com1 = "virsh suspend"
		args = vmname
		commandType = "CommonCommand"

	case "Reboot":
		com1 = "virsh reboot"
		args = vmname
		commandType = "CommonCommand"
	case "AutoStart":
		com1 = "virsh autostart"
		args = vmname
		commandType = "CommonCommand"
	case "DisableAutoStart":
		com1 = "virsh autostart"
		args = vmname + " " + "--disable"
		commandType = "CommonCommand"

	case "Resume":
		com1 = "virsh resume"
		args = vmname
		commandType = "CommonCommand"

	case "Destroy":
		com1 = "virsh destroy"
		args = vmname
		commandType = "CommonCommand"
	case "AttachDisk":
		disksize := v.GetString("disksize")
		driver := v.GetString("driver")
		_, err := strconv.Atoi(disksize)
		if err != nil {
			v.Ctx.WriteString("disksize Must be a digit")
			v.StopRun()
		}

		com1 = disksize
		args = vmname + " " + driver
		commandType = "AttachDisk"

		//for DetachDomainTargetDisk
	case "GetDomainDiskList":
		com1 = "NoNe"
		args = vmname
		commandType = "GetDomainDiskList"

	case "DetachDomainTargetDisk":
		targetdevice := v.GetString("targetdevice")
		com1 = "virsh detach-disk"
		args = vmname + " " + targetdevice + " " + "--live"
		commandType = "CommonCommand"

		//for DetachDomainAndXMLTargetDisk
	case "GetDomainDiskList2":
		com1 = "NoNe"
		args = vmname
		commandType = "GetDomainDiskList"

	case "DetachDomainAndXMLTargetDisk":
		targetdevice := v.GetString("targetdevice")
		com1 = "NoNe"
		args = vmname + " " + targetdevice
		commandType = "DetachDomainAndXMLTargetDisk"
	case "ShowVcpuNumber":
		com1 = "virsh vcpucount"
		args = vmname
		commandType = "CommonCommand"
	case "DomainMemory":
		com1 = "virsh dommemstat"
		args = vmname
		commandType = "CommonCommand"
	case "SetVCpuNumber":
		cpunum := v.GetString("cpumumber")
		cpuset := v.GetString("cpuset")
		var setmodel string
		_, err := strconv.Atoi(cpunum)
		if err != nil {
			v.Ctx.WriteString("cpu number must be a digit")
			v.StopRun()
		}
		switch cpuset {
		case "MaxConfig":
			setmodel = "--maximum --config"
		case "CurrentLive":
			setmodel = "--live"
		case "CurrentConfig":
			setmodel = "--current"
		}
		com1 = "virsh setvcpus"
		args = vmname + " " + "--count" + " " + cpunum + " " + setmodel
		commandType = "CommonCommand"
	//	v.Ctx.WriteString(args)

	case "SetMemory":
		memnumber := v.GetString("memnumber")
		mem, err := strconv.Atoi(memnumber)
		if err != nil {
			v.Ctx.WriteString("memory must be a digit type or float")
			v.StopRun()
		}
		umem := uint64(mem) * 1024 * 1024
		memset := v.GetString("memset")
		var setmodel string
		switch memset {
		case "MaxConfig":
			com1 = "virsh setmaxmem"
			setmodel = "--config"
		case "CurrentLive":
			com1 = "virsh setmem"
			setmodel = "--live"
		case "CurrentConfig":
			com1 = "virsh setmem"
			setmodel = "--current"
		}
		args = vmname + " " + "--size" + " " + strconv.FormatUint(umem, 10) + " " + setmodel
		commandType = "CommonCommand"

	case "VNC":
		com1 = "virsh vncdisplay"
		tokenpath := beego.AppConfig.String("vncTokenDir")
		args = vmname + "|" + hostip + "|" + tokenpath
		commandType = "VNCToken"
	default:
		v.Ctx.WriteString("No default, Forbiden")
		v.StopRun()
	}

	if com1 == "" {
		v.Ctx.WriteString("Get Remote command null,Forbiden")
		v.StopRun()
	}

	commandStr = com1 + " " + args
	//	commandStr = "virsh" + " " + command
	log.Println("commandStr---->", commandStr)
	host, err := models.GetHostInfoByIp(hostip)
	if err != nil {
		v.Ctx.WriteString("Get Host infomation err -> " + common.CustomerErr(err))
		v.StopRun()
	}

	//get the server socket info from DB
	WebsocketAgentPort := host.WebsocketAgentPort
	if WebsocketAgentPort == "" {
		v.Ctx.WriteString("The socket port is null ,Add firt")
		v.StopRun()
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
		v.Ctx.WriteString("err for Open socket--->" + common.CustomerErr(err))
		v.StopRun()
	}
	//write the messages
	websocketclient.WriteMsg(c, commandType, commandStr)
	//at the end of write message ,must give the end of single action,so must write the Over type to server
	websocketclient.WriteMsg(c, "Over", "")
	//read the message that return from server
	m, er := websocketclient.ForReadToReturn(c, C2, Cok)
	if er != nil {
		v.Data["Msg"] = common.CustomerErr(er)
	}
	//close the  read channel
	go websocketclient.ReturnCloseForRead(c, C2)
	if command != "VNC" {
		m = strings.Replace(m, " ", "&nbsp&nbsp", -1)
		m = strings.Replace(m, "\n", "<br>", -1)
	}
	// m = strings.Replace(m, "<br>", "#br#", -1)
	// m = strings.Replace(m, "<", " < ", -1)
	// m = strings.Replace(m, "/>", " /> ", -1)
	// m = strings.Replace(m, "#br#", "<br>", -1)
	if command == "VNC" {
		m = strings.Replace(m, "<br>", "", -1)
	}
	log.Println("m----====->", m)
	switch command {
	case "GetDomainDiskList":
		v.Data["Hostip"] = hostip
		v.Data["Vmname"] = vmname
		v.Data["IsMsg"] = false
		v.Data["IsGetDomainDiskList"] = true
		disklist := getDomainDiskList(m)
		v.Data["DiskList"] = disklist
		v.Data["DeleteFromDomain"] = true
		v.Data["DetachDomainAndXMLTargetDisk"] = false

	case "GetDomainDiskList2":
		v.Data["Hostip"] = hostip
		v.Data["Vmname"] = vmname
		v.Data["IsMsg"] = false
		v.Data["IsGetDomainDiskList"] = true
		disklist := getDomainDiskList(m)
		v.Data["DiskList"] = disklist
		v.Data["DeleteFromDomain"] = false
		v.Data["DetachDomainAndXMLTargetDisk"] = true
	case "VNC":
		tokenpath := beego.AppConfig.String("vncTokenDir")
		str := strings.Split(m, ":")
		tokenfilename := str[0]
		dstFile, err3 := os.Create(tokenpath + "/" + tokenfilename)
		if err3 != nil {
			log.Println("err3:", err3)
		}
		defer dstFile.Close()
		_, err2 := dstFile.WriteString(m)
		if err2 != nil {
			log.Println("err2:", err2)
		}

		vnchttpport := beego.AppConfig.String("vnchttpport")
		v.Redirect("/VncClient/vnc_auto.html?token="+tokenfilename+"&host="+hostip+"&port="+vnchttpport, 302)
	default:
		v.Data["Msg"] = string(m)
		v.Data["IsMsg"] = true
		v.Data["IsGetDomainDiskList"] = false

	}

	status := <-Cok
	//then you can follow the status do something
	log.Println("status:", status)

}

//getdomainlist
func getDomainDiskList(m string) []string {
	var s []string
	str := strings.Split(m, "|")
	for _, v := range str {
		if v != "" {
			v = strings.Replace(v, "<br>", "", -1)
			s = append(s, v)
		}
	}
	return s
}

//show CommonGroup list
func (c *VirtualOperationController) CommandGroup() {
	c.TplName = "virtualization/vm/commandgroup.html"
	if !common.IsLogin(c.GetSession("LoginName")) {
		c.Ctx.WriteString("login firt")
		return
	}

	vmip := c.GetString("vmip")
	vmname := c.GetString("vmname")
	vmid := c.GetString("vmid")
	hostip := c.GetString("hostip")
	if !common.IsIP(hostip) {
		c.Ctx.WriteString("HOST IP is inavid")
	}
	if !common.IsIP(vmip) {
		c.Ctx.WriteString("vm IP is inavid")
	}
	if vmip == "" || vmname == "" || hostip == "" || vmid == "" {
		c.Ctx.WriteString("some parameter inavid")
	}

	//log.Println("vmip->", vmip, "vmid-->", vmid, "vmname->", vmname, "hostip->", hostip)
	c.Data["Vmip"] = vmip
	c.Data["Vmid"] = vmid
	c.Data["Vmname"] = vmname
	c.Data["Hostip"] = hostip

}

// machinID := c.Input().Get("machinId")
// //commandStr := c.Input().Get("command")
// m, err2 := common.UrlParaIsInt(machinID)
// if err2 != nil {
// 	c.Ctx.WriteString("the machine id send error")
// }
// info, err3 := models.GetIPPoolById(m)
// if err3 != nil {
// 	c.Ctx.WriteString("sql ERROR at get ippool")
// 	log.Println(err3)
// }
// c.Data["Vminfo"] = info
// host, e := models.GetHostInfoByIp(info.HostsIp)
// if e != nil {
// 	c.Ctx.WriteString("Get Host websocket Port error -> " + common.CustomerErr(e))
// }

// //get the server socket info from DB
// WebsocketAgentPort := host.WebsocketAgentPort
// addr := info.HostsIp + ":" + WebsocketAgentPort
// //define the over channel
// var C2 = make(chan bool, 1)
// //open the socket
// w, err := websocketclient.Open(addr)
// defer w.Close()
// if err != nil {
// 	log.Println("err for Open -> ", err)
// 	c.Ctx.WriteString("err for Open socket--->" + common.CustomerErr(err))
// }
// //write the messages
// websocketclient.WriteMsg(w, "CommonCommand", commandStr)
// //at the end of write message ,must give the end of single action,so must write the Over type to server
// websocketclient.WriteMsg(w, "Over", "")
// //read the message that return from server
// m1, er := websocketclient.ForReadToReturn(w, C2)
// if er != nil {
// 	c.Data["Msg"] = common.CustomerErr(er)
// }
// //close the  read channel
// websocketclient.CloseForRead(w, C2)

// m1 = strings.Replace(m1, " ", "&nbsp&nbsp", -1)
// m1 = strings.Replace(m1, "<br>", "#br#", -1)
// m1 = strings.Replace(m1, "<", " < ", -1)
// m1 = strings.Replace(m1, "/>", " /> ", -1)
// m1 = strings.Replace(m1, "#br#", "<br>", -1)
// c.Data["Msg"] = string(m1)

//}

// //var addr = flag.String("addr", "127.0.0.1:9090", "http service address")

// //DelMachine VirtualOperationController delete virtual machine
// func (vo *VirtualOperationController) DelMachine() {
// 	vo.TplName = "virtualization/deletMachineInfo.html"
// 	switchType := vo.GetString("type")
// 	if switchType != "editMachine" {
// 		vo.Ctx.WriteString("parameter send error!!")
// 	}

// 	machinId := vo.GetString("machinId")
// 	log.Println("machinId", machinId)
// 	m, err := common.UrlParaIsInt(machinId)
// 	if err != nil {
// 		vo.Ctx.WriteString("the machine id send error")
// 	}
// 	info, err := models.GetIPPoolById(m)
// 	if err != nil {
// 		vo.Ctx.WriteString("sql ERROR at get ippool")
// 		log.Println(err)
// 	}

// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	vo.Data["PoolInfo"] = info
// 	vo.Data["ShowDetail"] = true

// }
// func (c *VirtualOperationController) OperateMachine() {
// 	c.TplName = "virtualization/OperateVM.html"
// 	Type := c.GetString("type")
// 	id, err1 := strconv.Atoi(c.GetString("machinId"))
// 	if err1 != nil {
// 		c.Ctx.WriteString(common.CustomerErr(err1))
// 	}
// 	inf, err := models.GetIPPoolById(id)
// 	if err != nil {
// 		c.Ctx.WriteString(common.CustomerErr(err))
// 	}
// 	c.Data["PoolInfo"] = inf
// 	c.Data["Type"] = Type

// 	//get host list
// 	s2 := strings.Split(inf.Ipaddr, ".")
// 	if len(s2) < 3 {
// 		c.Ctx.WriteString("host feauther get faild")
// 	}
// 	hostPrefix := s2[0] + "." + s2[1] + "." + s2[2] + "."
// 	hostlist, err := models.GetHostlistByPrefix(hostPrefix)
// 	if err != nil {
// 		c.Ctx.WriteString("get host list faild")
// 	}
// 	c.Data["HostLIst"] = hostlist

// }

func (c *VirtualOperationController) SSH() {
	stype := c.GetString("stype")
	sessionsid := c.GetSession("Id")
	if sessionsid == nil {
		c.Ctx.WriteString("session is lost ,relogin please.")
	}
	sid, err := strconv.ParseInt(sessionsid.(string), 10, 64)
	if err != nil {
		c.Ctx.WriteString("session id is wrong!!!")
	}
	u, err := models.GetUserinfoById(sid)
	if err != nil {
		c.Ctx.WriteString("Get user info error" + common.CustomerErr(err))
	}
	putty := u.Puttypath
	psftp := u.Psftppath
	switch stype {
	case "ssh":
		c.Data["path"] = putty
	case "sftp":
		c.Data["path"] = psftp
	}

}
