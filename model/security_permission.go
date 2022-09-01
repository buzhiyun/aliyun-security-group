package model

import "time"

type SecurityPermission struct {
	SourceGroupId           string    `json:"SourceGroupId"`
	Policy                  string    `json:"Policy"`
	Description             string    `json:"Description"`
	SourcePortRange         string    `json:"SourcePortRange"`
	Priority                int       `json:"Priority"`
	CreateTime              time.Time `json:"CreateTime"`
	DestPrefixListName      string    `json:"DestPrefixListName"`
	Ipv6SourceCidrIp        string    `json:"Ipv6SourceCidrIp"`
	NicType                 string    `json:"NicType"`
	DestGroupId             string    `json:"DestGroupId"`
	Direction               string    `json:"Direction"`
	SourceGroupName         string    `json:"SourceGroupName"`
	PortRange               string    `json:"PortRange"`
	DestGroupOwnerAccount   string    `json:"DestGroupOwnerAccount"`
	DestPrefixListId        string    `json:"DestPrefixListId"`
	SourceCidrIp            string    `json:"SourceCidrIp"`
	SourcePrefixListName    string    `json:"SourcePrefixListName"`
	IpProtocol              string    `json:"IpProtocol"`
	DestCidrIp              string    `json:"DestCidrIp"`
	DestGroupName           string    `json:"DestGroupName"`
	SourceGroupOwnerAccount string    `json:"SourceGroupOwnerAccount"`
	Ipv6DestCidrIp          string    `json:"Ipv6DestCidrIp"`
	SourcePrefixListId      string    `json:"SourcePrefixListId"`
}
