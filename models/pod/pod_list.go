package pod

import (
	//	js1 "encoding/json"
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/jetlwx/k8s_admin/common"
	"github.com/jetlwx/k8s_admin/models"
	"strconv"
	"strings"
)

type KubePodsList struct {
	Name                  string
	ReplicationController string
	NameSpace             string
	Ages                  string
	Labels                string
	PodsStatus            string
	HostIP                string
	PodIP                 string
	Contaions_ready       string // ok/all
	RestartPolicy         string
	//注释的为备用字段，勿删
	//DnsPolicy                     string
	NodeName string
	//TerminationGracePeriodSeconds int
	//SecurityContext               string
	RestartNumber    int
	Container_Status []KubePods_Container_status
}

type KubePods_Container_status struct {
	Name string
	//RuningAt     string
	Ready        string
	RestartCount int
	//Image        string
	//ImageID      string
	//ContainerID  string
}

func Get_models_KubePodsList() []KubePodsList {
	conten := common.API_URL("Pod_list", "", "")
	http_code, body := models.Get_json_Strem(conten)
	if http_code != 200 {
		fmt.Print(http_code)
		return nil
	}
	json, _ := gabs.ParseJSON(body)
	c, _ := json.S("items").Children()
	len_c, err := json.ArrayCount("items")
	if err != nil {
		fmt.Println("Get kube KubePodsList list err", err)
		common.Writelog("Critical", "KubePodsList"+common.CustomerErr(err))
		return nil
	}
	//fmt.Println("c--->", "len_c", len_c)

	pods := make([]KubePodsList, len_c)

	for k0, v := range c {
		t_name := v.Path("metadata.name").String()
		t_name = strings.Replace(t_name, "\"", "", -1)

		t_labels := v.Path("metadata.labels").String()
		t_labels = strings.Replace(t_labels, "\"", "", -1)
		t_labels = strings.Replace(t_labels, "}", "", -1)
		t_labels = strings.Replace(t_labels, "{", "", -1)

		t_namespace := v.Path("metadata.namespace").String()
		t_namespace = strings.Replace(t_namespace, "\"", "", -1)

		t_ages0 := v.Path("metadata.creationTimestamp").Data().(string)
		t_ages := common.TimetoDayandHour(t_ages0)

		t_rc0, _ := v.Path("metadata.annotations").Children()

		var t_rc_name string //controller
		for _, tb := range t_rc0 {
			v4 := strings.Replace(tb.String(), "\"", "", 1)
			v5 := strings.Split(v4, "\\n")
			v6 := strings.Replace(v5[0], "\\", "", -1)
			j2, _ := gabs.ParseJSON([]byte(v6))
			t_rc1 := j2.Path("reference.name").String()
			t_rc_name = t_rc1

			fmt.Println("t_rc1--->", t_rc1)
		}
		t_podsstatus := v.Path("status.phase").String()
		t_podsstatus = strings.Replace(t_podsstatus, "\"", "", -1)

		t_hostip := v.Path("status.hostIP").String()
		t_hostip = strings.Replace(t_hostip, "\"", "", -1)

		t_podip := v.Path("status.podIP").String()
		t_podip = strings.Replace(t_podip, "\"", "", -1)
		//注释的为备用字段，勿删

		// t_restartpolicy := v.Path("spec.restartPolicy").String()
		// t_restartpolicy = strings.Replace(t_restartpolicy, "\"", "", -1)

		// t_terminationGrace := int(v.Path("spec.terminationGracePeriodSeconds").Data().(float64))

		// t_dnsPolicy := v.Path("spec.dnsPolicy").String()
		// t_dnsPolicy = strings.Replace(t_dnsPolicy, "\"", "", -1)

		t_nodeName := v.Path("spec.nodeName").String()
		t_nodeName = strings.Replace(t_nodeName, "\"", "", -1)

		// t_securitycontext := v.Path("spec.securityContext").String()

		pods[k0].Name = t_name
		pods[k0].NameSpace = t_namespace
		pods[k0].Labels = t_labels
		pods[k0].Ages = t_ages

		t_rc_name = strings.Replace(t_rc_name, "\"", "", -1)
		pods[k0].ReplicationController = t_rc_name
		pods[k0].PodsStatus = t_podsstatus
		pods[k0].HostIP = t_hostip
		pods[k0].PodIP = t_podip
		// pods[k0].RestartPolicy = t_restartpolicy
		// pods[k0].TerminationGracePeriodSeconds = t_terminationGrace
		// pods[k0].DnsPolicy = t_dnsPolicy
		pods[k0].NodeName = t_nodeName
		// pods[k0].SecurityContext = t_securitycontext

		t_container, _ := v.Path("status.containerStatuses").Children()
		len_t_container := len(t_container)
		c_ready_num := len_t_container
		c_restart_num := 0
		container_status := make([]KubePods_Container_status, len_t_container)
		for k1, tv := range t_container {
			fmt.Println("k1=", k1, "tv=", tv)
			t_container_name := tv.Path("name").String()
			t_container_name = strings.Replace(t_container_name, "\"", "", -1)

			// t_container_running := tv.Path("state.running.startedAt").String()
			t_container_ready := tv.Path("ready").String()
			if t_container_ready != "true" {
				c_ready_num -= 1
			}
			t_container_restartCount := int(tv.Path("restartCount").Data().(float64))
			if t_container_restartCount > c_restart_num {
				c_restart_num = t_container_restartCount
			}
			// t_container_image := tv.Path("image").String()
			// t_container_imageId := tv.Path("imageID").String()
			// t_container_containerId := tv.Path("containerID").String()

			container_status[k1].Name = t_container_name
			// container_status[k1].RuningAt = t_container_running
			container_status[k1].Ready = t_container_ready
			container_status[k1].RestartCount = t_container_restartCount
			// container_status[k1].Image = t_container_image
			// container_status[k1].ImageID = t_container_imageId
			// container_status[k1].ContainerID = t_container_containerId

		}
		pods[k0].Container_Status = container_status
		c_ready := strconv.Itoa(c_ready_num) + "/" + strconv.Itoa(len_t_container)
		pods[k0].RestartNumber = c_restart_num
		pods[k0].Contaions_ready = c_ready

	}
	fmt.Println("pods------->", pods)
	return pods

}
