package virtualization

import (
	"log"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/jetlwx/k8s_admin/models"
	"github.com/jetlwx/k8s_admin/models/pager"
)

//VirtualizationMachineController
type VirtualizationMachineController struct {
	beego.Controller
}

func (this *VirtualizationMachineController) Get() {
	Name := this.GetSession("Name")
	log.Println("name:")
	if Name == nil {
		//this.Ctx.WriteString("Get userinfo [ name ]error")
		this.Redirect("/login", 302)
		return

	}
	sesid := this.GetSession("Id")
	if sesid == nil {
		//this.Ctx.WriteString("Get userinfo [ sesid ]error")
		this.Redirect("/login", 302)
		return

	}

	u, err := models.GetUserinfoById(sesid.(int64))
	if err != nil {
		this.Ctx.WriteString("User info get error")
		return
	}
	this.Data["PUTTY"] = u.Puttypath
	this.Data["PSFTP"] = u.Psftppath
	log.Println("putty---->", u.Puttypath)
	log.Println("psftp---->", u.Psftppath)
	this.Data["Islogin"] = true
	this.Data["Name"] = Name.(string)
	this.Data["IsIpPool"] = false
	this.Data["Ismachinelist"] = true
	this.Data["IsVirtualmachineMother"] = false
	this.Data["Proxy"] = beego.AppConfig.String("websocketPorxyHost")
	this.TplName = "virtualization/machinelist.html"
	var conditions string = "" //定义日志查询条件,格式为 "  "
	ser := this.GetString("ser")
	if ser != "" {
		conditions = ser
	}
	pno, _ := this.GetInt("pno") //获取当前请求页

	var po pager.PageOptions      //定义一个分页对象
	po.EnableFirstLastLink = true //是否显示首页尾页 默认false
	po.EnablePreNexLink = true    //是否显示上一页下一页 默认为false
	po.Conditions = conditions    // 传递分页条件 默认全表
	po.Currentpage = int64(pno)
	//传递当前页数,默认为1
	po.PageSize, _ = strconv.ParseInt(beego.AppConfig.String("vmlistNum"), 10, 64) //页面大小  默认为20

	//返回分页信息,
	//第一个:为返回的当前页面数据集合,ResultSet类型
	//第二个:生成的分页链接
	//第三个:返回总记录数
	//第四个:返回总页数
	//rs, pagerhtml, totalItem, _ := pager.GetPagerLinks(&po, this.Ctx)
	//log.Println("this.ctx", this.Ctx)
	totalItem := models.GetTotalItems(po.Conditions)
	totalpages, offset := pager.GetPageInfo(&po, totalItem)
	rs, err := models.GetpageRecordData(po.Conditions, int(po.PageSize), offset)
	Pagerhtml := pager.GetPagerLinks(&po, this.Ctx, totalpages)

	if err != nil {
		this.Ctx.WriteString("sql error")
		log.Println("err-->", err)
	}
	log.Println("totalpages", totalpages)

	//把当前页面的数据传递到前台
	this.Data["RecordData"] = rs
	this.Data["Pagerhtml"] = Pagerhtml
	this.Data["TotalItem"] = totalItem
	this.Data["PageSize"] = po.PageSize

}
