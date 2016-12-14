package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/jetlwx/k8s_admin/common"
	"github.com/jetlwx/k8s_admin/models"
	"github.com/jetlwx/k8s_admin/types"
)

func Get_EndpointsList() (*types.EndpointsList, int) {
	conten := common.API_URL("ListEndpoints", "", "")
	fmt.Println("operation:", "ListEndpoints")
	fmt.Println("Get_EndpointsList.conten", conten)
	httpcode, res := models.Get_json_Strem(conten)
	if httpcode != 200 {
		fmt.Println("httpcode:", httpcode)
		return nil, httpcode
	}
	//fmt.Println("res=", res)
	s_end := &types.EndpointsList{}
	json.Unmarshal(res, &s_end)
	fmt.Println("conten", s_end)
	for k, v := range s_end.Items {
		t0 := v.Metadata.CreationTimestamp
		s_end.Items[k].Ages = common.TimetoDayandHour(t0)
	}
	return s_end, httpcode
}
