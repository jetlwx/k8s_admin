package service

import (
	"fmt"
	//"github.com/Jeffail/gabs"
	"encoding/json"
	"github.com/jetlwx/k8s_admin/common"
	"github.com/jetlwx/k8s_admin/models"
	"github.com/jetlwx/k8s_admin/types"
)

func Get_spec_service(namespace string, name string) (*types.Service, int) {
	conten := common.API_URL("readSpecService", namespace, name)
	httpcode, res := models.Get_json_Strem(conten)
	if httpcode != 200 {
		fmt.Println("httpcode:", httpcode)
		return nil, httpcode
	}
	//fmt.Println("res=", res)
	s_end := &types.Service{}
	json.Unmarshal(res, &s_end)
	//fmt.Println(s_end)
	return s_end, 200
}
