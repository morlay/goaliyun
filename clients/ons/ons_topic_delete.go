package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsTopicDeleteRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	Cluster      string `position:"Query" name:"Cluster"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsTopicDeleteRequest) Invoke(client goaliyun.Client) (*OnsTopicDeleteResponse, error) {
	resp := &OnsTopicDeleteResponse{}
	err := client.Request("ons", "OnsTopicDelete", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsTopicDeleteResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
}
