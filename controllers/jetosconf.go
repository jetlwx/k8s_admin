package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jetlwx/k8s_admin/common"
	"github.com/jetlwx/k8s_admin/models"
	"strconv"
)

type JetosSysconfigController struct {
	beego.Controller
}

func (jetosconf *JetosSysconfigController) Get() {
	jetosconf.TplName = "sysconf/sysconfig.html"
	//Type用于控制各个标签的显示
	Type := jetosconf.GetString("type")
	if len(Type) < 1 || Type == "KubeDistribution" {

		jetosconf.Data["IsKubeDistribution"] = true
	} else {
		jetosconf.Data["IsKubeDistribution"] = false
	}
	switch Type {
	//if Type == "KubeMaster" {
	case "KubeMaster":
		jetosconf.Data["IsKubeMaster"] = true
		jetosconf.Data["IsKubeNode"] = false
		jetosconf.Data["IsKubeDistribution"] = false
		//获取kubemaster配置信息
		KubeMasterData, err := models.GetKubeMasterSetting()
		if err != nil {
			common.Writelog("Info", "获取KubeMaster配置出错")
		} else {
			jetosconf.Data["KubeMasterData"] = KubeMasterData

			fmt.Printf("%#v", KubeMasterData)
		}
	case "KubeNode":
		jetosconf.Data["IsKubeMaster"] = false
		jetosconf.Data["IsKubeNode"] = true
		jetosconf.Data["IsKubeDistribution"] = false
		fmt.Println("will done ....")
	case "KubeDistribution":
		jetosconf.Data["IsKubeMaster"] = false
		jetosconf.Data["IsKubeNode"] = false
		jetosconf.Data["IsKubeDistribution"] = true
		DistributionData, err := models.GetDistributionSetting()
		if err != nil {
			common.Writelog("Info", "获取DistributionData配置出错")
		} else {
			jetosconf.Data["DistributionData"] = DistributionData

			fmt.Printf("%#v", DistributionData)
		}

	}

}

func (jetoscof *JetosSysconfigController) Post() {

	Type := jetoscof.GetString("type")
	switch Type {
	case "KubeMaster":
		kbm_url := jetoscof.Input().Get("kbm_url")
		kbm_port, _ := strconv.ParseInt(jetoscof.Input().Get("kbm_port"), 10, 64)
		kbm_protocol := jetoscof.Input().Get("kmb_protocal1")
		fmt.Println("kbm_url", kbm_url)
		fmt.Println("kbm_port", kbm_port)
		fmt.Println("kmb_protocal1", kbm_protocol)
		err := models.AddKuberMasterSetting(kbm_url, kbm_port, kbm_protocol)

		if err != nil {
			fmt.Println(err)
			common.Writelog("Warn", "KubeMaster配置提交出错")
		} else {
			common.Writelog("Debug", "kbm_url="+kbm_url+"kbm_port="+string(kbm_port)+"kbm_protocol="+kbm_protocol+"提交至数据库成功！")
		}

		jetoscof.Redirect("/sysconfig?type=KubeMaster", 302)
	case "KubeNode":
		fmt.Println(" will done ....")
	case "KubeDistribution":
		distribution_url := jetoscof.Input().Get("distribution_url")
		distribution_protocol := jetoscof.Input().Get("distribution_protocol")
		err := models.AddDistributionSetting(distribution_url, distribution_protocol)
		if err != nil {
			fmt.Println(err)
			common.Writelog("Warn", "Distribution配置提交出错")
		} else {
			common.Writelog("Debug", "distribution_url="+distribution_url+"distribution_protocol="+distribution_protocol+"提交至数据库成功！")
		}

		jetoscof.Redirect("/sysconfig?type=KubeDistribution", 302)
	}
}
