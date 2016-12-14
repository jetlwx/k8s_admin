package componets

import (
	"encoding/json"
	"fmt"
	"github.com/jetlwx/k8s_admin/common"
	"github.com/jetlwx/k8s_admin/models"
	"github.com/jetlwx/k8s_admin/types"
)

func Get_ComponentStatus(componetsName string) (*types.ComponentStatus, int) {
	conten := common.API_URL("SpeccomponetsStatus", "", componetsName)
	fmt.Println("operation:", "ComponentStatus")
	fmt.Println("Get_ComponentStatus.conten", conten)
	httpcode, res := models.Get_json_Strem(conten)
	if httpcode != 200 {
		fmt.Println("httpcode:", httpcode)
		return nil, httpcode
	}
	//fmt.Println("res=", res)
	s_end := &types.ComponentStatus{}
	json.Unmarshal(res, &s_end)
	fmt.Println("conten", s_end)

	return s_end, httpcode
}
