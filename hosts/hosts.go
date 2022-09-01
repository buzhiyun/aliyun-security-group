package hosts

import (
	"bufio"
	"errors"
	"github.com/buzhiyun/go-utils/cfg"
	"github.com/buzhiyun/go-utils/file"
	"github.com/kataras/golog"
	"os"
	"path/filepath"
	"strings"
)

/*
更新域名
*/
func UpdateHosts(item,domain,ip string,isWindowsFile... bool) (changed bool, err error) {
	changed = true
	dir , ok :=cfg.Config().GetString("handler.template_dir")
	if !ok {
		err = errors.New("获取配置 handler.template_dir 错误")
		golog.Error(err)
		return
	}
	hostsFile := filepath.Join(dir,item,"hosts")
	if !file.FileExist(hostsFile){
		err = errors.New(item + " hosts 模板文件不存在 " + hostsFile )
		golog.Error(err)
		return
	}

	var winFile = false
	if len(isWindowsFile)>0 && isWindowsFile[0] {
		winFile = true
	}
	changed ,err = UpdateFile(hostsFile,domain,ip,winFile)
	if err == nil && changed{
		os.Rename(hostsFile + ".new",hostsFile)
	}

	return

}


/*
更新hosts 文件里面域名对应的IP
*/
func UpdateFile(fpath , domain,ip string,winFile bool) (changed bool,err error) {
	changed = true
	f ,err := os.Open(fpath)
	if err != nil{
		golog.Errorf("读取hosts 文件失败 %s ,%s", fpath,err.Error())
		return
	}

	fnew ,err := os.OpenFile(fpath + ".new", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil{
		golog.Errorf("打开新文件失败, %s",err.Error())
		return
	}
	defer f.Close()
	defer  fnew.Close()

	buf := bufio.NewScanner(f)
	golog.Infof("开始更新 %s , %s ,%s ",fpath,ip ,domain)

	var replaced = false
	endLine := "\n"
	if winFile { endLine = "\r\n" }
	for {
		if !buf.Scan() {
			break
		}
		line := buf.Text()
		line = strings.TrimSpace(line)
		//golog.Debugf(line)
		// 替换host
		if strings.Index(line,domain) > 0 {
			if strings.Index(line,ip + "\t") >=0 ||  strings.Index(line ,ip + " ") >=0 {
				golog.Debugf("hosts 中存在正确记录，无需变更")
				return false,nil
			}

			_ ,err = fnew.WriteString(ip + "\t"+ domain + endLine)
			replaced = true
			continue
		}
		_ ,err = fnew.WriteString(line + endLine)
	}
	// 在末尾添加新的host 记录
	if !replaced {
		_ ,err = fnew.WriteString(ip + "\t"+ domain + endLine)
	}

	return
}