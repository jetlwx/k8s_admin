package common

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	//"strings"
	"bytes"
	"io/ioutil"
)

/* Global Variables */
// var Log_T_memory, Log_D_memory, Log_I_memory, Log_W_memory, Log_C_memory string
// Log_T_memory = beego.AppConfig.String("Trace_log")
// Log_D_memory = beego.AppConfig.String("Debug_log")
// Log_I_memory = beego.AppConfig.String("Info_log")
// Log_W_memory = beego.AppConfig.String("Warn_log")
// Log_C_memory = beego.AppConfig.String("Critical_log")
var (
	Log_T_memory = beego.AppConfig.String("Trace_log")
	Log_D_memory = beego.AppConfig.String("Debug_log")
	Log_I_memory = beego.AppConfig.String("Info_log")
	Log_W_memory = beego.AppConfig.String("Warn_log")
	Log_C_memory = beego.AppConfig.String("Critical_log")
	Log_A_memory = beego.AppConfig.String("Access_log")
	Log_E_memory = beego.AppConfig.String("Error_log")
)

//返回distribution访问路径
func GetImgListUrl(url string, other string) (fullUrl string) {
	return url + "/v2/_catalog" + other
}

//返回镜像TARG列表 URL
func GetImgTagsUrl(imgName string, url string) (fullUrl string) {
	return url + "/v2/" + imgName + "/tags/list"
}

//返回指定镜像TAG URL
func GetImgFsLayerUrl(imgName string, imgTags string, url string) (fullUrl string) {
	return url + "/v2/" + imgName + "/manifests/" + imgTags
}

//返回kubemaster api路径
func GetKubeMasterApi(kbm_url string, kbm_protocol string, kbm_port int64, conten string) (fullurl string) {
	//int64 to string
	//strconv.FormatInt(kbm_port, 10)
	re_url := kbm_protocol + "://" + kbm_url + ":" + strconv.FormatInt(kbm_port, 10) + "/api/v1/" + conten
	fmt.Println("common.re_url=", re_url)
	return re_url
}

//返回日志存储名与路径
func Getlogpath(logtype string) (j string) {
	rotateName := logtype + fmt.Sprintf(".%s.%03d", time.Now().Format("2006-01-02"), 1) + ".log"
	logj := map[string]string{}
	logj["filename"] = beego.AppConfig.String("beegologdir") + rotateName
	loglog, _ := json.Marshal(&logj)
	//fmt.Println("logj[filename]===", string(loglog))
	return string(loglog)

}

func Writelog(log_level string, Logs ...string) {
	var log_type string

	if log_level == "ACCESS" || log_level == "A" {
		log_type = "Access_log"
	} else {
		log_type = "Running_log"
	}

	//level 日志级别
	var logshow string
	log := logs.NewLogger(100)

	//异步输出
	log.Async()
	log.SetLogger("file", Getlogpath(log_type))
	for i, _ := range Logs {
		logshow = logshow + " " + Logs[i]
	}
	//fmt.Println("Log_C_memory", Log_C_memory)
	switch log_level {
	case "Debug", "D":
		if Log_D_memory == "on" {
			log.Debug(logshow)
		}
	case "Info", "I":
		if Log_I_memory == "on" {
			log.Info(logshow)
		}
	case "Trace", "T":
		if Log_T_memory == "on" {
			log.Trace(logshow)
		}
	case "Warn", "W":
		if Log_W_memory == "on" {
			log.Warn(logshow)
		}
	case "Error", "E":
		if Log_E_memory == "on" {
			log.Error(logshow)
		}
	case "Critical", "C":
		if Log_C_memory == "on" {
			log.Critical(logshow)
		}
	case "ACCESS", "A":
		if Log_A_memory == "on" {
			log.Critical(logshow)
		}

	}
}

//error类型转string，方便重定向输出
func CustomerErr(err error) string {
	if err != nil {
		return fmt.Sprintf("%s", err.Error())
	}
	return ""
}

//返回几天几个小时
func GetDate(H float64) (Days int64, Hour int64) {
	Hours := float64(math.Remainder(H, 24))
	if Hours < 0 {
		Hours = 24.00 + Hours
	}
	Days = int64(H / 24)
	return Days, int64(Hours)
}

