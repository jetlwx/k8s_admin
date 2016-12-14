package pod

import (
	"github.com/astaxie/beego"
	//"github.com/jetlwx/common"
	m_pod "github.com/jetlwx/k8s_admin/models/pod"
	//"strconv"
	"fmt"
)

type KubePostlistContronller struct {
	beego.Controller
}

func (this *KubePostlistContronller) Get() {
	this.TplName = "kubecluster/pod/kube_list_pod.html"
	this.Data["IsComponets"] = false
	this.Data["IsClusternodes"] = false
	this.Data["IsEndpoints"] = false
	this.Data["IsServcie"] = false
	this.Data["IsReplication"] = false
	this.Data["IsPods"] = true
	s_data := m_pod.Get_models_KubePodsList()

	this.Data["PodS"] = s_data
	fmt.Println(m_pod.Get_spec_Pod("default", "kube-dns-v11-u1vih"))

}
