package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsSubscriptionCreateRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	Topic        string `position:"Query" name:"Topic"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsSubscriptionCreateRequest) Invoke(client goaliyun.Client) (*OnsSubscriptionCreateResponse, error) {
	resp := &OnsSubscriptionCreateResponse{}
	err := client.Request("ons", "OnsSubscriptionCreate", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsSubscriptionCreateResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
}
