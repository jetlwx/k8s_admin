package account

import (
	"github.com/astaxie/beego"
	"github.com/jetlwx/k8s_admin/common"
	"github.com/jetlwx/k8s_admin/models"
	"log"
	//	"os"
	"os/exec"
	"strconv"
)

type UsersCenterController struct {
	beego.Controller
}

func (c *UsersCenterController) UserList() {
	if c.GetSession("LoginName") == nil {
		c.Redirect("/login", 302)
	}
	c.TplName = "usersCenter/userCenter.html"
	c.Data["IsuserList"] = true
	c.Data["Isadduser"] = false
	c.Data["IssetSSHclient"] = false
	c.Data["IssetPrivileges"] = false
	userlist, err := models.ShowAllUser()
	if err != nil {
		c.Ctx.WriteString(common.CustomerErr(err))
		return
	}
	c.Data["Userlist"] = userlist
	return
}
func (c *UsersCenterController) AddUser() {
	if c.GetSession("LoginName") == nil {
		c.Redirect("/login", 302)
	}
	c.TplName = "usersCenter/addUser.html"
	c.Data["IsuserList"] = false
	c.Data["Isadduser"] = true
	c.Data["IssetSSHclient"] = false
	c.Data["IssetPrivileges"] = false
}

func (c *UsersCenterController) Post() {
	var state bool
	name := common.FliterSpecChars(c.Input().Get("name"))
	loginname := common.FliterSpecChars(c.Input().Get("loginname"))
	password := common.FliterSpecChars(c.Input().Get("password1"))
	s := c.Input().Get("state")
	if s == "1" {
		state = true
	} else {
		state = false
	}

	user := &models.Users{}
	user.LoginName = loginname
	user.Password = password
	user.Name = name
	user.Role = c.Input().Get("role")
	user.State = state

	err := models.AddUser(user)
	if err != nil {
		c.Ctx.WriteString(common.CustomerErr(err))
		return
	}
	c.Redirect("/UsersCenter/UserList", 302)
}

func (c *UsersCenterController) EditUser() {
	var seesionuid2 int64
	var urole2 string
	c.TplName = "usersCenter/editUser.html"
	uid := c.Input().Get("id")
	uid2, err := strconv.ParseInt(uid, 10, 64)

	seesionuid := c.GetSession("Id")
	if seesionuid != nil {
		seesionuid2 = seesionuid.(int64)
	}

	urole := c.GetSession("Role")
	if urole != nil {
		urole2 = urole.(string)
	}
	log.Println("urole2:", urole2)
	if (urole2 == "user" || urole2 == "") && uid2 != seesionuid2 {
		c.Ctx.WriteString("what do you want?")
		return
	}

	if err != nil {
		c.Ctx.WriteString(common.CustomerErr(err))
	}

	u, err2 := models.GetUserinfoById(uid2)
	if err2 != nil {
		c.Ctx.WriteString(common.CustomerErr(err2))
	}
	c.Data["User"] = u

}

func (c *UsersCenterController) UpdateUserInfo() {
	var S1 bool
	N := c.Input().Get("name")
	I := c.Input().Get("id")
	P := c.Input().Get("password")
	R := c.Input().Get("role")
	S := c.Input().Get("state")
	log.Println("N-->", N, "S-->", S)
	int64I, err := strconv.ParseInt(I, 10, 64)
	if err != nil {
		c.Ctx.WriteString(common.CustomerErr(err))
		return
	}
	if S == "true" {
		S1 = true
	}
	u := &models.Users{}
	u.Name = N
	u.Password = P
	u.Role = R
	u.Id = int64I
	u.State = S1

	e := models.UpdateUser(u)
	if e != nil {
		c.Ctx.WriteString(common.CustomerErr(e))
		return
	}
	c.Redirect("EditUser?id="+I, 302)
	return
}

func (c *UsersCenterController) DelUser() {
	var role1 string
	var idint int64
	var e error
	id := c.Input().Get("Id")
	role := c.GetSession("Role")
	log.Println("id", id)
	log.Println("role", role)
	idint, e = strconv.ParseInt(id, 10, 64)
	if e != nil {
		c.Ctx.WriteString("User id is inaviable")
		return
	}

	if role != nil {
		role1 = role.(string)
	}

	if role1 != "admin" {
		c.Ctx.WriteString("forbidden for role user")
		return
	}

	err := models.DelUserById(idint)
	if err != nil {
		c.Ctx.WriteString(common.CustomerErr(err))
		return
	}
	c.Redirect("/UsersCenter/UserList", 302)
	return
}

func (c *UsersCenterController) OpenSFTP() {
	c.TplName = "virtualization/machinelist.html"
	// var attr *os.ProcAttr
	// var st = make([]string, 2)
	// st[0] = "-u"
	// st[1] = "root"
	// proc, err := os.StartProcess("C:\\Users\\Jet\\Downloads\\putty_gr\\psftp.exe", st, attr)
	// log.Println("err---->", err.
	// proc.Wait()
	cmd := exec.Command("cmd.exe", "/c", "start C:\\Users\\Jet\\Downloads\\putty_gr\\psftp.exe ")
	cmd.Run()
	cmd.Wait()

}

func (c *UsersCenterController) SSHclient() {
	c.TplName = "usersCenter/sshclient.html"
	seesionuid := c.GetSession("Id")
	if seesionuid == nil {
		//c.Ctx.WriteString(" you must login first")
		c.Redirect("/login", 302)
	}

	sid := seesionuid.(int64)
	user, err := models.GetUserinfoById(sid)
	if err != nil {
		c.Ctx.WriteString(common.CustomerErr(err))
		return
	}
	c.Data["PUTTY"] = user.Puttypath
	c.Data["PSFTP"] = user.Psftppath

}
func (c *UsersCenterController) SSHclientUp() {
	c.TplName = "usersCenter/sshclient.html"
	var sessuser string
	putty := c.Input().Get("putty")
	psftp := c.Input().Get("psftp")
	seesionuid := c.GetSession("Id")
	sessusername := c.GetSession("LoginName")
	if seesionuid == nil {
		c.Ctx.WriteString(" you must login first")
		return
	}
	if sessusername != nil {
		sessuser = sessusername.(string)
	}

	sid := seesionuid.(int64)
	if putty != "" || psftp != "" {
		u := &models.Users{}
		u.Id = sid
		u.LoginName = sessuser
		u.Puttypath = putty
		u.Psftppath = psftp

		err2 := models.ModifyPuttyPath(u)
		if err2 != nil {
			c.Ctx.WriteString(common.CustomerErr(err2))
			return
		}
		c.Redirect("/UsersCenter/SSHclient", 302)
	}
}
