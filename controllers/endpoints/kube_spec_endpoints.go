package endpoints

import (
	"github.com/astaxie/beego"
	m_models "github.com/jetlwx/k8s_admin/models/endpoints"
)

type KubeSpecEndPointsController struct {
	beego.Controller
}

func (this *KubeSpecEndPointsController) Get() {
	this.TplName = "kubecluster/endpoints/Kube_spec_endpoints.html"
	namespace := this.GetString("namespace")
	name := this.GetString("name")
	this.Data["SPECENDPOINTS"] = m_models.Get_spec_endpoints(namespace, name)
}
