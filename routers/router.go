package routers

import (
	"github.com/astaxie/beego"
	"github.com/jetlwx/k8s_admin/controllers"
	"github.com/jetlwx/k8s_admin/controllers/componets"

	"github.com/jetlwx/k8s_admin/controllers/account"
	"github.com/jetlwx/k8s_admin/controllers/endpoints"
	"github.com/jetlwx/k8s_admin/controllers/node"
	"github.com/jetlwx/k8s_admin/controllers/pod"
	"github.com/jetlwx/k8s_admin/controllers/replication"
	"github.com/jetlwx/k8s_admin/controllers/service"
	"github.com/jetlwx/k8s_admin/controllers/virtualization"
	//	"github.com/jetlwx/k8s_admin/crontab"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//login
	beego.Router("/login", &account.AccountController{})
	//users center
	beego.AutoRouter(&account.UsersCenterController{})
	//初始化 distribution
	beego.Router("/distribution", &controllers.MainController{})
	//初始化Jetos配置
	beego.Router("/sysconfig", &controllers.JetosSysconfigController{})
	//kubecluster管理
	//beego.Router("/kubecluster", &controllers.KubeClusterController{})
	//show kubeNodeInf
	//beego.Router("/kubecluster/NodeInfo", &controllers.KubeNodeController{})
	//beego.Router("/kubecluster/specpodsinfo", &controllers.KubeSpecPodsController{})
	beego.Router("/kubecluster/Endpoints", &endpoints.EndPointsController{})
	//endpoints detail page show
	beego.Router("/kubecluster/specEndpoints", &endpoints.KubeSpecEndPointsController{})
	//componets
	beego.Router("/kubecluster/listComponets", &componets.KubeComponetsListController{})
	beego.Router("/kubecluster/specComponets", &componets.KubeSpecComponetsController{})

	beego.Router("/kubecluster/listServices", &service.KubelistServicesController{})
	//service detail page show
	beego.Router("/kubecluster/specServices", &service.KubeSpecServiceController{})
	// replication
	beego.Router("/kubecluster/listReplications", &replication.KubeReplicationListController{})
	beego.Router("/kubecluster/specReplications", &replication.KubeSpecReplicationController{})
	//nodes list
	beego.Router("/kubecluster/listNodes", &node.KubelistNodeController{})
	beego.Router("/kubecluster/specNode", &node.KubespecNodeController{})
	// pod
	beego.Router("/kubecluster/listPods", &pod.KubePostlistContronller{})
	//virtualization
	beego.Router("/virtualization/setting", &virtualization.VirtualizationController{})
	beego.Router("/virtualization/machinelist", &virtualization.VirtualizationMachineController{})

	beego.Router("/virtualization/editMachine", &virtualization.EditVirtualizationMachineController{})
	//virtualization auto router
	beego.AutoRouter(&virtualization.VirtualOperationController{})
	//vitual machine mother auto router
	beego.AutoRouter(&virtualization.VirtualHostMOtherController{})
	//timer and crontab
	// beego.AutoRouter(&crontab.CrontabController{})
	//test
	//beego.Router("/VirtualOperation/Msg", &virtualization.VtController{})
	//test
	beego.SetStaticPath("/logs", "logs")
	//vnc client
	beego.SetStaticPath("/VncClient", "VncClient")
}
