package virtualization

import (
	"strconv"
	"strings"

	"log"

	"github.com/jetlwx/k8s_admin/common"
)

//Can transfer ip statement to ip list
func SplitIPsToip(ips string, suffix string) (s []string, err error) {
	isip := common.IsIP(ips)
	if !isip {
		return nil, nil
	}
	str := strings.Split(ips, ".")
	prefix := str[0] + "." + str[1] + "." + str[2]

	lastNum, _ := strconv.Atoi(str[3])
	switch {
	case lastNum > 0 && suffix == "": //just an ip
		log.Println("doing 1......")
		s = append(s, ips)
		return s, nil

	case lastNum >= 0:
		log.Println("doing 2......")
		var start int
		var stu int
		if suffix != "" {
			stu, _ = strconv.Atoi(suffix)
		} else {
			stu = 254
		}

		if lastNum == 0 {
			start = 1
		} else {
			start = lastNum
		}
		for i := start; i <= stu; i++ {
			newip := prefix + "." + strconv.Itoa(i)
			s = append(s, newip)

		}
		return s, nil
	default:
		return nil, nil
	}
	return
}
