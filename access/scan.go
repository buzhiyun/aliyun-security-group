package access

import (
	"github.com/buzhiyun/aliyun-security-group/aliyun"
	"github.com/buzhiyun/aliyun-security-group/db"
	"github.com/buzhiyun/aliyun-security-group/handler"
	"github.com/buzhiyun/aliyun-security-group/model"
	"github.com/kataras/golog"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
扫描阿里云的安全组列表，写入到数据库 record 记录表中
*/
func ScanAliyunPermission()  {
	var sg []model.SecurityGroup
	db.DB.Model(&model.SecurityGroup{}).Where(&model.SecurityGroup{Active: true}).Find(&sg)

	r, _ := regexp.Compile(`(.+)\ -\ (.+)\ (.+)`)
	for _, group := range sg {
		golog.Debugf("%s",group.SecurityGroupName)
		permissions , err := aliyun.GetPermissions(group.SecurityGroupId)
		if err != nil {
			continue
		}
		for _, permission := range *permissions {

			desc := r.FindStringSubmatch(permission.Description)
			if len(desc) == 4 {  // 查询到符合规则的说明
				golog.Debugf("发现阿里云安全组规则：[ %s ]\t%s\t[ %s ]\t%s",group.SecurityGroupName, desc[1],desc[2],desc[3])

				ports :=strings.Split(permission.PortRange,"/")
				if len(ports) != 2 {
					golog.Errorf("异常的端口 %s %s %s %s",group.SecurityGroupName, permission.DestCidrIp, permission.PortRange , permission.Description)
					continue
				}
				port ,err := strconv.ParseInt(ports[0],10,16)
				if err != nil {
					golog.Errorf("异常的端口解析 %s %s %s %s",group.SecurityGroupName, permission.DestCidrIp, permission.PortRange , permission.Description)
					continue
				}

				var biz model.Business
				db.DB.Model(&model.Business{}).Where("biz_name = ?",desc[1]).First(&biz)

				// 如果业务线表没有记录，就添加记录
				if biz.BizName == ""|| biz.BizCode == 0 {
					result := db.DB.Create(&model.Business{
						BizName:  desc[1],
					})
					golog.Debugf("写入业务线数据，RowsAffected %v",result.RowsAffected)
					biz.BizName = desc[1]

				}

				//db.DB.Where("biz_name = ?", biz.BizName).First(&biz)

				result := db.DB.Create(&model.FirewallRecord{
				//result := db.DB.Debug().Create(&model.FirewallRecord{
						Domain:          strings.TrimSpace(desc[2]),
						IP:              permission.DestCidrIp,
						Port:            port,
						BizName: 		 biz.BizName,
						SecurityGroupId: group.SecurityGroupId,
						ExpiryDate:      time.Now().AddDate(100,0,0),
						Comment:         desc[3],
					})

				golog.Debugf("写入防火墙记录数据，RowsAffected %v",result.RowsAffected)

			}

		}
	}

}

type invaildRecord struct {
	record		model.FirewallRecord
	vaildIp		[]string
}

type resolvMap map[string][]string

var resolvRecord resolvMap

/*
检查现有记录是否 dns 解析异常
*/
func CheckRecord() (invalidRec []invaildRecord) {
	resolvRecord = make(resolvMap)

	var record []model.FirewallRecord
	result := db.DB.Where("expiry_date >= ? and is_delete = ?" , time.Now() ,false ).Find(&record)
	if result.Error!= nil{
		golog.Errorf("查询数据库 firewall_record 错误,%s",result.Error.Error())
		return
	}
	for _, firewallRecord := range record {
		var ok = false
		ipList , err := getIp(strings.Split(firewallRecord.Domain,",")[0])
		if err != nil {
			continue
		}
		for _, ip := range ipList {
			if ip == firewallRecord.IP {
				ok = true
				break
			}
		}
		if !ok {
			golog.Warnf("%s %s %s ip %s 异常，解析到的正确ip ：%v ",firewallRecord.SecurityGroupId, firewallRecord.BizName , firewallRecord.Domain ,firewallRecord.IP , ipList )

			invalidRec = append(invalidRec, invaildRecord{
				record:  firewallRecord,
				vaildIp: ipList,
			})
		}
	}
	return
}


/*
对异常记录做处理
*/
func HandleInvaildRecord()  {
	invaildRec := CheckRecord()
	for _, record := range invaildRec {
		newIp := ""
		// 从合法IP 中挑一个ipv4可用IP
		for _, ip := range record.vaildIp {
			if strings.Index(ip,":") < 0 {
				newIp = ip
				break
			}
		}
		if len(newIp) == 0 {
			golog.Errorf("%s %s %s %s 合法IP 中没有 ipv4 地址 %v" ,record.record.SecurityGroupId,record.record.BizName, record.record.Domain ,record.record.Comment ,record.vaildIp)
			continue
		}

		// 开始真正处理异常记录

		// 改主机host 或者改主机解析

		if len(strings.TrimSpace(record.record.Handlers)) == 0 {
			golog.Warnf("[%s] %s  %s:%v  %s 没有设置handler",record.record.SecurityGroupId, record.record.BizName,record.record.IP,record.record.Port,record.record.Domain)
			continue
		}


		handlerNames := strings.Split(strings.TrimSpace(record.record.Handlers),",")

		success := false  // 必须所有的handler都正确处理完，才能 去删除或者清理阿里云记录
		for _, name := range handlerNames {
			if len(name) == 0 {
				golog.Errorf("[%s] %s  %s:%v  %s 存在空的handler %#v",record.record.SecurityGroupId, record.record.BizName,record.record.IP,record.record.Port,record.record.Domain,handlerNames)
				break
			}
			hdl , ok := handler.HandlerMap[name]
			if !ok {
				golog.Errorf("[%s] %s  %s:%v  %s  未找到对应的 handle %s", record.record.SecurityGroupId,record.record.BizName,record.record.IP,record.record.Port,record.record.Domain,name)
				break
			}
			if !hdl.Available(){
				golog.Errorf("%s 目前不可用不能处理",hdl.String())
				success = false
				break
			}

			// 更新解析
			if err := hdl.UpdateHost(record.record.Domain, newIp);err != nil {
				golog.Errorf("%s 更新hosts 失败, %s",hdl.String(),err.Error())
				success = false
				break
			}

			success = true
		}

		if success{
			golog.Infof("替换阿里云安全组策略")

			aliyun.DeletePermission(record.record.IP,strconv.FormatInt(record.record.Port,10),record.record.Domain,record.record.BizName,record.record.Comment,record.record.SecurityGroupId)
			aliyun.AddPermission(newIp,strconv.FormatInt(record.record.Port,10),record.record.Domain,record.record.BizName,record.record.Comment,record.record.SecurityGroupId)

			result := db.DB.Model(&model.FirewallRecord{}).Where("id = ?" , record.record.Id).Update("ip",newIp)
			if result.Error != nil {
				if strings.Index(result.Error.Error() ,"Error 1062") >= 0 {
					golog.Warnf("发现已有相同配置，该配置设置为删除")
					rst := db.DB.Model(&model.FirewallRecord{}).Where("id = ?" , record.record.Id).Update("is_delete",true)
					if rst.Error != nil {
						if strings.Index(rst.Error.Error() ,"Error 1062") >= 0 {
							db.DB.Model(&model.FirewallRecord{}).Where("id = ?" , record.record.Id).Delete(&model.FirewallRecord{})
						}
					}
				} else {
					golog.Errorf("更新firewall_record记录失败")
				}
			}
			//golog.Info(record.record)
		}

	}
}