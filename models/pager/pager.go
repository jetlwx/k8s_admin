package pager

/**
 * 分页功能
//得到记录总数
//models layer
func GetTotalItems() int64 {
	tab := new(IpPool)
	total, err := engine.Count(tab)
	log.Println("err--------->", err)
	return total
}

//得到页面数据
func GetpageRecordData(conditions string, pagesize int, offset int) (interface{}, error) {
	//lens := endrecord - startrecorder

	var pool []IpPool
	log.Println("conditions-->", conditions)
	log.Println("pagesize-->", pagesize)
	log.Println("offset-->", offset)
	engine.ShowSQL()
	err := engine.Where("1>0 "+conditions).Desc("ipaddr").Limit(pagesize, offset).Find(&pool)
	return pool, err

}

//css

.a1{
  padding-left: 10px;
  padding-right: 10px;
  margin-left: 2px;
  margin-right: 2px;
  font-size: 14px;
  border:1px solid slategray;
  text-align: center;
  line-height: 30px;
  display:inline-block;
}
//html view
<div style="float:left;width:85%">{{.Pagerhtml}}</div>
<div style="float:left; "> Total:{{.TotalItem}}</div>

//controller
pno, _ := this.GetInt("pno") //获取当前请求页
	var conditions string = ""   //定义日志查询条件

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
	log.Println("this.ctx", this.Ctx)

	totalpages, totalItem, offset := pager.GetPageInfo(&po)
	rs, err := models.GetpageRecordData(po.Conditions, int(po.PageSize), offset)
	Pagerhtml := pager.GetPagerLinks(&po, this.Ctx, totalpages)
	// log.Println("totalpages", totalpages)
	// log.Println("rs2--->", rs)
	// log.Println("err1-->", err)
	// log.Println("Pagerhtml-->", Pagerhtml)
	//log.Println("err2-->", err)
	// if err1 != nil {

	// 	return
	// }
	//把当前页面的数据传递到前台
	this.Data["RecordData"] = rs
	this.Data["Pagerhtml"] = Pagerhtml
	this.Data["TotalItem"] = totalItem
	this.Data["PageSize"] = po.PageSize

*/
import (
	// "fmt"
	html "html/template"
	"log"
	con "strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/context"
)

type PageOptions struct {
	TableName           string //表名
	Conditions          string //条件
	Currentpage         int64  //当前页 ,默认1 每次分页,必须在前台设置新的页数,不设置始终默认1.在控制器中使用方式:cp, _ := this.GetInt("pno")   po.Currentpage = int(cp)
	PageSize            int64  //页面大小,默认20
	LinkItemCount       int64  //生成A标签的个数 默认10个
	Href                string //A标签的链接地址  ---------[不需要设置]
	ParamName           string //参数名称  默认是pno
	FirstPageText       string //首页文字  默认"首页"
	LastPageText        string //尾页文字  默认"尾页"
	PrePageText         string //上一页文字 默认"上一页"
	NextPageText        string //下一页文字 默认"下一页"
	EnableFirstLastLink bool   //是否启用首尾连接 默认false 建议开启
	EnablePreNexLink    bool   //是否启用上一页,下一页连接 默认false 建议开启
}

//得到分页信息－－>总记录条数，分页数，LIMIT offset
//offset IS pagesize*limit
func GetPageInfo(po *PageOptions, totalItem int64) (totalpages int64, offset int) {
	currentpage := po.Currentpage
	pagesize := po.PageSize
	defaultPagesize := int(po.PageSize)
	if currentpage <= 1 {
		currentpage = 1
		offset = int(currentpage) * defaultPagesize

	} else {
		offset = int(currentpage) * defaultPagesize
	}
	if pagesize == 0 {
		pagesize = int64(defaultPagesize)
		offset = int(currentpage) * defaultPagesize
	}

	if totalItem <= pagesize {
		totalpages = 1
		offset = 0
	} else if totalItem > pagesize {
		temp := totalItem / pagesize
		if (totalItem % pagesize) != 0 {

			temp = temp + 1
			offset = int(currentpage*pagesize - pagesize)

		}
		totalpages = temp
	}
	//get limit's start and  end
	//startrecorder = int((currentpage - 1) * pagesize)
	//	endrecord = int(currentpage)
	log.Println("totalItem--->", totalItem)
	log.Println("totalpages--->", totalpages)
	log.Println("offset--->", offset)
	return totalpages, offset
}

