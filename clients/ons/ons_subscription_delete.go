package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsSubscriptionDeleteRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	Topic        string `position:"Query" name:"Topic"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsSubscriptionDeleteRequest) Invoke(client goaliyun.Client) (*OnsSubscriptionDeleteResponse, error) {
	resp := &OnsSubscriptionDeleteResponse{}
	err := client.Request("ons", "OnsSubscriptionDelete", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsSubscriptionDeleteResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
}
