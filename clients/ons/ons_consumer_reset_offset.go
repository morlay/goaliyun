package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsConsumerResetOffsetRequest struct {
	PreventCache   int64  `position:"Query" name:"PreventCache"`
	OnsRegionId    string `position:"Query" name:"OnsRegionId"`
	OnsPlatform    string `position:"Query" name:"OnsPlatform"`
	ConsumerId     string `position:"Query" name:"ConsumerId"`
	Topic          string `position:"Query" name:"Topic"`
	ResetTimestamp int64  `position:"Query" name:"ResetTimestamp"`
	Type           int64  `position:"Query" name:"Type"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *OnsConsumerResetOffsetRequest) Invoke(client goaliyun.Client) (*OnsConsumerResetOffsetResponse, error) {
	resp := &OnsConsumerResetOffsetResponse{}
	err := client.Request("ons", "OnsConsumerResetOffset", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsConsumerResetOffsetResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
}
