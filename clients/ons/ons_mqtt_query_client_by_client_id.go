package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsMqttQueryClientByClientIdRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	ClientId     string `position:"Query" name:"ClientId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsMqttQueryClientByClientIdRequest) Invoke(client goaliyun.Client) (*OnsMqttQueryClientByClientIdResponse, error) {
	resp := &OnsMqttQueryClientByClientIdResponse{}
	err := client.Request("ons", "OnsMqttQueryClientByClientId", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsMqttQueryClientByClientIdResponse struct {
	RequestId        goaliyun.String
	HelpUrl          goaliyun.String
	MqttClientInfoDo OnsMqttQueryClientByClientIdMqttClientInfoDo
}

type OnsMqttQueryClientByClientIdMqttClientInfoDo struct {
	Online          bool
	ClientId        goaliyun.String
	SocketChannel   goaliyun.String
	LastTouch       goaliyun.Integer
	SubScriptonData OnsMqttQueryClientByClientIdSubscriptionDoList
}

type OnsMqttQueryClientByClientIdSubscriptionDo struct {
	ParentTopic goaliyun.String
	SubTopic    goaliyun.String
	Qos         goaliyun.Integer
}

type OnsMqttQueryClientByClientIdSubscriptionDoList []OnsMqttQueryClientByClientIdSubscriptionDo

func (list *OnsMqttQueryClientByClientIdSubscriptionDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMqttQueryClientByClientIdSubscriptionDo)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
