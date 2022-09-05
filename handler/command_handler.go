package handler

import (
	"github.com/buzhiyun/aliyun-security-group/utils"
	"github.com/kataras/golog"
)

/*
该处理器只更新模板中的hosts文件，然后根据cmd 任务去 推送hosts 文件
*/
type commandHandler struct {
	isAvailable bool
	name        string
	cmdList     []cmd
}

/*
name 参数 第一个是 handle名字 ， 第二个是 是否windows文件 "true"|"false"
*/
func (h *commandHandler) NewHandler(name string, cmdList []cmd, args ...string) handler {
	h.name = name
	h.cmdList = cmdList
	h.isAvailable = true
	//if err := h.initConfig();err != nil{
	//	h.isAvailable = false
	//}
	return h
}

func (h *commandHandler) String() string {
	return h.name
}

func (h *commandHandler) Available() bool {
	return h.isAvailable
}

func (h *commandHandler) UpdateHost(domain, ip string) (err error) {
	for _, cmd := range h.cmdList {
		args := make([]string, len(cmd.Args))
		copy(args, cmd.Args)
		args[len(args)-1] = args[len(args)-1] + " -domain " + domain + " -ip " + ip
		result := utils.Exec(cmd.Command, args...)
		golog.Infof("执行cmd：\n%s", result)
	}
	return
}

//func (h *commandHandler)initConfig() (err error)  {
//	//h.cmdList , err = getHandleCmd(h.name)
//	return
//}
