package service

import (
	"encoding/json"
	"fmt"
	"github.com/jetlwx/k8s_admin/common"
	"github.com/jetlwx/k8s_admin/models"
	"github.com/jetlwx/k8s_admin/types"
)

func Get_ListAllNamespaceService() (*types.ServiceList, int) {
	conten := common.API_URL("NamespaceServiceList", "", "")
	fmt.Println("operation:", "NamespaceServiceList")
	fmt.Println("Get_ListAllNamespaceService.conten", conten)
	httpcode, res := models.Get_json_Strem(conten)
	if httpcode != 200 {
		fmt.Println("httpcode:", httpcode)
		return nil, httpcode
	}
	//fmt.Println("res=", res)
	s_end := &types.ServiceList{}
	json.Unmarshal(res, &s_end)
	fmt.Println("conten", s_end)
	//Ages
	lens := len(s_end.Items)
	for i := 0; i < lens; i++ {
		t0 := s_end.Items[i].Metadata.CreationTimestamp
		t1 := common.TimetoDayandHour(t0)
		s_end.Items[i].Ages = t1

	}
	return s_end, httpcode
}
