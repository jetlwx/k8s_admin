package componets

import (
	"github.com/astaxie/beego"
	m_componets "github.com/jetlwx/k8s_admin/models/componets"
	"strconv"
)

type KubeComponetsListController struct {
	beego.Controller
}

func (this *KubeComponetsListController) Get() {
	this.TplName = "kubecluster/componets/kube_list_componets.html"
	this.Data["IsComponets"] = true
	this.Data["IsClusternodes"] = false
	this.Data["IsServcie"] = false
	this.Data["IsReplication"] = false
	this.Data["IsPods"] = false

	c_d, httpcode := m_componets.Get_ComponentStatusList()
	if httpcode == 200 {
		this.Data["ComponetsList"] = c_d
	} else {
		this.Abort(strconv.Itoa(httpcode))
	}

}
