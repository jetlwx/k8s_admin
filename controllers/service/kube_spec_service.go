package service

import (
	"github.com/astaxie/beego"
	m_service "github.com/jetlwx/k8s_admin/models/service"
	"strconv"
)

type KubeSpecServiceController struct {
	beego.Controller
}

func (this *KubeSpecServiceController) Get() {
	this.TplName = "kubecluster/service/Kube_spec_service.html"
	namespace := this.GetString("namespace")
	name := this.GetString("name")
	s_data, code := m_service.Get_spec_service(namespace, name)
	if code == 200 {
		this.Data["SPECSERVICE"] = s_data
	} else {
		this.Abort(strconv.Itoa(code))
	}
}
