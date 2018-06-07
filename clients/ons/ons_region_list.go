package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsRegionListRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsRegionListRequest) Invoke(client goaliyun.Client) (*OnsRegionListResponse, error) {
	resp := &OnsRegionListResponse{}
	err := client.Request("ons", "OnsRegionList", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsRegionListResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsRegionListRegionDoList
}

type OnsRegionListRegionDo struct {
	Id          goaliyun.Integer
	OnsRegionId goaliyun.String
	RegionName  goaliyun.String
	ChannelId   goaliyun.Integer
	ChannelName goaliyun.String
	CreateTime  goaliyun.Integer
	UpdateTime  goaliyun.Integer
}

type OnsRegionListRegionDoList []OnsRegionListRegionDo

func (list *OnsRegionListRegionDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsRegionListRegionDo)
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
