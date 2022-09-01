package access

import (
	"github.com/kataras/golog"
	"net"
	"time"
)


func getIp(domain string) (ip []string, err error) {
	record , ok := resolvRecord[domain]
	if ok {
		return record, nil
	}

	ipSet := map[string]interface{}{}

	for i:= 0;i< 15 ; i++ {
		iplist , err := net.LookupHost(domain)
		if err != nil {
			golog.Errorf("域名解析错误, %s",err.Error())
			continue
		}
		for _, s := range iplist {
			ipSet[s] = nil
		}
		time.Sleep(time.Millisecond * 10)
	}
	//ip , err = net.LookupHost(domain)
	//if err != nil {
	//	golog.Errorf("域名解析错误, %s",err.Error())
	//	return
	//}
	for _ip, _ := range ipSet {
		ip = append(ip,_ip)
	}
	resolvRecord[domain] = ip
	golog.Debugf("解析到域名 %s 有以下ip: %s",domain,ip)
	return
}
