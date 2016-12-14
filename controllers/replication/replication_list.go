package replication

import (
	//"fmt"
	"github.com/astaxie/beego"
	m_models "github.com/jetlwx/k8s_admin/models/replication"
	"strconv"
)

type KubeReplicationListController struct {
	beego.Controller
}

func (this *KubeReplicationListController) Get() {
	this.TplName = "kubecluster/replication/kube_replication_list.html"
	this.Data["IsComponets"] = false
	this.Data["IsClusternodes"] = false
	this.Data["IsEndpoints"] = false
	this.Data["IsServcie"] = false
	this.Data["IsReplication"] = true
	this.Data["IsPods"] = false
	replication, httpcode := m_models.Get_ReplicationControllerList()
	if httpcode == 200 {
		this.Data["Replication"] = replication
	} else {
		this.Abort(strconv.Itoa(httpcode))
	}

}
