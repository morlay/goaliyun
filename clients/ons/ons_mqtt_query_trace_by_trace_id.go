package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsMqttQueryTraceByTraceIdRequest struct {
	TraceId      string `position:"Query" name:"TraceId"`
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsMqttQueryTraceByTraceIdRequest) Invoke(client goaliyun.Client) (*OnsMqttQueryTraceByTraceIdResponse, error) {
	resp := &OnsMqttQueryTraceByTraceIdResponse{}
	err := client.Request("ons", "OnsMqttQueryTraceByTraceId", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsMqttQueryTraceByTraceIdResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	PushCount goaliyun.Integer
}