//生成链结地址
func GetPagerLinks(po *PageOptions, ctx *context.Context, totalpages int64) html.HTML {
	var str string = ""
	//	totalItem, totalpages, rs := GetPagesInfo(po.TableName, po.Currentpage, po.PageSize, po.Conditions)
	//	totalpages, _, _ := GetPageInfo(po)
	//rs, err1 := models.GetpageRecordData(po.Conditions, int(po.PageSize), offset)
	po = setDefault(po, totalpages)
	DealUri(po, ctx)
	if totalpages <= po.LinkItemCount {
		str = fun1(po, totalpages) //显示完全  12345678910
	} else if totalpages > po.LinkItemCount {
		if po.Currentpage < po.LinkItemCount {
			str = fun2(po, totalpages) //123456789...200
		} else {
			if po.Currentpage+po.LinkItemCount < totalpages {
				str = fun3(po, totalpages)
			} else {
				str = fun4(po, totalpages)
			}
		}
	}
	//	str = str + "<span Total:" + con.FormatInt(totalItem, 10) + "</span>"
	return html.HTML(str)
}

/**
 * 处理url,目的是保存参数
 */
func DealUri(po *PageOptions, ctx *context.Context) {
	uri := ctx.Request.RequestURI
	host := ctx.Request.Host

	protocol := strings.Split(ctx.Request.Proto, "/")[0]
	prefix := protocol + "://" + host
	// log.Println("url", ctx.Request.URL)
	// log.Println("host-->", ctx.Request.Host)
	// log.Println("uri->", uri)
	var rs string
	if strings.Contains(uri, "?") {
		arr := strings.Split(uri, "?")
		// log.Println("arr[0]->", arr[0])
		// log.Println("arr[1]->", arr[1])
		// log.Println("po.ParamName---> ", po.ParamName)
		rs = prefix + arr[0] + "?" + po.ParamName + "time=" + con.Itoa(time.Now().Second())
		arr2 := strings.Split(arr[1], "&")
		for _, v := range arr2 {
			log.Println("v->", v)
			if !strings.Contains(v, po.ParamName) {
				rs += "&" + v
			}
		}
	} else {
		rs = prefix + uri + "?" + po.ParamName + "time=" + con.Itoa(time.Now().Second())
	}
	po.Href = rs
	log.Println("po.Href--->", po.Href)
}

/**
 * 1...197 198 199 200
 */
func fun4(po *PageOptions, totalpages int64) string {
	var rs string = ""
	rs += getHeader(po, totalpages)
	log.Println("rs4-->head:", rs)
	rs += "<a class='a1' href=" + po.Href + "&" + po.ParamName + "=" + con.FormatInt(1, 10) + ">" + con.FormatInt(1, 10) + "</a>"
	rs += "<a  class='a1' href=>...</a>"
	for i := totalpages - po.LinkItemCount; i <= totalpages; i++ {
		if po.Currentpage != i {
			rs += "<a class='a1' href=" + po.Href + "&" + po.ParamName + "=" + con.FormatInt(i, 10) + ">" + con.FormatInt(i, 10) + "</a>"
		} else {
			rs += "<span class=\"current\">" + con.FormatInt(i, 10) + "</span>"
		}
	}
	rs += getFooter(po, totalpages)
	return rs

}

/**
 * 1...6 7 8 9 10 11 12  13  14 15... 200
 */
func fun3(po *PageOptions, totalpages int64) string {
	var rs string = ""
	rs += getHeader(po, totalpages)
	log.Println("rs3-->head:", rs)
	rs += "<a class='a1' href=" + po.Href + "&" + po.ParamName + "=" + con.FormatInt(1, 10) + ">" + con.FormatInt(1, 10) + "</a>"
	rs += "<a  class='a1' href=>...</a>"
	for i := po.Currentpage - po.LinkItemCount/2 + 1; i <= po.Currentpage+po.LinkItemCount/2-1; i++ {
		if po.Currentpage != i {

			C1 := con.FormatInt(i, 10)
			rs += "<a  class='a1' href=" + po.Href + "&" + po.ParamName + "=" + C1 + ">" + C1 + "</a>"
		} else {
			rs += "<span class=\"current\">" + con.FormatInt(i, 10) + "</span>"
		}
	}
	rs += "<a class='a1'  href=>...</a>"
	rs += "<a class='a1' href=" + po.Href + "&" + po.ParamName + "=" + con.FormatInt(totalpages, 10) + ">" + con.FormatInt(totalpages, 10) + "</a>"
	rs += getFooter(po, totalpages)
	return rs

}

/**
 * totalpages > po.LinkItemCount   po.Currentpage < po.LinkItemCount
 * 123456789...200
 */
