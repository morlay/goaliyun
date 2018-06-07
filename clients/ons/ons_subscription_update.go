package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsSubscriptionUpdateRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	ReadEnable   string `position:"Query" name:"ReadEnable"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsSubscriptionUpdateRequest) Invoke(client goaliyun.Client) (*OnsSubscriptionUpdateResponse, error) {
	resp := &OnsSubscriptionUpdateResponse{}
	err := client.Request("ons", "OnsSubscriptionUpdate", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsSubscriptionUpdateResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
}
