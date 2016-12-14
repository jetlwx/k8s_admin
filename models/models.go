package models

import (
	"fmt"
	"log"
	"strconv"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/jetlwx/k8s_admin/common"
)

//日志级别设置(弃用)
// type Jlogs struct {
// 	Id            int64  `xorm:"notnull autoincr"`
// 	LevelTrace    string `xorm:"varchar(5) notnull default 'off'"`
// 	LevelDebug    string `xorm:"varchar(5) notnull default 'off'"`
// 	LevelInfo     string `xorm:"varchar(5) notnull default 'on'"`
// 	LevelWarn     string `xorm:"varchar(5) notnull default 'on'"`
// 	LevelCritical string `xorm:"varchar(5) notnull default 'on'"`
// }

//kubeMaster设置
type KubeMasterSetting struct {
	Id                 int64  `xorm:"notnull autoincr"`
	KubeMasterIp       string `xorm:"varchar(50)"`
	KubeMasterPort     int64  `xorm:"notnull default 8080"`
	KubeMasterProtocol string `xorm:"varchar(8) notnull default 'http'"`
}

// users and privileges
type Users struct {
	Id        int64  `xorm:"notnull autoincr"`
	Name      string `xorm:"varchar(10)"`
	LoginName string `xorm:"varchar(10) notnull"`
	Password  string `xorm:"varchar(50) notnull"`
	State     bool   `xorm:"default 0"`
	Role      string `xorm:"notnull default 'user'"`
	Puttypath string `xorm:"varchar(100)"`
	Psftppath string `xorm:"varchar(100)"`
}

//distribution数据仓库设置
type DistributionSetting struct {
	Id                   int64  `xorm:"notnull autoincr"`
	DistributionUrl      string `xorm:"varchar(255)"`
	DistributionProtocol string `xorm:"varchar(8) notnull default 'https'"`
}

// IpPool ip pool
type IpPool struct {
	Id          int64  `xorm:"notnull autoincr"`
	Ipaddr      string `xorm:"varchar(15)"`
	Useing      bool   `xorm:"notnull bool default false"`
	VmName      string `xorm:"varchar(255)"`
	VmState     string `xorm:" varchar(10) default 'Down'"`
	VmDesc      string `xorm:"varchar(255)"`
	HostsIp     string `xorm:"varchar(15)"`
	VmMac       string `xorm:"varchar(20)"`
	Sshuser     string `xorm:"varchar(20) default 'root'"`
	Sshport     int64  `xorm:"default 22"`
	Sshpassword string `xorm:"varchar(30)"`
}

//Hosts ,means virtual machine mother
type Hosts struct {
	Id                 int64  `xorm:"notnull autoincr"`
	Ipaddr             string `xorm:"varchar(15)"`
	State              string `xorm:" varchar(10) default 'Down'"`
	Location           string `xorm:"varchar(30)"`
	CpuMode            string `xorm:"varchar(10)"`
	Cpus               int    `xorm:"int"`
	CpuFrequency       string `xorm:"varchar(10)"`
	CpuSockets         int    `xorm:"int"`
	CoresPeerCpu       int    `xorm:"int"`
	MemorySizeKB       uint32 `xorm:"int"`
	Hostname           string `xorm:"varchar(30)"`
	MaxVcpus           int    `xorm:"int"`
	StorageCapacity    string `xorm:"varchar(30)"`
	StorageAllocation  string `xorm:"varchar(30)"`
	StorageAvailable   string `xorm:"varchar(30)"`
	NumberOfDomains    int    `xorm:"int"`
	WebsocketAgentPort string `xorm:"varchar(5)"`
	ImageTemplatePath1 string `xorm:"varchar(100)"`
	XmlTemplatePath1   string `xorm:"varchar(100)"`
	ImageTemplatePath2 string `xorm:"varchar(100)"`
	XmlTemplatePath2   string `xorm:"varchar(100)"`
	ImageTemplatePath3 string `xorm:"varchar(100)"`
	XmlTemplatePath3   string `xorm:"varchar(100)"`
	ImageTemplatePath4 string `xorm:"varchar(100)"`
	XmlTemplatePath4   string `xorm:"varchar(100)"`
	ImageTemplatePath5 string `xorm:"varchar(100)"`
	XmlTemplatePath5   string `xorm:"varchar(100)"`
}

//---------------------------------------------------------------
var engine *xorm.Engine

// 初始化数据库
func DBinit() {
	var err error
	dbuser := beego.AppConfig.String("mysqluser")
	dbpass := beego.AppConfig.String("mysqlpassword")
	dbhost := beego.AppConfig.String("mysqlhost")
	dbname := beego.AppConfig.String("mysqldb")
	dbport := beego.AppConfig.String("mysqlport")
	dbsource := dbuser + ":" + dbpass + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"

	engine, err = xorm.NewEngine("mysql", dbsource)
	if err != nil {
		common.Writelog("Critical", "创建数据库引擎失败", common.CustomerErr(err))
		log.Fatalf("创建数据库引擎失败 %v", err)

	}
	//是否显示打印SQL信息
	b, _ := strconv.ParseBool(beego.AppConfig.String("showsqllog"))
	if b {
		engine.ShowSQL(b)
	}
	//数据库PING测试
	if engine.Ping() == nil {
		fmt.Println("数据库" + dbhost + ":" + dbport + "连接成功！！！")
		common.Writelog("Info", "数据库"+dbhost+":"+dbport+"连接成功！！！")
	}

	//同步至数据库
	//注意：增加表后，要增加new(表名)
	err = engine.Sync2(new(KubeMasterSetting),
		new(DistributionSetting),
		new(IpPool),
		new(Hosts),
		new(Users))
	if err != nil {
		common.Writelog("Critical", "同步数据结构至数据库失败,"+common.CustomerErr(err))
		log.Fatalf("同步数据结构至数据库失败，可能原因为：%v ", err)

	}
}
