package models

import (
	"errors"
	"log"

	"github.com/jetlwx/k8s_admin/common"
)

//add machine host mother
func AddHosts(ip string, location string) (int64, string, error) {
	host := new(Hosts)
	host.Ipaddr = ip
	host.Location = location
	pool := new(IpPool)
	host.State = "Down"
	//check the ip is exist in vm ippool
	has1, _ := engine.Where("ipaddr= ?", ip).Get(pool)
	has, err := engine.Where("ipaddr= ?", ip).Get(host)
	if !has && !has1 {
		ok, err := engine.InsertOne(host)
		if err == nil {
			return ok, "", nil
		} else {
			log.Printf("插入入数据出错 %#v", err)
			common.Writelog("Critical", common.CustomerErr(err))
			return 0, ip, err
		}

	}
	return 0, ip, err
}

//show all Hosts
func ShowAllHost() ([]Hosts, error) {
	host := []Hosts{}
	engine.ShowSQL()
	err := engine.Desc("id").Find(&host)
	return host, err

}

//get host list by ip prefix ,ex: 192.168.2
func GetHostlistByPrefix(conditions string) ([]Hosts, error) {
	h := []Hosts{}
	sqlwhere := "image_template_path1 <> '' and xml_template_path1 <> '' and ipaddr like '" + conditions + "%'"
	engine.ShowSQL()
	err := engine.Where(sqlwhere).Find(&h)
	return h, err
}

//Return the host infomation by host id
func GetHostInfo(hostid int64) (*Hosts, error) {
	h := new(Hosts)
	_, err := engine.Where("id=?", hostid).Get(h)
	engine.ShowSQL()
	return h, err
}

//retrun the host information by host ip
func GetHostInfoByIp(hostip string) (*Hosts, error) {
	h := new(Hosts)
	_, err := engine.Where("ipaddr=?", hostip).Get(h)
	engine.ShowSQL()
	return h, err
}

//edit host infomain
func EditHostinfo(h *Hosts) error {
	a := &Hosts{}
	has, err := engine.Id(h.Id).Get(a)
	if !has {
		return err
	}
	a.WebsocketAgentPort = h.WebsocketAgentPort
	a.ImageTemplatePath1 = h.ImageTemplatePath1
	a.XmlTemplatePath1 = h.XmlTemplatePath1

	a.ImageTemplatePath2 = h.ImageTemplatePath2
	a.XmlTemplatePath2 = h.XmlTemplatePath2

	a.ImageTemplatePath3 = h.ImageTemplatePath3
	a.XmlTemplatePath3 = h.XmlTemplatePath3

	a.ImageTemplatePath4 = h.ImageTemplatePath4
	a.XmlTemplatePath4 = h.XmlTemplatePath4

	a.ImageTemplatePath5 = h.ImageTemplatePath5
	a.XmlTemplatePath5 = h.XmlTemplatePath5
	_, err1 := engine.Id(h.Id).Update(a)
	return err1
}

//delete host
func DelHost(ip string, id int64) error {
	host := new(Hosts)
	pool := new(IpPool)

	//check ip in pool
	has1, _ := engine.Where("hosts_ip= ?", ip).Get(pool)
	has, _ := engine.Where("ipaddr= ? and id = ?", ip, id).Get(host)
	engine.ShowSQL()
	if has1 {
		return errors.New("The delete IP has in useing at  vm pool")
	}
	if has && !has1 {
		_, err := engine.Delete(&Hosts{Id: id, Ipaddr: ip})
		if err != nil {
			return err
		}

	}
	return nil
}

//Update ip pool's ip infomation after createvm success
func ChangeIPpoolAfterCreateVm(vmip string, vmname string, hostip string) error {
	h := new(IpPool)
	h.Ipaddr = vmip
	h.Useing = false
	has, err := engine.Get(h)
	engine.ShowSQL()
	if !has {
		return err
	}
	sql := "update ip_pool set useing=1,vm_name= ?,hosts_ip= ? where ipaddr=?"
	res, err := engine.Exec(sql, vmname, hostip, vmip)
	engine.ShowSQL()
	log.Println("the res of ChangeIPpoolAfterCreateVm ", res)
	return err
}
