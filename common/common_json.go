package common

import (
	//"github.com/Jet/models"   //for log 弃用
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//得到访问distribution API 访问结果
//请求地址类似于http://172.16.6.160:8080/api/v1/componentstatuses 或
//https://172.16.6.160/api/v1/componentstatuses
func GetData(url string) (HttpStatusCode int, res []byte) {
	httpOrhttps := strings.Split(url, ":")
	//fmt.Println(httpOrhttps[0])
	switch httpOrhttps[0] {
	case "https":
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		resp, err := client.Get(url)
		if err != nil {
			Writelog("ACCESS", "获取url"+url+"出错")
			log.Fatalf("获取HTTPs内容出错 ", err)

		}
		Writelog("ACCESS", "成功获取url"+url+"内容")
		HttpStatusCode = resp.StatusCode
		defer resp.Body.Close()
		if HttpStatusCode == 200 {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				Writelog("ACCESS", "获取url "+url+"出错")
				log.Fatalf("读取HTTP流出错 ", err)

			}
			Writelog("ACCESS", "成功获取url "+url+"内容")
			res = body
			Writelog("ACCESS", "获取url "+url+"内容为"+string(res))
		}
		return HttpStatusCode, res

	case "http":
		resp, err := http.Get(url)
		if err != nil {
			Writelog("ACCESS", "获取url "+url+"出错")
			log.Fatalf("获取HTTP内容出错 ", err)

		}
		Writelog("Info", "成功获取url "+url+"内容")
		HttpStatusCode = resp.StatusCode
		defer resp.Body.Close()
		if HttpStatusCode == 200 {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				Writelog("ACCESS", "获取http流出错")
				log.Fatalf("读取HTTP流出错 ", err)

			}
			res = body
			Writelog("ACCESS", "获取url"+url+"内容为"+string(res))

		}

	}
	//fmt.Println("httpcode", HttpStatusCode)
	return HttpStatusCode, res

}
