package replication

import (
	//"fmt"
	"github.com/astaxie/beego"
	//"github.com/jetlwx/k8s_admin/models"
	m_models "github.com/jetlwx/k8s_admin/models/replication"
	//"github.com/jetlwx/k8s_admin/types"
	//"github.com/jetlwx/k8s_admin/yaml.v2"
	//"log"
	"strconv"
)

type KubeSpecReplicationController struct {
	beego.Controller
}

func (this *KubeSpecReplicationController) Get() {
	this.TplName = "kubecluster/replication/kube_spec_replication.html"
	namespace := this.GetString("namespace")
	name := this.GetString("name")
	rc, httpcode := m_models.Get_spec_Replication(namespace, name)
	if httpcode == 200 {
		this.Data["SPECReplication"] = rc
	} else {
		this.Abort(strconv.Itoa(httpcode))
	}

	// replication_two := "namespaces/" + namespace + "/replicationcontrollers/" + name

	// m := &types.ReplicationController{}
	// _, mydata := models.Get_json_Strem(replication_two)
	// //_, mydata := common.GetData("http://172.16.6.160:8080/api/v1/pods")
	// err := yaml.Unmarshal(mydata, &m)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("--- m:\n%v\n\n", m)

	// d, err := yaml.Marshal(&m)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }

	// fmt.Printf("--- m dump:\n%s\n\n", string(d))
	// this.Data["PODSDETAIL"] = string(d)
}
