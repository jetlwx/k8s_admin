package models

import (
	//"encoding/json"
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jetlwx/k8s_admin/common"
)

//flit speace charsert
func SqlConditionsFliter(conditions string) string {
	conditions = strings.Replace(conditions, "'", "", -1)
	conditions = strings.Replace(conditions, "!", "", -1)
	conditions = strings.Replace(conditions, "#", "", -1)
	conditions = strings.Replace(conditions, "$", "", -1)
	conditions = strings.Replace(conditions, "%", "", -1)
	conditions = strings.Replace(conditions, "^", "", -1)
	conditions = strings.Replace(conditions, "&", "", -1)
	conditions = strings.Replace(conditions, "*", "", -1)
	conditions = strings.Replace(conditions, "(", "", -1)
	conditions = strings.Replace(conditions, ")", "", -1)
	conditions = strings.Replace(conditions, "]", "", -1)
	conditions = strings.Replace(conditions, "[", "", -1)
	conditions = strings.Replace(conditions, "<", "", -1)
	conditions = strings.Replace(conditions, ">", "", -1)
	conditions = strings.Replace(conditions, "/", "", -1)
	conditions = strings.Replace(conditions, "\\", "", -1)
	return conditions
}

//从数据库中读一条记录
func GetOneRecord(bean interface{}) (interface{}, error) {
	a := bean
	fmt.Sprintf("dddddd %v#", a)
	//	log.Println("查询了表")
	common.Writelog("Trace", "查询了表")
	has, err := engine.Get(a)
	if err != nil {
		return false, err
	} else if !has {
		//log.Println("在表中没有找到记录！！！")
		common.Writelog("Trace", "在表中没有找到记录！！！")
		common.Writelog("Debug", common.CustomerErr(err))
		return nil, err
	}
	return a, nil
}

//插入kubemaster配置
func AddKuberMasterSetting(kbm_url string, kbm_port int64, kbm_protocol string) (e error) {
	//接收传入的值
	bean := &KubeMasterSetting{
		KubeMasterIp:       kbm_url,
		KubeMasterPort:     kbm_port,
		KubeMasterProtocol: kbm_protocol}
	//定义一个结构体表达式
	bean1 := &KubeMasterSetting{}
	//查询表内是否记录select count(*) from ....
	//有则更新，无则添加
	has, _ := engine.Count(bean1)
	if has == 0 {
		_, err := engine.InsertOne(bean)
		e = err
	} else {
		_, err := engine.Update(bean)
		e = err
	}
	return e
}

//获取kubemaster设置
func GetKubeMasterSetting() (*KubeMasterSetting, error) {
	bean := &KubeMasterSetting{}
	has, err := engine.Get(bean)
	if err != nil {
		common.Writelog("ACCESS", common.CustomerErr(err))
		return nil, err
	} else if !has {
		return nil, errors.New("记录未找到！")
	}
	return bean, nil
}

//直接获取kubemaster设置详细信息
func GetKubeMasterSettingDetail() (kbm_url string, kbm_protocol string, kbm_port int64, err error) {
	bean := &KubeMasterSetting{}
	has, err := engine.Get(bean)
	if err != nil {
		common.Writelog("ACCESS", common.CustomerErr(err))
		return "", "", int64(0), err
	} else if !has {
		return "", "", int64(0), errors.New("记录未找到！")
	}
	kbm_protocol = bean.KubeMasterProtocol
	kbm_url = bean.KubeMasterIp
	kbm_port = bean.KubeMasterPort

	return kbm_url, kbm_protocol, kbm_port, nil
}

// 添加distribution设置
func AddDistributionSetting(distribution_url string, distribution_protocol string) (e error) {
	bean := &DistributionSetting{
		DistributionUrl:      distribution_url,
		DistributionProtocol: distribution_protocol}
	bean1 := &DistributionSetting{}
	has, _ := engine.Count(bean1)
	if has == 0 {
		_, err := engine.InsertOne(bean)
		e = err
	} else {
		_, err := engine.Update(bean)
		e = err
	}
	return e
}

//获取distribution 设置
func GetDistributionSetting() (*DistributionSetting, error) {
	bean := &DistributionSetting{}
	has, err := engine.Get(bean)
	if err != nil {
		common.Writelog("ACCESS", common.CustomerErr(err))
		return nil, err
	} else if !has {
		return nil, errors.New("记录未找到！")
	}
	return bean, nil
}

// //插入表
// func InsertOneRecord(bean interface{}, col interface{}) (int64, error) {
// 	log.Println(col)
// 	has, e := engine.Id(ipaddr).Get(bean)
// 	log.Println("has===", has)
// 	if !has {
// 		ok, err := engine.InsertOne(bean)
// 		if err == nil {
// 			return ok, nil
// 		}
// 		log.Printf("插入入数据出错 %#v", err)
// 		common.Writelog("Critical", common.CustomerErr(err))
// 		return 0, err

// 	}
// 	return 0, e
// }
