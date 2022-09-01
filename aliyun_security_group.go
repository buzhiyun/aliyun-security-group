package main

import (
	"flag"
	"github.com/buzhiyun/aliyun-security-group/access"
	"github.com/buzhiyun/aliyun-security-group/db"
	"github.com/kataras/golog"
)

func main() {
	golog.SetTimeFormat("2006-01-02 15:04:05.000")


	migrate := flag.Bool("migrate",false,"初始化数据表")
	check := flag.Bool("check",false,"检查现有记录DNS解析地址是否有效")
	debug := flag.Bool("debug",false,"开启debug")
	update := flag.Bool("update", false , "检查失效IP，并更新")
	flag.Parse()


	if *debug {
		golog.Info("设置日志等级： debug")
		golog.SetLevel("debug")
	}

	db.InitDb()


	if *migrate {
		_ = db.Migrate()
		return
	}

	if *check {
		access.CheckRecord()
		return
	}


	if *update {
		access.HandleInvaildRecord()
		return
	}

}

