package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsMqttQueryMsgByPubInfoRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	ClientId     string `position:"Query" name:"ClientId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	BeginTime    int64  `position:"Query" name:"BeginTime"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsMqttQueryMsgByPubInfoRequest) Invoke(client goaliyun.Client) (*OnsMqttQueryMsgByPubInfoResponse, error) {
	resp := &OnsMqttQueryMsgByPubInfoResponse{}
	err := client.Request("ons", "OnsMqttQueryMsgByPubInfo", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsMqttQueryMsgByPubInfoResponse struct {
	RequestId      goaliyun.String
	HelpUrl        goaliyun.String
	MqttMessageDos OnsMqttQueryMsgByPubInfoMqttMessageDoList
}

type OnsMqttQueryMsgByPubInfoMqttMessageDo struct {
	TraceId goaliyun.String
	PubTime goaliyun.Integer
}

type OnsMqttQueryMsgByPubInfoMqttMessageDoList []OnsMqttQueryMsgByPubInfoMqttMessageDo

func (list *OnsMqttQueryMsgByPubInfoMqttMessageDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMqttQueryMsgByPubInfoMqttMessageDo)
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
