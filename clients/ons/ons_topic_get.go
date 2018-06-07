package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsTopicGetRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsTopicGetRequest) Invoke(client goaliyun.Client) (*OnsTopicGetResponse, error) {
	resp := &OnsTopicGetResponse{}
	err := client.Request("ons", "OnsTopicGet", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsTopicGetResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsTopicGetPublishInfoDoList
}

type OnsTopicGetPublishInfoDo struct {
	Id           goaliyun.Integer
	ChannelId    goaliyun.Integer
	ChannelName  goaliyun.String
	OnsRegionId  goaliyun.String
	RegionName   goaliyun.String
	Topic        goaliyun.String
	Owner        goaliyun.String
	Relation     goaliyun.Integer
	RelationName goaliyun.String
	Status       goaliyun.Integer
	StatusName   goaliyun.String
	Appkey       goaliyun.String
	CreateTime   goaliyun.Integer
	UpdateTime   goaliyun.Integer
	Remark       goaliyun.String
}

type OnsTopicGetPublishInfoDoList []OnsTopicGetPublishInfoDo

func (list *OnsTopicGetPublishInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTopicGetPublishInfoDo)
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
