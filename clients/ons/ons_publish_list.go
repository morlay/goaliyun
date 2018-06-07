package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsPublishListRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsPublishListRequest) Invoke(client goaliyun.Client) (*OnsPublishListResponse, error) {
	resp := &OnsPublishListResponse{}
	err := client.Request("ons", "OnsPublishList", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsPublishListResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsPublishListPublishInfoDoList
}

type OnsPublishListPublishInfoDo struct {
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

type OnsPublishListPublishInfoDoList []OnsPublishListPublishInfoDo

func (list *OnsPublishListPublishInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsPublishListPublishInfoDo)
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
