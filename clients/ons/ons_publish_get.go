package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsPublishGetRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	ProducerId   string `position:"Query" name:"ProducerId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsPublishGetRequest) Invoke(client goaliyun.Client) (*OnsPublishGetResponse, error) {
	resp := &OnsPublishGetResponse{}
	err := client.Request("ons", "OnsPublishGet", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsPublishGetResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsPublishGetPublishInfoDoList
}

type OnsPublishGetPublishInfoDo struct {
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

type OnsPublishGetPublishInfoDoList []OnsPublishGetPublishInfoDo

func (list *OnsPublishGetPublishInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsPublishGetPublishInfoDo)
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
