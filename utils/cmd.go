package utils

import (
	"github.com/kataras/golog"
	"os/exec"
)

func Exec(command string,args... string) (result []byte) {
	golog.Debugf("exec command: %s %s",command ,args)
	cmd := exec.Command(command ,args...)
	result = syncCmdOut(cmd)
	return
}


func syncCmdOut(cmd *exec.Cmd) (cmdOutput []byte) {

	stdout, err := cmd.CombinedOutput()

	if err != nil {
		golog.Errorf("执行出错 %s %v：%s", cmd.Path , cmd.Args ,  err.Error())
		return make([]byte, 0)
	}
	//Linux 读取需要加上这段  从GBK读取转换UTF-8字符集
	//b, err := simplifiedchinese.GBK.NewDecoder().Bytes(stdout)
	//if err != nil {
	//	log.Println("转换出错", err.Error())
	//}
	//return b
	return stdout
}
