package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsMqttQueryClientByGroupIdRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	GroupId      string `position:"Query" name:"GroupId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsMqttQueryClientByGroupIdRequest) Invoke(client goaliyun.Client) (*OnsMqttQueryClientByGroupIdResponse, error) {
	resp := &OnsMqttQueryClientByGroupIdResponse{}
	err := client.Request("ons", "OnsMqttQueryClientByGroupId", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsMqttQueryClientByGroupIdResponse struct {
	RequestId       goaliyun.String
	HelpUrl         goaliyun.String
	MqttClientSetDo OnsMqttQueryClientByGroupIdMqttClientSetDo
}

type OnsMqttQueryClientByGroupIdMqttClientSetDo struct {
	OnlineCount  goaliyun.Integer
	PersistCount goaliyun.Integer
}
