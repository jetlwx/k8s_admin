package models

import (
	"fmt"
	"reflect"

	"github.com/jetlwx/k8s_admin/common"
)

// return the type of parm
func TypeOfParm(i interface{}) (res string) {
	t := reflect.TypeOf(i)
	switch t.Kind() {
	case reflect.String:
		res = "string"
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		res = "int"
	default:
		res = "nothing"
	}
	return res
}

//获取HTTP,HTTPS json数据流的数据
func Get_json_Strem(conten string) (HttpStatusCode int, res []byte) {
	kbm_url, kbm_protocol, kbm_port, err := GetKubeMasterSettingDetail()
	if err != nil {
		return 404, nil
	}
	//url := common.GetKubeMasterApi(kbm_url, kbm_protocol, kbm_port, conten)
	url := common.GetKubeMasterApi(kbm_url, kbm_protocol, kbm_port, conten)
	fmt.Println("GetKubeMasterApi.url=", url)
	HttpStatusCode, res = common.GetData(url)
	fmt.Println("httpcod=", HttpStatusCode)
	return HttpStatusCode, res

}

func Get_Api_url(operation string, namespace string, name string) (url string) {
	kbm_url, kbm_protocol, kbm_port, err := GetKubeMasterSettingDetail()
	if err != nil {
		return ""
	}
	conten := common.API_URL(operation, namespace, name)
	url = common.GetKubeMasterApi(kbm_url, kbm_protocol, kbm_port, conten)
	fmt.Println("models.comm.url=", url)
	return url
}
