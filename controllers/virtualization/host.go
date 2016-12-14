package virtualization

import (
	"log"
	"strconv"
	"strings"

	"github.com/astaxie/beego"

	"github.com/jetlwx/k8s_admin/common"
	m "github.com/jetlwx/k8s_admin/models"
	"github.com/jetlwx/websocket/websocketclient"
)

type VirtualHostMOtherController struct {
	beego.Controller
}

func (h *VirtualHostMOtherController) DeleteHost() {
	hostid := h.Input().Get("hostid")
	hostip := h.Input().Get("hostip")
	hid, e := strconv.Atoi(hostid)
	if e != nil {
		h.Ctx.WriteString("Get the HOSTID error:" + common.CustomerErr(e))
	}
	if !common.IsIP(hostip) {
		h.Ctx.WriteString("Get an invaid IP address")
	}
	err := m.DelHost(hostip, int64(hid))
	if err != nil {
		h.Ctx.WriteString("Delete host error--> " + common.CustomerErr(err))
	}
	h.Ctx.WriteString("Delete host " + hostip + " successful !!")

}

//common command
func (h *VirtualHostMOtherController) CommonCommand() {
	h.TplName = "virtualization/host/Command_group.html"
	if !common.IsLogin(h.GetSession("LoginName")) {
		h.Ctx.WriteString("login firt")
		return
	}
	defer func() {
		if r := recover(); r != nil {
			h.Ctx.WriteString(common.CustomerErr(r.(error)))
		}
	}()
	h.Data["IsCommandGroup"] = true
	h.Data["IsCommonComand"] = true
	h.Data["IsCreateVM"] = false
	log.Println("submit->", h.GetString("submit"))
	cType := h.Input().Get("command")
	var commandStr0 string
	switch cType {
	case "list":
		commandStr0 = "list"
	case "AllList":
		commandStr0 = "list --all"
	case "nodecpumap":
		commandStr0 = "nodecpumap"
	case "nodecpustats":
		commandStr0 = "nodecpustats"
	case "nodeinfo":
		commandStr0 = "nodeinfo"
	case "sysinfo":
		commandStr0 = "sysinfo"
	case "hostname":
		commandStr0 = "hostname"
	case "pool-list":
		commandStr0 = "pool-list"
	case "iface-list":
		commandStr0 = "iface-list"
	case "nodedev-list":
		commandStr0 = "nodedev-list"
	}
	if commandStr0 == "" {
		h.Ctx.WriteString("Get command error")
	}
	commandStr := "virsh" + " " + commandStr0
	hostid := h.Input().Get("hostid")

	hid, e := strconv.Atoi(hostid)
	if e != nil {
		h.Ctx.WriteString("Get the HOSTID error:" + common.CustomerErr(e))
	}

	hostinfo, err := m.GetHostInfo(int64(hid))
	if err != nil {
		h.Ctx.WriteString("get host infomation error")
	}

	hostip := hostinfo.Ipaddr
	//get the server socket info from DB
	WebsocketAgentPort := hostinfo.WebsocketAgentPort
	if WebsocketAgentPort == "" {
		h.Ctx.WriteString("The socket port is null ,Add firt")
		return
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
		h.Ctx.WriteString("err for Open socket--->" + common.CustomerErr(err))
		return
	}
	//write the messages
	websocketclient.WriteMsg(c, "CommonCommand", commandStr)
	//at the end of write message ,must give the end of single action,so must write the Over type to server
	websocketclient.WriteMsg(c, "Over", "")
	//read the message that return from server
	m, er := websocketclient.ForReadToReturn(c, C2, Cok)
	if er != nil {
		h.Data["Msg"] = common.CustomerErr(er)
	}
	//close the  read channel
	websocketclient.ReturnCloseForRead(c, C2)

	m = strings.Replace(m, " ", "&nbsp&nbsp", -1)
	m = strings.Replace(m, "<br>", "#br#", -1)
	m = strings.Replace(m, "<", " < ", -1)
	m = strings.Replace(m, "/>", " /> ", -1)
	m = strings.Replace(m, "#br#", "<br>", -1)
	h.Data["Msg"] = string(m)
	h.Data["HostID"] = hostid
	status := <-Cok
	//then you can follow the status do something
	log.Println("status:", status)
}

