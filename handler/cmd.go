package handler

type cmd struct {
	Command		string
	Args	[]string
}

//func getHandleCmd(handlerName string) (cmds []cmd, err error) {
//	commands ,ok := cfg.Config().Get("handler." + handlerName + ".cmd")
//	if !ok {
//		err = errors.New("获取配置 handler." + handlerName + ".cmd 失败")
//		golog.Error(err.Error())
//		return
//	}
//
//	cmdList ,ok := commands.([]interface{})
//	if !ok {
//		err = errors.New("解析配置 handler." + handlerName + ".cmd 失败")
//		golog.Error(err.Error())
//		return
//	}
//
//	for _, c := range cmdList {
//		c2 ,ok := c.(map[string]interface {})
//		if !ok {
//			err = errors.New("解析配置 handler." + handlerName + ".cmd 失败")
//			golog.Error(err.Error())
//			return
//		}
//
//		var _cmd = cmd{}
//		for k, v := range c2 {
//			if k == "command" {
//				_cmd.Cmd ,ok = v.(string)
//				if !ok {
//					err = errors.New("解析配置 handler." + handlerName + ".cmd 失败")
//					golog.Error(err.Error())
//					return
//				}
//			}
//			if k == "args"{
//				var args = []string{}
//				argList , ok := v.([]interface {})
//				if !ok {
//					err = errors.New("解析配置 handler." + handlerName + ".cmd 失败")
//					golog.Error(err.Error())
//					return
//				}
//				for _, a := range argList {
//					arg , ok := a.(string)
//					if !ok {
//						err = errors.New("解析配置 handler." + handlerName + ".cmd 失败")
//						golog.Error(err.Error())
//						return
//					}
//					args = append(args, arg)
//				}
//				_cmd.Args = args
//			}
//		}
//		cmds = append(cmds,_cmd)
//	}
//
//
//	return
//
//}
