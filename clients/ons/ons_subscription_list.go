package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsSubscriptionListRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsSubscriptionListRequest) Invoke(client goaliyun.Client) (*OnsSubscriptionListResponse, error) {
	resp := &OnsSubscriptionListResponse{}
	err := client.Request("ons", "OnsSubscriptionList", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsSubscriptionListResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsSubscriptionListSubscribeInfoDoList
}

type OnsSubscriptionListSubscribeInfoDo struct {
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

type OnsSubscriptionListSubscribeInfoDoList []OnsSubscriptionListSubscribeInfoDo

func (list *OnsSubscriptionListSubscribeInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsSubscriptionListSubscribeInfoDo)
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
