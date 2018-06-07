package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsMqttGroupIdCreateRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	GroupId      string `position:"Query" name:"GroupId"`
	Topic        string `position:"Query" name:"Topic"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsMqttGroupIdCreateRequest) Invoke(client goaliyun.Client) (*OnsMqttGroupIdCreateResponse, error) {
	resp := &OnsMqttGroupIdCreateResponse{}
	err := client.Request("ons", "OnsMqttGroupIdCreate", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsMqttGroupIdCreateResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
}
