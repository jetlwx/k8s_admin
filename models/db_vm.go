package models

import (
	"log"

	"errors"
	"github.com/jetlwx/k8s_admin/common"
	"strconv"
)

func InsertIPPool(ip string) (int64, error, string) {
	pool := new(IpPool)
	pool.Ipaddr = ip
	host := new(Hosts)
	has1, _ := engine.Where("ipaddr= ?", ip).Get(host)
	has, err := engine.Where("ipaddr= ?", ip).Get(pool)
	if !has && !has1 {
		ok, err := engine.InsertOne(pool)
		if err == nil {
			return ok, nil, ""
		} else {
			log.Printf("插入入数据出错 %#v", err)
			common.Writelog("Critical", common.CustomerErr(err))
			return 0, err, ip
		}

	}
	return 0, err, ip
}

//Get a table all records of IpPool
func Getvirtualmachinelist() []IpPool {
	var machine []IpPool
	engine.Find(&machine)
	return machine
}

//得到记录总数
func GetTotalItems(conditions string) int64 {
	tab := new(IpPool)
	conditions = SqlConditionsFliter(conditions)
	if conditions != "" {
		conditions = "and (ipaddr like '%" + conditions + "%' or vm_name like '%" + conditions + "%' or vm_desc like '%" + conditions + "%')"
	}

	engine.ShowSQL()
	total, err := engine.Where("1>0 " + conditions).Count(tab)
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
	conditions = SqlConditionsFliter(conditions)
	if conditions != "" {
		conditions = "and (ipaddr like '%" + conditions + "%' or vm_name like '%" + conditions + "%' or vm_desc like '%" + conditions + "%' or hosts_ip like '%" + conditions + "%')"
	}
	log.Println("conditions-->", conditions)
	engine.ShowSQL()
	err := engine.Where("1>0 "+conditions).Desc("useing").Limit(pagesize, offset).Find(&pool)
	return pool, err

}

//get an recorder by ip_pool id
func GetIPPoolById(id int) (*IpPool, error) {
	info := new(IpPool)
	engine.ShowSQL()
	//err := engine.Where("id = ?", id).Limit(1, 0).Find(&info)
	_, err := engine.Where("id = ?", id).Get(info)
	return info, err
}

//update ip_pool by id
func UpdateIPPool(p IpPool) (int64, error) {
	h := new(IpPool)
	id := p.Id
	vmname := p.VmName
	vmdesc := p.VmDesc
	host := p.HostsIp
	engine.ShowSQL()
	c1 := "vm_name='" + vmname + "' and hosts_ip='" + host + "'"
	count, err2 := engine.Where(c1).Count(h)
	log.Println("count--->", count)
	if count == 0 {
		//	affected, err := engine.Exec("update ip_pool set  vm_name = ? ,vm_desc = ? where id = ?", vmname, vmdesc, id)
		affected, err := engine.Id(id).Cols("vm_name", "vm_desc", "hosts_ip").Update(&IpPool{VmName: vmname, VmDesc: vmdesc, HostsIp: host})

		if err != nil {
			return 0, err
		}
		return affected, nil
	}
	return 0, err2
}

//update ip_pool ,vm mac address
func UpdateVmIPMacaddr(p IpPool) (int64, error) {
	h := new(IpPool)
	vmid := strconv.FormatInt(p.Id, 10)
	vmmac := p.VmMac
	vmname := p.VmName

	sqlstr := "id=" + vmid + " " + "and vm_name=" + vmname + " " + "and vm_mac=" + vmmac
	count, err := engine.Where(sqlstr).Count(h)
	if count == 0 {
		affect, err := engine.Id(p.Id).Cols("vm_mac").Update(&IpPool{VmMac: vmmac})
		if err != nil {
			return 0, err
		}
		return affect, nil
	}
	return 0, err
}

//update ip_pool ,vm status
func UpdateVmvmstate(p IpPool) (int64, error) {
	h := new(IpPool)
	vmid := strconv.FormatInt(p.Id, 10)
	vmstate := p.VmState
	vmname := p.VmName

	sqlstr := "id=" + vmid + " " + "and vm_name=" + vmname + " " + "and vm_state=" + vmstate
	count, err := engine.Where(sqlstr).Count(h)
	if count == 0 {
		affect, err := engine.Id(p.Id).Cols("vm_state").Update(&IpPool{VmState: vmstate})
		if err != nil {
			return 0, err
		}
		return affect, nil
	}
	return 0, err
}

// get vm list in ippool where useing state [useing| nouseing|all]
func GetvmIPpoolByprefix(prefix string, usestate string) ([]IpPool, error) {
	var p []IpPool
	uCondition := ""
	switch usestate {
	case "useing":
		uCondition = " and useing = 1"
	case "nouseing":
		uCondition = " and useing = 0"
	}
	condition := "1> 0 and ipaddr like '" + prefix + "%' " + uCondition

	log.Println("condition-->", condition)
	err := engine.Where(condition).Find(&p)
	engine.ShowSQL()
	return p, err
}

//get vm list By condition
func GetvmIPlistByConditions(conditions string) ([]IpPool, error) {
	if len(conditions) <= 3 {
		return nil, errors.New("conditions is null ,Forbidden")
	}
	var p []IpPool
	condition := "1 >0 and" + " " + conditions
	err := engine.Where(condition).Find(&p)
	engine.ShowSQL()
	return p, err
}
