package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsTopicSearchRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	Search       string `position:"Query" name:"Search"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsTopicSearchRequest) Invoke(client goaliyun.Client) (*OnsTopicSearchResponse, error) {
	resp := &OnsTopicSearchResponse{}
	err := client.Request("ons", "OnsTopicSearch", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsTopicSearchResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsTopicSearchPublishInfoDoList
}

type OnsTopicSearchPublishInfoDo struct {
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

type OnsTopicSearchPublishInfoDoList []OnsTopicSearchPublishInfoDo

func (list *OnsTopicSearchPublishInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTopicSearchPublishInfoDo)
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
