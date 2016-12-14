package virtualization

import (
	"log"
	"regexp"
	"strconv"

	"strings"

	"github.com/astaxie/beego"
	"github.com/jetlwx/k8s_admin/common"
	"github.com/jetlwx/k8s_admin/models"
)

type EditVirtualizationMachineController struct {
	beego.Controller
}

func (this *EditVirtualizationMachineController) Get() {
	this.TplName = "virtualization/editMachineInfo.html"
	switchType := this.GetString("type")
	if switchType != "editMachine" {
		this.Ctx.WriteString("parameter send error!!")
	}

	machinId := this.GetString("machinId")
	log.Println("machinId", machinId)
	b, _ := regexp.MatchString("[0-9]", machinId)
	if !b {
		this.Ctx.WriteString("parameter send error!!")
	}
	m, _ := strconv.Atoi(machinId)
	info, err := models.GetIPPoolById(m)
	if err != nil {
		this.Ctx.WriteString("sql ERROR at get ippool")
		log.Println(err)
	}
	this.Data["PoolInfo"] = info
	log.Println("info->", info)

	//get host list
	s2 := strings.Split(info.Ipaddr, ".")
	if len(s2) < 3 {
		this.Ctx.WriteString("host feauther get faild")
	}
	hostPrefix := s2[0] + "." + s2[1] + "." + s2[2] + "."
	hostlist, err := models.GetHostlistByPrefix(hostPrefix)
	if err != nil {
		this.Ctx.WriteString("get host list faild")
	}
	this.Data["HostLIst"] = hostlist
}
func (this *EditVirtualizationMachineController) Post() {
	vmhost := new(models.IpPool)
	// vmhost.VmName = this.Input().Get("VmName")
	vmhost.VmDesc = this.Input().Get("VmDesc")
	// vmhost.HostsIp = this.Input().Get("Host")
	// log.Println("vmhost.VmName", vmhost.VmName)
	log.Println("vmhost.VmDesc", vmhost.VmDesc)
	// log.Println("vmhost.HostsIp", vmhost.HostsIp)
	id1, err := strconv.Atoi(this.Input().Get("id"))
	vmhost.Id = int64(id1)
	if err != nil {
		this.Ctx.WriteString("id is null")
	}

	if vmhost.Id != 0 && (vmhost.VmName != "" || vmhost.VmDesc != "") {
		eff, err := models.UpdateIPPool(*vmhost)
		log.Println("eff----->", eff)
		if eff == 0 {
			this.Ctx.WriteString("the virtual machine name is already exist ")
		}
		if err != nil {
			log.Println("update ip pool err:", err)
			this.Ctx.WriteString(common.CustomerErr(err))
			//this.Ctx.WriteString("sql ERROR at update ippool")

		}
		// if res != 1 {
		// 	this.Ctx.WriteString("update error")
		// }
		//this.Ctx.WriteString("update ok!")
		this.Redirect("/virtualization/editMachine?machinId="+strconv.Itoa(id1)+"&type=editMachine", 302)
	}
}
