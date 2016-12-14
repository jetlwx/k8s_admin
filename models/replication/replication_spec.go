package replication

import (
	"fmt"
	//"github.com/Jeffail/gabs"
	"encoding/json"
	"github.com/jetlwx/k8s_admin/common"
	"github.com/jetlwx/k8s_admin/models"
	"github.com/jetlwx/k8s_admin/types"
)

func Get_spec_Replication(namespace string, name string) (*types.ReplicationController, int) {
	conten := common.API_URL("ReadSpecReplication", namespace, name)
	fmt.Println("Get_spec_Replication.conten", conten)
	httpcode, res := models.Get_json_Strem(conten)
	if httpcode != 200 {
		fmt.Println("httpcode:", httpcode)
		return nil, httpcode
	}
	//fmt.Println("res=", res)
	s_end := &types.ReplicationController{}
	json.Unmarshal(res, &s_end)
	//fmt.Println(s_end)
	Ages := common.TimetoDayandHour(s_end.Metadata.CreationTimestamp)
	s_end.Ages = Ages
	//fmt.Println("--->", s_end.Spec.Template.Spec.Volumes[0].HostPath)
	return s_end, httpcode
}
