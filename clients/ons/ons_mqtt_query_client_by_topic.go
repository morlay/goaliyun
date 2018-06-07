package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsMqttQueryClientByTopicRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ParentTopic  string `position:"Query" name:"ParentTopic"`
	SubTopic     string `position:"Query" name:"SubTopic"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsMqttQueryClientByTopicRequest) Invoke(client goaliyun.Client) (*OnsMqttQueryClientByTopicResponse, error) {
	resp := &OnsMqttQueryClientByTopicResponse{}
	err := client.Request("ons", "OnsMqttQueryClientByTopic", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsMqttQueryClientByTopicResponse struct {
	RequestId       goaliyun.String
	HelpUrl         goaliyun.String
	MqttClientSetDo OnsMqttQueryClientByTopicMqttClientSetDo
}

type OnsMqttQueryClientByTopicMqttClientSetDo struct {
	OnlineCount  goaliyun.Integer
	PersistCount goaliyun.Integer
}
