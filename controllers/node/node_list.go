package node

import (
	"github.com/astaxie/beego"
	m_node "github.com/jetlwx/k8s_admin/models/node"
	//"github.com/jetlwx/common"
	"strconv"
)

type KubelistNodeController struct {
	beego.Controller
}

func (this *KubelistNodeController) Get() {
	this.TplName = "kubecluster/node/list_node.html"
	this.Data["IsComponets"] = false
	this.Data["IsClusternodes"] = true
	this.Data["IsEndpoints"] = false
	this.Data["IsServcie"] = false
	this.Data["IsReplication"] = false
	this.Data["IsPods"] = false
	s_data, code := m_node.Get_NodeList()
	if code == 200 {
		this.Data["LISTNode"] = s_data
	} else {
		this.Abort(strconv.Itoa(code))
	}
}
