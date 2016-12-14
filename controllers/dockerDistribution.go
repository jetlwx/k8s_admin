package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/logs"
	"github.com/jetlwx/k8s_admin/common"
	"github.com/jetlwx/k8s_admin/dockerDistribution"
	"github.com/jetlwx/k8s_admin/models"
	//"log"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "distribution/admin.html"
	Dis, err := models.GetDistributionSetting()
	if err != nil {
		c.Ctx.WriteString("获取镜像出错")
		return
	}
	Dis_url := Dis.DistributionUrl
	Dis_protocol := Dis.DistributionProtocol
	url := Dis_protocol + "://" + Dis_url
	fmt.Println("url", url)
	PagePerNumber := beego.AppConfig.String("distributionPerNum")
	c.Data["PagePerNumber"] = PagePerNumber

	//accessurl := comm.GetImgListUrl(url)
	other := "?n=" + PagePerNumber + "&last=" + c.GetString("last") + "&pre=" + c.GetString("pre")
	weburl := common.GetImgListUrl(url, other)
	//定义一个MAP
	Mimages := make(map[string]string)

	keywords := c.GetString("search")
	Httpcode, imagesName := dockerDistribution.GetImageNames(weburl, keywords)
	fmt.Println(imagesName)
	if Httpcode == 200 {
		ILen := len(imagesName)
		if len(c.GetString("pre")) == 0 {
			c.Data["Isnotpre"] = true
		}
		if ILen > 0 {
			for _, Iname := range imagesName {
				fmt.Println("Iname", Iname)
				hcode, tags := dockerDistribution.GetImageTags(url, Iname)
				fmt.Println("tags==", tags)
				if hcode == 200 {
					//将MAP的键和值写进MAP中
					Mimages[Iname] = tags
				}

			}

			pg, _ := strconv.Atoi(PagePerNumber)
			if ILen < pg || len(imagesName[ILen-1]) == 0 {
				c.Data["IstoLast"] = true
				c.Data["Pre"] = true

			}

			c.Data["Images"] = Mimages
			c.Data["LastImage"] = imagesName[ILen-1]
			c.Data["PreImage"] = imagesName[0]

			if imagesName[ILen-1] == imagesName[0] {
				c.Data["Allthelast"] = false
			}

		} //ILen over

		//	fmt.Println(imagesName[ILen-1])
		// lenM:=len(Mimages)
		// for key,value:=Mimages{
		//            fmt.Println(Key,value)
		// }

		//del imgaes start

		//del immages over
	} // htcode=200 over

}

// func Delimages(delimagesName string) bool{

// }
