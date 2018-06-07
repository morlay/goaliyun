package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsPublishDeleteRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	ProducerId   string `position:"Query" name:"ProducerId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsPublishDeleteRequest) Invoke(client goaliyun.Client) (*OnsPublishDeleteResponse, error) {
	resp := &OnsPublishDeleteResponse{}
	err := client.Request("ons", "OnsPublishDelete", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsPublishDeleteResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
}