// 返回xxxDxxxH格式
func TimetoDayandHour(t0 string) (DaysandHours string) {
	tm2, _ := time.Parse("2006-01-02T15:04:05Z", t0)
	t1 := time.Now()
	d_sub := t1.Sub(tm2).Hours()
	kbm_days_s, kbm_hours_s := GetDate(d_sub)
	kbm_days := strconv.FormatInt(kbm_days_s, 10)
	kbm_hours := strconv.FormatInt(kbm_hours_s, 10)
	return kbm_days + "d" + kbm_hours + "h"
}

func PostJsonToApiServer(url string, parm string, bodystr string) (code int, reson string) {
	end := []byte(bodystr)
	body := bytes.NewBuffer(end)
	res, err := http.Post(url, parm, body)
	if err != nil {
		Writelog("Warn", "POST出错："+CustomerErr(err))
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	j, _ := gabs.ParseJSON(result)
	var err1 error
	code, err1 = strconv.Atoi(j.S("code").String())
	if err1 != nil {
		code = 0
	}
	status := j.S("status").String()
	r1 := j.S("reason").String()
	msg := j.S("message").String()
	reson = "status: " + status + "reson: " + r1 + "message:" + msg

	if err != nil {
		log.Fatal(err)
		return
	}

	//fmt.Printf("%s", result)
	return code, reson

}

//计算机单位自动匹配

const (
	_        = iota // iota = 0
	KB int64 = 1 << (10 * iota)
	MB       //iota=1
	GB       // 与 KB 表达式相同,但 iota = 2
	TB
	PB
	YB
)

//check ip
func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}

func TransferByte(size int64) (val string) {
	aB := size
	aKB := (size / KB)
	aMB := (size / MB)
	aGB := (size / GB)
	aTB := (size / TB)
	aPB := (size / PB)
	aYB := (size / YB)
	f := int64(1024)
	var res int64
	var dw string

	if aB > f {
		if aKB > f {
			if aMB > f {
				if aGB > f {
					if aTB > f {
						if aPB > f {
							res = aYB
							dw = "YB"
						} else {
							res = aPB
							dw = "PB"
						}
					} else {
						res = aTB
						dw = "TB"
					}
				} else {
					res = aGB
					dw = "GB"
				}
			} else {
				res = aMB
				dw = "MB"
			}
		} else {
			res = aKB
			dw = "KB"
		}
	} else {
		res = aB
		dw = "Byte"
	}

	return strconv.FormatInt(res, 10) + dw
}

//[]byte to string
func ByteArrToString(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, ",")
}

//check the url para is int or int64
func UrlParaIsInt(s string) (int, error) {
	b, _ := regexp.MatchString("[0-9].*", s)
	if b {
		m, err := strconv.Atoi(s)
		return m, err
	}
	return 0, nil
}

//check the url para is int
func UrlParaIsInt2(s string) (bool, error) {
	b, _ := regexp.MatchString("[0-9].*", s)
	if b {
		_, err := strconv.Atoi(s)
		return true, err
	}
	return false, nil
}

// 随机字符串
// KC_RAND_KIND_NUM   = 0  // 纯数字
// KC_RAND_KIND_LOWER = 1  // 小写字母
// KC_RAND_KIND_UPPER = 2  // 大写字母
// KC_RAND_KIND_ALL   = 3  // 数字、大小写字母
//生成随机字符串
func RandomString() string {
	size := 16
	kind := 3
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

//fliter the specia charset
func FliterSpecChars(s string) string {
	log.Println("before fliter:", s)
	s = strings.Replace(s, "\\ ", "", -1)
	s = strings.Replace(s, "=", "", -1)
	s = strings.Replace(s, "^", "", -1)
	s = strings.Replace(s, "*", "", -1)
	s = strings.Replace(s, "^", "", -1)
	s = strings.Replace(s, "!", "", -1)
	s = strings.Replace(s, "+", "", -1)
	s = strings.Replace(s, "-", "", -1)
	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	s = strings.Replace(s, "&", "", -1)
	s = strings.Replace(s, "$", "", -1)
	s = strings.Replace(s, "#", "", -1)
	s = strings.Replace(s, "@", "", -1)
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, "\\.", "", -1)
	log.Println("after fliter:", s)
	return s
}
