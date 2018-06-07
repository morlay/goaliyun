package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsTopicListRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsTopicListRequest) Invoke(client goaliyun.Client) (*OnsTopicListResponse, error) {
	resp := &OnsTopicListResponse{}
	err := client.Request("ons", "OnsTopicList", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsTopicListResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsTopicListPublishInfoDoList
}

type OnsTopicListPublishInfoDo struct {
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

type OnsTopicListPublishInfoDoList []OnsTopicListPublishInfoDo

func (list *OnsTopicListPublishInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTopicListPublishInfoDo)
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
