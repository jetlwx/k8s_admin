package componets

import (
	"fmt"
	"github.com/astaxie/beego"
	m_componets "github.com/jetlwx/k8s_admin/models/componets"
	"strconv"
)

type KubeSpecComponetsController struct {
	beego.Controller
}

func (this *KubeSpecComponetsController) Get() {
	this.TplName = "kubecluster/componets/kube_spec_compontes.html"
	componet_name := this.GetString("name")
	fmt.Println("componet_name", componet_name)
	componets_data, httpcode := m_componets.Get_ComponentStatus(componet_name)
	if httpcode == 200 {
		this.Data["SpecComponets"] = componets_data
	} else {
		this.Abort(strconv.Itoa(httpcode))
	}
}
