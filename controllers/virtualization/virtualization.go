package virtualization

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/jetlwx/k8s_admin/common"
	m "github.com/jetlwx/k8s_admin/models"
)

type VirtualizationController struct {
	beego.Controller
}

//Get method
func (this *VirtualizationController) Get() {
	this.TplName = "virtualization/setting.html"
	if !common.IsLogin(this.GetSession("LoginName")) {
		this.Ctx.WriteString("login firt")
		return
	}

	this.Data["IsIpPool"] = true
	this.Data["Ismachinelist"] = false
	this.Data["IsVirtualmachineMother"] = false

}

//Post to models layers
func (this *VirtualizationController) Post() {
	//this.TplName = "virtualization/setting.html"

	this.Data["IsIpPool"] = true
	this.Data["Ismachinelist"] = false
	ipips := this.Input().Get("ip_ips")
	ipsuffix := this.Input().Get("ip_suffix")
	s, err := SplitIPsToip(ipips, ipsuffix)
	var msg []string
	if err == nil {
		for _, v := range s {
			log.Println("v->", v)
			_, err, errip := m.InsertIPPool(v)
			if err != nil {
				msg = append(msg, errip+"insert into db faild!")
				//	this.Ctx.WriteString(errip + "insert into db faild!<br>")
				log.Println(errip, "insert into db faild")
			}
			msg = append(msg, errip+"insert into db ok!")
			log.Println(v, "insert into db ok!")
			//this.Ctx.WriteString(v + "insert into db ok!<br>")
		}
	}

	this.Redirect("/virtualization/machinelist", 302)
	//this.Data["MSG"] = msg
	//fmt.Println("sssssss", s)

	return
}
