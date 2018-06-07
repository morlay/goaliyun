package afs

import (
	"github.com/morlay/goaliyun"
)

type ConfigurationStyleRequest struct {
	ResourceOwnerId     int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp            string `position:"Query" name:"SourceIp"`
	ConfigurationMethod string `position:"Query" name:"ConfigurationMethod"`
	ApplyType           string `position:"Query" name:"ApplyType"`
	Scene               string `position:"Query" name:"Scene"`
	RegionId            string `position:"Query" name:"RegionId"`
}

func (req *ConfigurationStyleRequest) Invoke(client goaliyun.Client) (*ConfigurationStyleResponse, error) {
	resp := &ConfigurationStyleResponse{}
	err := client.Request("afs", "ConfigurationStyle", "2018-01-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ConfigurationStyleResponse struct {
	RequestId goaliyun.String
	BizCode   goaliyun.String
	CodeData  ConfigurationStyleCodeData
}

type ConfigurationStyleCodeData struct {
	Html   goaliyun.String
	Net    goaliyun.String
	Php    goaliyun.String
	Python goaliyun.String
	Java   goaliyun.String
	NodeJs goaliyun.String
}
