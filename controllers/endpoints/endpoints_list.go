package endpoints

import (
	//"fmt"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/httplib"
	//"github.com/jetlwx/k8s_admin/common"
	//"encoding/json"
	//"github.com/jetlwx/k8s_admin/controllers/endpoints"
	m_models "github.com/jetlwx/k8s_admin/models/endpoints"
	//"github.com/jetlwx/k8s_admin/models/service"
	"strconv"
)

type EndPointsController struct {
	beego.Controller
}

func (this *EndPointsController) Get() {
	this.TplName = "kubecluster/endpoints/Kube_Endpoints.html"
	this.Data["IsComponets"] = false
	this.Data["IsClusternodes"] = false
	this.Data["IsEndpoints"] = true
	this.Data["IsServcie"] = false
	this.Data["IsReplication"] = false
	this.Data["IsPods"] = false
	endpoints, httpcode := m_models.Get_EndpointsList()
	if httpcode == 200 {
		this.Data["EndPoints"] = endpoints
	} else {
		this.Abort(strconv.Itoa(httpcode))
	}

}
