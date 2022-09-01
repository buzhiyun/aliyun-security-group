package handler

import (
	"github.com/buzhiyun/aliyun-security-group/hosts"
	"github.com/buzhiyun/aliyun-security-group/utils"
	"github.com/kataras/golog"
)

type szoneHandler struct {
	IsAvailable bool
	cmdList		[]cmd
}

func (h *szoneHandler)NewHandler(name string ,cmdList []cmd,args... string) handler {
	h.IsAvailable = true
	h.cmdList = cmdList
	//if err := h.initConfig();err != nil{
	//	h.IsAvailable = false
	//}
	return h
}

func (h *szoneHandler)String() string {
	return "szoneHandler"
}

func (h *szoneHandler)Available() bool {
	return h.IsAvailable
}


func (h *szoneHandler)UpdateHost(domain,ip string) (err error) {
	if err != nil { return }

	//golog.Debugf("替换 %s 到新的ip %s",domain,ip)
	changed , err := hosts.UpdateHosts("szone",domain,ip,true)
	if err != nil || !changed {	return }

	for _, cmd := range h.cmdList {
		result := utils.Exec(cmd.Command,cmd.Args...)
		golog.Infof("执行发送 hosts 文件：\n%s",result)
	}

	return
}

//func (h *szoneHandler)initConfig() (err error)  {
//
//	return
//}