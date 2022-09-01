package model

import "time"

type SecurityGroup struct {
	//Id					int				`gorm:"primaryKey"`
	SecurityGroupId		string			`gorm:"primaryKey;not null;comment:安全组ID"`			// 开墙记录表的外键
	SecurityGroupName	string			`gorm:"comment:安全组名称"`
	Active				bool			`gorm:"comment:是否有效"`
	FirewallRecord	FirewallRecord	`gorm:"foreignKey:SecurityGroupId;references:SecurityGroupId"`
}

type Business struct {
	BizCode		int				`gorm:"primaryKey;not null;comment:业务代码"`
	BizName		string			`gorm:"unique:comment:业务名称"`				// 开墙记录表的外键
	FirewallRecord	FirewallRecord	`gorm:"foreignKey:BizName;references:BizName"`
}

type FirewallRecord struct {
	Id			uint					`gorm:"primaryKey"`
	Domain		string					`gorm:"uniqueIndex:ip_info;size:512;comment:域名"`
	IP			string					`gorm:"uniqueIndex:ip_info;size:32;comment:IP地址"`
	Port		int64					`gorm:"uniqueIndex:ip_info;comment:端口号"`   //   -1 时候不生效
	IsDelete	bool					`gorm:"uniqueIndex:ip_info;comment:是否已经删除"`
	BizName		string					`gorm:"not null;comment:业务名称"`
	SecurityGroupId		string			`gorm:"uniqueIndex:ip_info;size:64;not null;comment:安全组ID"`
	Handlers	string					`gorm:"size:2048;comment:同步记录到主机的模块"`
	ExpiryDate			time.Time		`gorm:"comment:有效期"`
	Comment		string					`gorm:"comment:记录描述"`
}