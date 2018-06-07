package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsMqttManualUpdateRuleRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	OwnerId      string `position:"Query" name:"OwnerId"`
	AdminKey     string `position:"Query" name:"AdminKey"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsMqttManualUpdateRuleRequest) Invoke(client goaliyun.Client) (*OnsMqttManualUpdateRuleResponse, error) {
	resp := &OnsMqttManualUpdateRuleResponse{}
	err := client.Request("ons", "OnsMqttManualUpdateRule", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsMqttManualUpdateRuleResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
}
