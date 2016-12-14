package node

import (
	"encoding/json"
	"fmt"
	"github.com/jetlwx/k8s_admin/common"
	"github.com/jetlwx/k8s_admin/models"
	"github.com/jetlwx/k8s_admin/types"
)

func Get_Node(nodename string) (*types.Node, int) {
	conten := common.API_URL("SpecNode", "", nodename)
	fmt.Println("operation:", "SpecNode")
	fmt.Println("Get_Node.conten", conten)
	httpcode, res := models.Get_json_Strem(conten)
	if httpcode != 200 {
		fmt.Println("httpcode:", httpcode)
		return nil, httpcode
	}
	//fmt.Println("res=", res)
	s_end := &types.Node{}
	json.Unmarshal(res, &s_end)
	fmt.Println("conten", s_end)
	for k, v := range s_end.Status.Images {
		s_end.Status.Images[k].SizeAuto = common.TransferByte(v.SizeBytes)
		fmt.Println("s_end.Status.Images[k].SizeAuto", s_end.Status.Images[k].SizeAuto)
	}
	return s_end, httpcode
}
