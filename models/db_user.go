package models

import (
	"errors"
	"log"
)

func ShowAllUser() ([]Users, error) {
	u := []Users{}
	engine.ShowSQL()
	err := engine.Desc("id").Find(&u)
	return u, err

}
func GetUserinfoById(id int64) (Users, error) {
	u := Users{}
	_, err := engine.Where("id=?", id).Get(&u)
	if err != nil {
		return u, err
	}
	return u, nil

}

//add user to DB
func AddUser(u *Users) error {
	username := u.LoginName
	password := u.Password
	if len(username) <= 2 {
		return errors.New("the length of username must ge 3")
	}
	if len(password) < 4 {
		return errors.New("the length of password must gt 4")
	}
	loginuser := new(Users)
	loginuser.LoginName = u.LoginName
	loginuser.Name = u.Name
	loginuser.Password = u.Password
	loginuser.Role = u.Role
	loginuser.State = u.State
	has, err := engine.Where("login_name=? and password=?", username, password).Get(loginuser)
	if err != nil {
		return err
	}
	if has {
		str := "user:" + username + "exist!!"
		return errors.New(str)
	}
	_, err2 := engine.InsertOne(loginuser)
	return err2
}

func LoginCheck(u *Users) (bool, *Users, error) {
	username := u.LoginName
	password := u.Password
	loginuser := new(Users)
	has, err := engine.Where("login_name=? and password=? and state=1", username, password).Get(loginuser)
	if err != nil {
		log.Println("login check: user=", username, "login faild")
		log.Println("login err log:", err)
	}
	return has, loginuser, err

}

func UpdateUser(u *Users) error {
	var sql string
	var err error

	id := u.Id
	name := u.Name
	pass := u.Password
	role := u.Role
	state := u.State
	if id == 0 {
		return errors.New("user id is null")
	}
	if pass == "" {
		sql = "update users set name = ? , role= ? ,state = ? where id= ?"
		_, err = engine.Exec(sql, name, role, state, id)
	} else {
		sql = "update users set name= ? ,password= ? ,role = ? ,state = ? where id = ?"
		_, err = engine.Exec(sql, name, pass, role, state, id)
	}

	engine.ShowSQL()
	//	log.Println("the user info has been update", res)
	return err
}

func DelUserById(id int64) error {
	sql := "delete from users where id= ?"
	_, err := engine.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}

func ModifyPuttyPath(U *Users) error {
	uid := U.Id
	ulogname := U.LoginName
	putty := U.Puttypath
	psftp := U.Psftppath
	upuser := new(Users)
	if putty == "" && psftp == "" {
		return errors.New("nothing to update")
	}
	has, err := engine.Where("login_name=? and id=? and state=1", ulogname, uid).Get(upuser)
	if err != nil {
		return err
	}
	if has {
		sql := "update users set puttypath = ? , psftppath = ? where id = ?"
		_, err2 := engine.Exec(sql, putty, psftp, uid)
		return err2
	}
	return errors.New("No user find")
}
