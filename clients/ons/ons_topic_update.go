package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsTopicUpdateRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Perm         int64  `position:"Query" name:"Perm"`
	Topic        string `position:"Query" name:"Topic"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsTopicUpdateRequest) Invoke(client goaliyun.Client) (*OnsTopicUpdateResponse, error) {
	resp := &OnsTopicUpdateResponse{}
	err := client.Request("ons", "OnsTopicUpdate", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsTopicUpdateResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
}
