package main

import (
	"flag"
	"github.com/buzhiyun/aliyun-security-group/hosts"
	"github.com/kataras/golog"
	"os"
)

func main() {
	domain := flag.String("domain","","要修改的域名")
	ip := flag.String("ip","","要修改的IP")

	flag.Parse()

	if len(*domain) == 0 || len(*ip) == 0 {
		golog.Fatalf("域名或IP 输入不正确")
		return
	}

	err := hosts.UpdateFile("/etc/hosts",*domain,*ip,false)
	if err != nil {
		return
	}
	err =os.Rename("/etc/hosts.new","/etc/hosts")

	if err != nil {
		golog.Errorf("更新 /etc/hosts 文件错误, %s",err.Error())
		return
	}

	golog.Infof("更新 hosts 成功 ，【 %s 】→【 %s 】",*domain,*ip)

}