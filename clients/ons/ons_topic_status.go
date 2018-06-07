package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsTopicStatusRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsTopicStatusRequest) Invoke(client goaliyun.Client) (*OnsTopicStatusResponse, error) {
	resp := &OnsTopicStatusResponse{}
	err := client.Request("ons", "OnsTopicStatus", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsTopicStatusResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsTopicStatusData
}

type OnsTopicStatusData struct {
	TotalCount    goaliyun.Integer
	LastTimeStamp goaliyun.Integer
}
