package service

import (
	"github.com/astaxie/beego"
	m_service "github.com/jetlwx/k8s_admin/models/service"
	//"github.com/jetlwx/common"
	"strconv"
)

type KubelistServicesController struct {
	beego.Controller
}

func (this *KubelistServicesController) Get() {
	this.TplName = "kubecluster/service/Kube_list_Service.html"
	this.Data["IsComponets"] = false
	this.Data["IsClusternodes"] = false
	this.Data["IsEndpoints"] = false
	this.Data["IsServcie"] = true
	this.Data["IsReplication"] = false
	this.Data["IsPods"] = false
	s_data, code := m_service.Get_ListAllNamespaceService()
	if code == 200 {
		this.Data["LISTSERVICES"] = s_data
	} else {
		this.Abort(strconv.Itoa(code))
	}
}
