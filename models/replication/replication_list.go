package replication

import (
	"encoding/json"
	"fmt"
	"github.com/jetlwx/k8s_admin/common"
	"github.com/jetlwx/k8s_admin/models"
	"github.com/jetlwx/k8s_admin/types"
)

func Get_ReplicationControllerList() (*types.ReplicationControllerList, int) {
	conten := common.API_URL("replication_listall", "", "")
	fmt.Println("operation:", "replication_listall")
	fmt.Println("Get_ComponentStatusList.conten", conten)
	httpcode, res := models.Get_json_Strem(conten)
	if httpcode != 200 {
		fmt.Println("httpcode:", httpcode)
		return nil, httpcode
	}
	//fmt.Println("res=", res)
	s_end := &types.ReplicationControllerList{}
	json.Unmarshal(res, &s_end)
	for k, v := range s_end.Items {
		t0 := v.Metadata.CreationTimestamp
		s_end.Items[k].Ages = common.TimetoDayandHour(t0)
	}
	fmt.Println("conten", s_end)

	return s_end, httpcode
}