//create vm and show cratee logs
func (h *VirtualHostMOtherController) Createvm() {
	h.TplName = "virtualization/host/showcreatevmlog.html"
	defer func() {
		if r := recover(); r != nil {
			h.Ctx.WriteString(common.CustomerErr(r.(error)))
			return
		}
	}()

	var vmtemplateimages, vmtemplatexml string
	vmname := h.Input().Get("vmname")
	hostid := h.Input().Get("hostid")
	vmip := h.Input().Get("vmip")
	hostip := h.Input().Get("hostip")
	vmgw := h.Input().Get("GateWay")
	vmnetmask := h.Input().Get("NetMask")
	vmdns := h.Input().Get("DNS")
	vmtemplateimgtag := h.Input().Get("vmtemplate")
	script := h.Input().Get("createScript")

	hid, e := strconv.Atoi(hostid)
	if e != nil {
		h.Ctx.WriteString("Get the HOSTID error:" + common.CustomerErr(e))
	}

	hostinfo, err := m.GetHostInfo(int64(hid))
	if err != nil {
		h.Ctx.WriteString("get host infomation error")
	}
	WebsocketAgentPort := hostinfo.WebsocketAgentPort
	switch vmtemplateimgtag {
	case "1":
		vmtemplateimages = hostinfo.ImageTemplatePath1
		vmtemplatexml = hostinfo.XmlTemplatePath1
	case "2":
		vmtemplateimages = hostinfo.ImageTemplatePath2
		vmtemplatexml = hostinfo.XmlTemplatePath2
	case "3":
		vmtemplateimages = hostinfo.ImageTemplatePath3
		vmtemplatexml = hostinfo.XmlTemplatePath3
	case "4":
		vmtemplateimages = hostinfo.ImageTemplatePath4
		vmtemplatexml = hostinfo.XmlTemplatePath4
	case "5":
		vmtemplateimages = hostinfo.ImageTemplatePath5
		vmtemplatexml = hostinfo.XmlTemplatePath5
	}
	//prepare create vm
	if vmname == "" || vmip == "" || vmnetmask == "" || vmgw == "" || vmdns == "" || vmtemplateimages == "" || vmtemplatexml == "" {
		h.Ctx.WriteString(" some parameter of the create vm  require")
		return
	}
	createCommand := script + " " + vmname + " " + vmip + " " + vmnetmask + " " + vmgw + " " + vmdns + " " + vmtemplateimages + " " + vmtemplatexml

	//define an single chann
	var C2 = make(chan bool, 1)
	// define an single for check the command success or faild ,true is success and false is faild
	var CoK = make(chan bool, 1)
	addr := hostip + ":" + WebsocketAgentPort
	//open the websocket
	c, err := websocketclient.Open(addr)
	if err != nil {
		h.Ctx.WriteString("err for Open socket--->" + common.CustomerErr(err))
		log.Println("err for Open -> ", err)
		//attendtion: open success is the mostest condition, if open faild ,should Over current function advanced
		return
	}

	log.Println("Open Ok!")
	//define the tempfile for store the message of server return
	//Why do this ,because this command will run long time , so this way ,we can did it that write and read at the once
	filename, filepath := websocketclient.TempLogfile()
	// an gorouting thread to read the message from server and write to the tmp file
	go websocketclient.ForReadToTmpFile(c, C2, filepath, CoK)
	//write the command
	websocketclient.WriteMsg(c, "longTimeCommand", createCommand)
	//an gorouteing thread monitor the states single ,if recv the "Over" from server then close it
	go websocketclient.CloseForRead(c, C2)

	var status bool
	go func() {
		status = <-CoK
		log.Println("status->:", status)
		//you can follow the value of status do something
		if status == true {
			err := m.ChangeIPpoolAfterCreateVm(vmip, vmname, hostip)
			if err != nil {
				log.Println("func ChangeIPpoolAfterCreateVm error:", err)
			}
			log.Println("update ip pool succces")
		}
	}()

	logpath2 := beego.AppConfig.String("readcreatevmlogDIR")
	h.Data["Filename"] = logpath2 + filename
	return
}

//show Command_group pannel
func (h *VirtualHostMOtherController) CommandGroup() {
	h.TplName = "virtualization/host/Command_group.html"
	if !common.IsLogin(h.GetSession("LoginName")) {
		h.Ctx.WriteString("login firt")
		return
	}

	hostid := h.GetString("hostid")
	log.Println("hostid--->", hostid)
	hid, e := strconv.Atoi(hostid)
	if e != nil {
		h.Ctx.WriteString("Get the HOSTID error:" + common.CustomerErr(e))
	}

	command := h.Input().Get("command")
	switch command {
	case "createvm":
		h.Data["IsCommandGroup"] = false
		h.Data["IsCreateVM"] = true

		hostinfo, err := m.GetHostInfo(int64(hid))
		log.Println("hostinfo", hostinfo)
		if err != nil {
			h.Ctx.WriteString("get host infomation error")
			return
		}
		h.Data["HOST"] = hostinfo

		s2 := strings.Split(hostinfo.Ipaddr, ".")
		if len(s2) < 3 {
			h.Ctx.WriteString("host feauther get faild")
		}
		hostPrefix := s2[0] + "." + s2[1] + "." + s2[2] + "."
		log.Println("hostPrefix-->", hostPrefix)
		vmpool, err := m.GetvmIPpoolByprefix(hostPrefix, "nouseing")
		if err != nil {
			h.Ctx.WriteString("Get an error when get the vm pool list  " + common.CustomerErr(err))
			return
		}
		h.Data["VMPOOL"] = vmpool

	default:
		h.Data["IsCommandGroup"] = true
		h.Data["IsCreateVM"] = false
	}

}

