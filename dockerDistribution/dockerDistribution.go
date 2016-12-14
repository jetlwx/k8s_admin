package dockerDistribution

//package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/jetlwx/k8s_admin/common"
	"io/ioutil"
	"log"
	"net/http"
	//"reflect"
	"regexp"
	"strings"
)

type ImagesRepos struct {
	Repositories []string `json:"repositories"`
}
type ImageTags struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

//得到访问distribution API 访问结果
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
			common.Writelog("ACCESS", "获取url"+url+"出错")
			log.Fatalf("获取HTTPs内容出错 ", err)

		}
		common.Writelog("ACCESS", "成功获取url"+url+"内容")
		HttpStatusCode = resp.StatusCode
		defer resp.Body.Close()
		if HttpStatusCode == 200 {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				common.Writelog("ACCESS", "获取url "+url+"出错")
				log.Fatalf("读取HTTP流出错 ", err)

			}
			common.Writelog("ACCESS", "成功获取url "+url+"内容")
			res = body
			common.Writelog("ACCESS", "获取url "+url+"内容为"+string(res))
		}
		return HttpStatusCode, res

	case "http":
		resp, err := http.Get(url)
		if err != nil {
			common.Writelog("ACCESS", "获取url "+url+"出错")
			log.Fatalf("获取HTTP内容出错 ", err)

		}
		common.Writelog("Info", "成功获取url "+url+"内容")
		HttpStatusCode = resp.StatusCode
		defer resp.Body.Close()
		if HttpStatusCode == 200 {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				common.Writelog("ACCESS", "获取http流出错")
				log.Fatalf("读取HTTP流出错 ", err)

			}
			res = body
			common.Writelog("ACCESS", "获取url"+url+"内容为"+string(res))

		}

	}
	//fmt.Println("httpcode", HttpStatusCode)
	return HttpStatusCode, res

}

// 得到DISTRIBUTION所有镜像名,keyword可省
func GetImageNames(url string, keyword string) (HttpStatusCode int, res []string) {
	//imgListUrl := comm.GetfullUrl("imglist", url, "", "")  //del
	imgListUrl := common.GetImgListUrl(url, "")
	hcode, body := GetData(imgListUrl)
	ser := []string{}
	if hcode == 200 {
		var jsonImage ImagesRepos
		json.Unmarshal(body, &jsonImage)
		if len(keyword) != 0 {
			for _, i := range jsonImage.Repositories {
				t, _ := regexp.Match(keyword, []byte(i))
				if t {
					ser = append(ser, i)
				}
				//fmt.Println(i)
			}
			res = ser
		} else {
			res = jsonImage.Repositories
		}
	}
	common.Writelog("ACCESS", "调用函数  GetImageNames  成功！")
	return hcode, res
}

//得到指定镜像的TAGS,可能有多个
func GetImageTags(url string, img string) (HttpStatusCode int, res string) {
	//tagsUrl := comm.GetfullUrl("imgTags", url, img, "")  //del
	if len(img) > 0 {
		tagsUrl := common.GetImgTagsUrl(img, url)
		hcode, body := GetData(tagsUrl)
		if hcode == 200 {
			var jsonImage ImageTags
			json.Unmarshal(body, &jsonImage)
			restemp := jsonImage.Tags
			for k, i := range restemp {
				if k != 0 {
					res = res + ", " + i
				} else {
					res = i
				}
			}
		}
		common.Writelog("Debug", "call func 'GetImageTags' success!! ")
		return hcode, res
	}
	return
}

// //how to use
// func main() {

// 	httpcode, names := GetImageNames("https://172.16.6.142/v2/_catalog", "base")
// 	fmt.Println("GetImageNames====>", httpcode, names)
// 	// httpcode1, tags := GetImageTags("https://172.16.6.142","base/centos")
// 	// fmt.Println("GetImageTags===>", httpcode1, tags)
// }

// 字符串去重函数
func Rm_duplicate(list *[]string) []string {
	var x []string = []string{}
	for _, i := range *list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	return x
}

//得到指定镜像指定TAG的文件系统层 ,url 如 http://172.16.6.142:5000 或 https://172.16.6.142
func GetImagesfsLayers(imgName string, url string, imgTags string) (s1 []string, t error) {
	//myurl := url + "/v2/" + imageName + "/manifests/" + imagestag
	//myurl := comm.GetfullUrl("imgfsLayer", url, imgName, imgTags) // del
	myurl := common.GetImgFsLayerUrl(imgName, imgTags, url)
	hcode, body := GetData(myurl)
	if hcode == 200 {
		g, _ := simplejson.NewJson(body)
		g_body, _ := g.Get("fsLayers").Array()

		for _, value := range g_body {
			//	fmt.Println("key==", key, "value==", value)
			if v2, ok := value.(map[string]interface{}); ok { //value.(map[string]interface{})  就是断言上面value的类型
				//	fmt.Println("v2==", v2, "ok==", ok)
				for _, v := range v2 {
					//fmt.Println("v==", v.(string))
					s1 = append(s1, v.(string))

				}

			}
		}

	}
	return s1, nil

}

//得到准备删除镜像 imageName 的BLOBSUB
func GetdelimageBlobSums(imageName string, url string, imagestag string) (s2 []string, t bool) {
	b1, err := GetImagesfsLayers(imageName, url, imagestag)
	if err == nil {
		s2 = Rm_duplicate(&b1) //去重
		if len(s2) != 0 {
			t = true
		}
	}
	return s2, t

}

//在所有镜像中寻找是否有被依赖，若有依赖则返回被依赖的镜像名：TAG和TRUE

func FindRelationShipsBetweenImages(delimagesName string, delimagestag string, url string) (dependedImg string) {
	//得到准备删除IMG的BLOBSUMS
	delBlobSumList, _ := GetdelimageBlobSums(delimagesName, url, delimagestag)
	fmt.Println("delBlobSumList==", delBlobSumList)
	lenDelBlobSumList := len(delBlobSumList)
	if lenDelBlobSumList == 1 {
		return
	} else {
		htcode, imgList := GetImageNames(url, "")
		fmt.Println("imgList==", imgList)
		if htcode == 200 {
			for k, img := range imgList { // img列表
				fmt.Println("k==", k, "img==", img)
				if len(img) != 0 {
					htcode2, tags := GetImageTags(url, img) //img列表中对应的img的tag
					fmt.Println("tags==", tags)
					if htcode2 == 200 {
						tagone := strings.Split(tags, ",")
						fmt.Println("tagone", tagone)
						//mytag := tagone[0] //只取第一个进行判断
						for _, mytag := range tagone {
							fmt.Println("mytag==", mytag)
							//已得到img列表的imgname,tags，准备获取其 blogsum
							imgofListBlogSums, err := GetImagesfsLayers(img, url, mytag)
							fmt.Println("imgofListBlogSums==", imgofListBlogSums)
							if err == nil {
								for _, delblobs := range delBlobSumList {
									fmt.Println("delblobs==", delblobs)
									for _, imgblobs := range imgofListBlogSums {
										if delblobs == imgblobs {

											fmt.Println("imgblobs==", imgblobs)
											dependedImg = img + ":" + mytag
											goto JUSTOVER
										}
									}
								}

							}
						}

					}
				}
			}

		}

	}
JUSTOVER:
	return dependedImg
}
