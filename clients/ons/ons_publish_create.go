package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsPublishCreateRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	AppName      string `position:"Query" name:"AppName"`
	Topic        string `position:"Query" name:"Topic"`
	ProducerId   string `position:"Query" name:"ProducerId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsPublishCreateRequest) Invoke(client goaliyun.Client) (*OnsPublishCreateResponse, error) {
	resp := &OnsPublishCreateResponse{}
	err := client.Request("ons", "OnsPublishCreate", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsPublishCreateResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
}
