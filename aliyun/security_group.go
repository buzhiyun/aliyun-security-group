package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/kataras/golog"
	"time"
)


/*
添加安全组策略
*/
func AddPermission(ip,port,domain,bizName,describe,securityGroupId string) (err error)  {
	request := ecs.CreateAuthorizeSecurityGroupEgressRequest()
	request.Scheme = "https"
	request.SetConnectTimeout(10 * time.Second)

	request.SecurityGroupId = securityGroupId

	request.Permissions = &[]ecs.AuthorizeSecurityGroupEgressPermissions{
		{
			Policy:                "accept",
			Priority:              "2",
			IpProtocol:            "TCP",
			DestCidrIp:            ip,
			//Ipv6DestCidrIp:        "",
			//DestGroupId:           "",
			//DestPrefixListId:      "",
			PortRange:             port + "/" + port,
			//SourceCidrIp:          "",
			//Ipv6SourceCidrIp:      "",
			//SourcePortRange:       "",
			//DestGroupOwnerAccount: "",
			//DestGroupOwnerId:      "",
			NicType:               "intranet",
			Description:           bizName + " - " + domain + " " + describe,
		},
	}


	response, err := aliyunClient.AuthorizeSecurityGroupEgress(request)
	if err != nil {
		golog.Errorf("添加安全组 %s 策略错误 ,%s",securityGroupId , err.Error())
		return
	}
	golog.Infof("添加安全组 %s 策略返回 %s", securityGroupId, response.GetHttpContentString())
	return
}


/*
删除安全组策略
*/
func DeletePermission(ip,port,domain,bizName,describe,securityGroupId string) (err error) {

	request := ecs.CreateRevokeSecurityGroupEgressRequest()
	request.SetConnectTimeout(10 * time.Second)

	request.Scheme = "https"

	request.SecurityGroupId = securityGroupId
	request.Permissions = &[]ecs.RevokeSecurityGroupEgressPermissions{
		{
			Policy:                "accept",
			Priority:              "2",
			IpProtocol:            "TCP",
			DestCidrIp:            ip,
			//Ipv6DestCidrIp:        "",
			//DestGroupId:           "",
			//DestPrefixListId:      "",
			PortRange:             port + "/" + port,
			//SourceCidrIp:          "",
			//Ipv6SourceCidrIp:      "",
			//SourcePortRange:       "",
			//DestGroupOwnerAccount: "",
			//DestGroupOwnerId:      "",
			NicType:               "intranet",
			Description:           bizName + " - " + domain + " " + describe,
		},
		{
			Policy:                "accept",
			Priority:              "1",
			IpProtocol:            "TCP",
			DestCidrIp:            ip,
			//Ipv6DestCidrIp:        "",
			//DestGroupId:           "",
			//DestPrefixListId:      "",
			PortRange:             port + "/" + port,
			//SourceCidrIp:          "",
			//Ipv6SourceCidrIp:      "",
			//SourcePortRange:       "",
			//DestGroupOwnerAccount: "",
			//DestGroupOwnerId:      "",
			NicType:               "intranet",
			Description:           bizName + " - " + domain + " " + describe,
		},
	}

	response, err := aliyunClient.RevokeSecurityGroupEgress(request)
	if err != nil {
		golog.Errorf("删除安全组 %s 策略错误 ,%s" ,securityGroupId ,err.Error())
		return
	}
	golog.Infof("删除安全组 %s 策略返回 %s", securityGroupId ,response.GetHttpContentString())
	return
}

/*
获取阿里云安全组的所有策略
*/
func GetPermissions(securityGroupId string) (permissions *[]ecs.Permission , err error) {
	request := ecs.CreateDescribeSecurityGroupAttributeRequest()

	request.Scheme = "https"
	request.SetConnectTimeout(10 * time.Second)
	request.SecurityGroupId = securityGroupId

	response, err := aliyunClient.DescribeSecurityGroupAttribute(request)
	if err != nil {
		golog.Errorf("获取安全组 %s 策略错误 ,%s", securityGroupId,err.Error())
		return
	}

	permissions = &response.Permissions.Permission
	return
}