package node

import (
	"github.com/astaxie/beego"
	m_node "github.com/jetlwx/k8s_admin/models/node"
	//"github.com/jetlwx/common"
	"strconv"
)

type KubespecNodeController struct {
	beego.Controller
}

func (this *KubespecNodeController) Get() {
	this.TplName = "kubecluster/node/spec_node.html"
	name := this.GetString("nodename")
	s_data, code := m_node.Get_Node(name)
	if code == 200 {
		this.Data["SPECNode"] = s_data
	} else {
		this.Abort(strconv.Itoa(code))
	}
}