func fun2(po *PageOptions, totalpages int64) string {
	var rs string = ""
	rs += getHeader(po, totalpages)
	log.Println("rs2-->head:", rs)
	for i := int64(1); i <= po.LinkItemCount+1; i++ {
		if i == po.LinkItemCount {
			rs += "<a  class='a1' href=\"" + po.Href + "&" + po.ParamName + "=" + con.FormatInt(i, 10) + "\">...</a>"
		} else if i == po.LinkItemCount+1 {
			rs += "<a   class='a1' href=\"" + po.Href + "&" + po.ParamName + "=" + con.FormatInt(totalpages, 10) + "\">" + con.FormatInt(totalpages, 10) + "</a>"
		} else {
			if po.Currentpage != i {
				rs += "<a  class='a1' href=" + po.Href + "&" + po.ParamName + "=" + con.FormatInt(i, 10) + ">" + con.FormatInt(i, 10) + "</a>"
			} else {
				rs += "<span class=\"current\">" + con.FormatInt(i, 10) + "</span>"
			}
		}
	}
	rs += getFooter(po, totalpages)
	return rs
}

/**
 * totalpages <= po.LinkItemCount
 * 显示完全  12345678910
 */
func fun1(po *PageOptions, totalpages int64) string {

	var rs string = ""
	rs += getHeader(po, totalpages)
	log.Println("rs1-->head:", rs)
	for i := int64(1); i <= totalpages; i++ {
		if po.Currentpage != i {
			rs += "<a class='a1' href=" + po.Href + "&" + po.ParamName + "=" + con.FormatInt(i, 10) + ">" + con.FormatInt(i, 10) + "</a>"
		} else {
			rs += "<span class=\"current\">" + con.FormatInt(i, 10) + "</span>"
		}
	}
	rs += getFooter(po, totalpages)
	return rs
}

/**
 * 头部
 */
func getHeader(po *PageOptions, totalpages int64) string {
	var rs string = "<div style='text-align:center;'>"
	if po.EnableFirstLastLink { //当首页,尾页都设定的时候,就显示

		rs += "<a class='a1' " + judgeDisable(po, totalpages, 0) + " href=" + po.Href + "&" + po.ParamName + "=" + con.FormatInt(1, 10) + ">" + po.FirstPageText + "</a>"
	}
	if po.EnablePreNexLink { // disabled=\"disabled\"
		var a int64 = po.Currentpage - 1
		if po.Currentpage == 1 {
			a = 1
		}
		rs += "<a class='a1' " + judgeDisable(po, totalpages, 0) + " href=" + po.Href + "&" + po.ParamName + "=" + con.FormatInt(a, 10) + ">" + po.PrePageText + "</a>"
	}
	return rs
}

/**
 * 尾部
 */
func getFooter(po *PageOptions, totalpages int64) string {
	var rs string = ""
	if po.EnablePreNexLink {
		var a int64 = po.Currentpage + 1
		if po.Currentpage == totalpages {
			a = totalpages
		}
		rs += "<a class='a1' " + judgeDisable(po, totalpages, 1) + "  href=" + po.Href + "&" + po.ParamName + "=" + con.FormatInt(a, 10) + ">" + po.NextPageText + "</a>"
	}
	if po.EnableFirstLastLink { //当首页,尾页都设定的时候,就显示
		rs += "<a class='a1' " + judgeDisable(po, totalpages, 1) + " href=" + po.Href + "&" + po.ParamName + "=" + con.FormatInt(totalpages, 10) + ">" + po.LastPageText + "</a>"
	}
	rs += "</div>"
	return rs
}

/**
 * 设置默认值
 */
func setDefault(po *PageOptions, totalpages int64) *PageOptions {
	if len(po.FirstPageText) <= 0 {
		//首页
		po.FirstPageText = "Index"
	}
	if len(po.LastPageText) <= 0 {
		//尾页
		po.LastPageText = "Last"
	}
	if len(po.PrePageText) <= 0 {
		//上一页
		po.PrePageText = "Pre"
	}
	if len(po.NextPageText) <= 0 {
		// 下一页
		po.NextPageText = "Next"
	}
	if po.Currentpage >= totalpages {
		po.Currentpage = totalpages
	}
	if po.Currentpage <= 1 {
		po.Currentpage = 1
	}
	if po.LinkItemCount == 0 {
		po.LinkItemCount = 10
	}
	if po.PageSize == 0 {
		po.PageSize = 20
	}
	if len(po.ParamName) <= 0 {
		po.ParamName = "pno"
	}
	return po
}

/**
 *判断首页尾页  上一页下一页是否能用
 */
func judgeDisable(po *PageOptions, totalpages int64, h_f int64) string {
	var rs string = ""
	//判断头部
	if h_f == 0 {
		if po.Currentpage == 1 {
			rs = "disabled=\"disabled\"  style=‘pointer-events:none;‘"
		}
	} else {
		if po.Currentpage == totalpages {
			rs = "disabled=\"disabled\"  style=‘pointer-events:none;‘"
		}
	}
	return rs
}
