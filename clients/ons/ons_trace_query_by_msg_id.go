package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsTraceQueryByMsgIdRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	MsgId        string `position:"Query" name:"MsgId"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	BeginTime    int64  `position:"Query" name:"BeginTime"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsTraceQueryByMsgIdRequest) Invoke(client goaliyun.Client) (*OnsTraceQueryByMsgIdResponse, error) {
	resp := &OnsTraceQueryByMsgIdResponse{}
	err := client.Request("ons", "OnsTraceQueryByMsgId", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsTraceQueryByMsgIdResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	QueryId   goaliyun.String
}
