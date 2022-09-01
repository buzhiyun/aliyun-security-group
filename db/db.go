package db

import (
	"github.com/buzhiyun/aliyun-security-group/model"
	"github.com/buzhiyun/go-utils/cfg"
	"github.com/kataras/golog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

/*
初始化数据库连接
*/
func InitDb()  {
	dsn ,ok := cfg.Config().GetString("mysql.dsn")
	if !ok {
		golog.Fatalf("读取数据库配置文件失败 mysql.dsn")
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // DSN data source name
		//DSN: "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize: 256, // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		//DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		golog.Fatalf("初始化数据库失败，%s",err.Error())
	}

	DB = db

}

/*
初始化各表
*/
func Migrate() (err error) {
	err = DB.AutoMigrate(&model.Business{},model.SecurityGroup{})
	err = DB.AutoMigrate(&model.FirewallRecord{})
	if err != nil {
		golog.Errorf("Migrate 异常, %s",err.Error())
	}
	golog.Info("初始表成功")
	return
}