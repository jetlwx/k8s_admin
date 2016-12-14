package account

import (
	"github.com/astaxie/beego"
	"github.com/jetlwx/k8s_admin/common"
	"github.com/jetlwx/k8s_admin/models"
	"log"
)

//AccountController is for user controll and priviles
type AccountController struct {
	beego.Controller
}

func (ac *AccountController) Get() {
	ac.TplName = "usersCenter/login.html"
	if ac.GetString("logout") == "1" {
		ac.DelSession("LoginName")
		ac.DelSession("Name")
		ac.DelSession("Role")
	}
}

//check user
func (ac *AccountController) Post() {
	UserName := ac.Input().Get("UserName")
	Password := ac.Input().Get("Password")
	n := &models.Users{}
	n.LoginName = common.FliterSpecChars(UserName)
	n.Password = common.FliterSpecChars(Password)
	has, userinfo, err := models.LoginCheck(n)
	if err != nil {
		ac.Ctx.WriteString(common.CustomerErr(err))
		return
	}
	if has == false {
		ac.Redirect("/login", 302)
		return
	}
	log.Println("user", userinfo)
	// if !models.Setsession(ac.Ctx, userinfo) {
	// 	ac.Ctx.WriteString("Session set error")
	// 	return
	// }
	ac.SetSession("Id", userinfo.Id)
	ac.SetSession("LoginName", userinfo.LoginName)
	ac.SetSession("Name", userinfo.Name)
	ac.SetSession("Role", userinfo.Role)
	ac.Redirect("/", 302)

}

func Checklogin(c *AccountController) bool {
	ck := c.GetSession("LoginName")
	if ck != nil {
		return true
	}
	return false
}
