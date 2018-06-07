package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsSubscriptionSearchRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	Search       string `position:"Query" name:"Search"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsSubscriptionSearchRequest) Invoke(client goaliyun.Client) (*OnsSubscriptionSearchResponse, error) {
	resp := &OnsSubscriptionSearchResponse{}
	err := client.Request("ons", "OnsSubscriptionSearch", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsSubscriptionSearchResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsSubscriptionSearchSubscribeInfoDoList
}

type OnsSubscriptionSearchSubscribeInfoDo struct {
	Id          goaliyun.Integer
	ChannelId   goaliyun.Integer
	ChannelName goaliyun.String
	OnsRegionId goaliyun.String
	RegionName  goaliyun.String
	Owner       goaliyun.String
	ConsumerId  goaliyun.String
	Topic       goaliyun.String
	Status      goaliyun.Integer
	StatusName  goaliyun.String
	CreateTime  goaliyun.Integer
	UpdateTime  goaliyun.Integer
}

type OnsSubscriptionSearchSubscribeInfoDoList []OnsSubscriptionSearchSubscribeInfoDo

func (list *OnsSubscriptionSearchSubscribeInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsSubscriptionSearchSubscribeInfoDo)
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
