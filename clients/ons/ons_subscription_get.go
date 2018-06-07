package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsSubscriptionGetRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	Topic        string `position:"Query" name:"Topic"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsSubscriptionGetRequest) Invoke(client goaliyun.Client) (*OnsSubscriptionGetResponse, error) {
	resp := &OnsSubscriptionGetResponse{}
	err := client.Request("ons", "OnsSubscriptionGet", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsSubscriptionGetResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsSubscriptionGetSubscribeInfoDoList
}

type OnsSubscriptionGetSubscribeInfoDo struct {
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

type OnsSubscriptionGetSubscribeInfoDoList []OnsSubscriptionGetSubscribeInfoDo

func (list *OnsSubscriptionGetSubscribeInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsSubscriptionGetSubscribeInfoDo)
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
