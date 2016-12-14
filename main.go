package main

import (
	//	"github.com/Jet/common"
	"github.com/astaxie/beego"
	"github.com/jetlwx/k8s_admin/crontab"
	"github.com/jetlwx/k8s_admin/models"
	_ "github.com/jetlwx/k8s_admin/routers"

	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func opensesion() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 14400
	sessionTime, err := strconv.ParseInt(beego.AppConfig.String("sessionSeconds"), 10, 64)
	if err != nil {
		sessionTime = 14400
	}
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = sessionTime
}

/* Global Variables */
//var log_T_memory, log_D_memory, log_I_memory, log_W_memory, log_C_memory string

func main() {
	opensesion()
	models.DBinit()
	//models.GetOneRecord(&models.KubeMasterSetting{})
	go crontab.StartCrontab()
	beego.Run()

}