//edit the host information submit page
func (h *VirtualHostMOtherController) Hostsubmit() {
	h.TplName = "virtualization/host/hostedit.html"
	host := m.Hosts{}
	log.Println("id", h.Input().Get("Id"))
	id, err := strconv.Atoi(h.Input().Get("Id"))
	if err != nil {
		h.Ctx.WriteString("get Hostinfo error:" + common.CustomerErr(err))
	}
	host.Id = int64(id)
	host.WebsocketAgentPort = h.Input().Get("WebsocketAgentPort")
	host.ImageTemplatePath1 = h.Input().Get("ImageTemplatePath1")
	host.XmlTemplatePath1 = h.Input().Get("XmlTemplatePath1")

	host.ImageTemplatePath2 = h.Input().Get("ImageTemplatePath2")
	host.XmlTemplatePath2 = h.Input().Get("XmlTemplatePath2")

	host.ImageTemplatePath3 = h.Input().Get("ImageTemplatePath3")
	host.XmlTemplatePath3 = h.Input().Get("XmlTemplatePath3")

	host.ImageTemplatePath4 = h.Input().Get("ImageTemplatePath4")
	host.XmlTemplatePath4 = h.Input().Get("XmlTemplatePath4")

	host.ImageTemplatePath5 = h.Input().Get("ImageTemplatePath5")
	host.XmlTemplatePath5 = h.Input().Get("XmlTemplatePath5")
	log.Println("host-------->", host)
	err1 := m.EditHostinfo(&host)
	if err1 != nil {
		h.Ctx.WriteString("update host info err" + common.CustomerErr(err1))
	}
	h.Redirect("/VirtualHostMOther/Hostedit?hostid="+strconv.Itoa(id), 302)
}

//edit the host information show page
func (h *VirtualHostMOtherController) Hostedit() {
	h.TplName = "virtualization/host/hostedit.html"
	if !common.IsLogin(h.GetSession("LoginName")) {
		h.Ctx.WriteString("login firt")
		return
	}

	hostid := h.GetString("hostid")
	log.Println("hostid--->", hostid)
	hid, e := strconv.Atoi(hostid)
	if e != nil {
		h.Ctx.WriteString("Get the HOSTID error:" + common.CustomerErr(e))
	}
	hostinfo, err := m.GetHostInfo(int64(hid))
	log.Println("hostinfo", hostinfo)
	if err != nil {
		h.Ctx.WriteString("get host infomation error")
	}
	h.Data["HOST"] = hostinfo
}

//add host and show the host list
func (vh *VirtualHostMOtherController) Host() {
	vh.TplName = "virtualization/host/list_add.html"
	if !common.IsLogin(vh.GetSession("LoginName")) {
		vh.Ctx.WriteString("login firt")
		return
	}
	vh.Data["IsIpPool"] = false
	vh.Data["Ismachinelist"] = false
	vh.Data["IsVirtualmachineMother"] = true

	//show hosts
	hosts, err := m.ShowAllHost()
	if err != nil {
		log.Println("get host list error:", err)
		vh.Ctx.WriteString("get host list err,show the log")
	}
	vh.Data["Hostlist"] = hosts
}

//host add post page and action
func (vh *VirtualHostMOtherController) Post() {
	//vh.Ctx.WriteString("dddddd")
	if !common.IsLogin(vh.GetSession("LoginName")) {
		vh.Ctx.WriteString("login firt")
		return
	}

	ipips := vh.Input().Get("hostsip")
	ipsuffix := vh.Input().Get("ip_suffix")
	location := vh.Input().Get("location")
	log.Println("v->")
	s, err := SplitIPsToip(ipips, ipsuffix)
	log.Println("s-->", s)
	var msg []string
	if err == nil {
		for _, v := range s {
			log.Println("v->", v)
			_, errip, err := m.AddHosts(v, location)
			if err != nil {
				msg = append(msg, errip+"insert into db faild!")
				//	vh.Ctx.WriteString(errip + "insert into db faild!<br>")
				log.Println(errip, "insert into db faild")
			}
			msg = append(msg, errip+"insert into db ok!")
			log.Println(v, "insert into db ok!")
			//vh.Ctx.WriteString(v + "insert into db ok!<br>")
		}
	}

	vh.Redirect("/VirtualHostMOther/Host", 302)
	return
}
