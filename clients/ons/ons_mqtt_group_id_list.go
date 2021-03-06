package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsMqttGroupIdListRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsMqttGroupIdListRequest) Invoke(client goaliyun.Client) (*OnsMqttGroupIdListResponse, error) {
	resp := &OnsMqttGroupIdListResponse{}
	err := client.Request("ons", "OnsMqttGroupIdList", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsMqttGroupIdListResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsMqttGroupIdListMqttGroupIdDoList
}

type OnsMqttGroupIdListMqttGroupIdDo struct {
	Id          goaliyun.Integer
	ChannelId   goaliyun.Integer
	OnsRegionId goaliyun.String
	RegionName  goaliyun.String
	Owner       goaliyun.String
	GroupId     goaliyun.String
	Topic       goaliyun.String
	Status      goaliyun.Integer
	CreateTime  goaliyun.Integer
	UpdateTime  goaliyun.Integer
}

type OnsMqttGroupIdListMqttGroupIdDoList []OnsMqttGroupIdListMqttGroupIdDo

func (list *OnsMqttGroupIdListMqttGroupIdDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMqttGroupIdListMqttGroupIdDo)
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
