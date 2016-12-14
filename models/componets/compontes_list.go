package componets

import (
	"encoding/json"
	"fmt"
	"github.com/jetlwx/k8s_admin/common"
	"github.com/jetlwx/k8s_admin/models"
	"github.com/jetlwx/k8s_admin/types"
)

func Get_ComponentStatusList() (*types.ComponentStatusList, int) {
	conten := common.API_URL("ComponetStatusList", "", "")
	fmt.Println("operation:", "ComponetStatusList")
	fmt.Println("Get_ComponentStatusList.conten", conten)
	httpcode, res := models.Get_json_Strem(conten)
	if httpcode != 200 {
		fmt.Println("httpcode:", httpcode)
		return nil, httpcode
	}
	//fmt.Println("res=", res)
	s_end := &types.ComponentStatusList{}
	json.Unmarshal(res, &s_end)
	fmt.Println("conten", s_end)

	return s_end, httpcode
}
