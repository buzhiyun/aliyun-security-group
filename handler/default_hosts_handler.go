package handler

import (
	"github.com/buzhiyun/aliyun-security-group/hosts"
	"github.com/buzhiyun/aliyun-security-group/utils"
	"github.com/kataras/golog"
)

/*
该处理器只根据cmd 任务去执行
*/
type hostsHandler struct {
	isAvailable bool
	name 		string
	cmdList		[]cmd
	isWindows	bool
}


/*
name 参数 第一个是 handle名字 ， 第二个是 是否windows文件 "true"|"false"
*/
func (h *hostsHandler)NewHandler(name string,cmdList []cmd,args... string) handler {
	h.name = name
	h.cmdList = cmdList
	h.isWindows = false
	if len(args) > 0 && args[0] == "true"{
		h.isWindows = true
	}


	h.isAvailable = true
	//if err := h.initConfig();err != nil{
	//	h.isAvailable = false
	//}
	return h
}

func (h *hostsHandler)String() string {
	return h.name
}

func (h *hostsHandler)Available() bool {
	return h.isAvailable
}

func (h *hostsHandler)UpdateHost(domain,ip string) (err error) {
	//cmdList , err := getHandleCmd(h.name)
	//if err != nil { return }

	//golog.Debugf("替换 %s 到新的ip %s",domain,ip)
	changed , err := hosts.UpdateHosts(h.name,domain,ip,h.isWindows)
	if err != nil || !changed {	return }

	for _, cmd := range h.cmdList {
		result := utils.Exec(cmd.Command,cmd.Args...)
		golog.Infof("执行发送 hosts 文件：\n%s",result)
	}
	return
}

//func (h *hostsHandler)initConfig() (err error)  {
//	//h.cmdList , err = getHandleCmd(h.name)
//	return
//}