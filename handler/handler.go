package handler

import (
	"github.com/buzhiyun/go-utils/cfg"
	"github.com/kataras/golog"
)

type handler interface {
	NewHandler(name string,cmdList []cmd, args... string) handler
	UpdateHost(domain,ip string) (err error)		// 更新主机上的dns信息
	String() string					// 获取当前handler名字
	Available() bool				// 是否初始化成功
}

var (
	HandlerMap = map[string]handler{}
)

type hadlerCfg struct {
	Type	string
	Args	[]string
	Cmd		[]cmd
}

func init() {
	golog.Debugf("初始化 handler")

	var handlerCfg map[string]hadlerCfg
	if ok :=cfg.Config().Scan("handler.handlers",&handlerCfg); !ok {
		golog.Fatal("读取配置异常，初始化 handler 失败")
	}

	for handlerName, handler := range handlerCfg {
		switch handler.Type {
		case "szone":
			_handler := szoneHandler{}
			HandlerMap[handlerName] = _handler.NewHandler("szone", handler.Cmd)
		case "hosts":
			_handler := hostsHandler{}
			HandlerMap[handlerName] = _handler.NewHandler(handlerName, handler.Cmd,handler.Args...)
		case "command":
			_handler := commandHandler{}
			HandlerMap[handlerName] = _handler.NewHandler(handlerName, handler.Cmd,handler.Args...)
		}
	}

	golog.Debugf("handler 配置 %#v",handlerCfg)


}
