package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsPublishSearchRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	Search       string `position:"Query" name:"Search"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsPublishSearchRequest) Invoke(client goaliyun.Client) (*OnsPublishSearchResponse, error) {
	resp := &OnsPublishSearchResponse{}
	err := client.Request("ons", "OnsPublishSearch", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsPublishSearchResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsPublishSearchPublishInfoDoList
}

type OnsPublishSearchPublishInfoDo struct {
	Id          goaliyun.Integer
	ChannelId   goaliyun.Integer
	ChannelName goaliyun.String
	OnsRegionId goaliyun.String
	RegionName  goaliyun.String
	Owner       goaliyun.String
	ProducerId  goaliyun.String
	Topic       goaliyun.String
	Status      goaliyun.Integer
	StatusName  goaliyun.String
	CreateTime  goaliyun.Integer
	UpdateTime  goaliyun.Integer
}

type OnsPublishSearchPublishInfoDoList []OnsPublishSearchPublishInfoDo

func (list *OnsPublishSearchPublishInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsPublishSearchPublishInfoDo)
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
