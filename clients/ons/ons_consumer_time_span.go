package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsConsumerTimeSpanRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	Topic        string `position:"Query" name:"Topic"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsConsumerTimeSpanRequest) Invoke(client goaliyun.Client) (*OnsConsumerTimeSpanResponse, error) {
	resp := &OnsConsumerTimeSpanResponse{}
	err := client.Request("ons", "OnsConsumerTimeSpan", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsConsumerTimeSpanResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsConsumerTimeSpanData
}

type OnsConsumerTimeSpanData struct {
	Topic            goaliyun.String
	MinTimeStamp     goaliyun.Integer
	MaxTimeStamp     goaliyun.Integer
	ConsumeTimeStamp goaliyun.Integer
}
