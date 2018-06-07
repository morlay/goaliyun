package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsTraceQueryByMsgKeyRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	BeginTime    int64  `position:"Query" name:"BeginTime"`
	MsgKey       string `position:"Query" name:"MsgKey"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsTraceQueryByMsgKeyRequest) Invoke(client goaliyun.Client) (*OnsTraceQueryByMsgKeyResponse, error) {
	resp := &OnsTraceQueryByMsgKeyResponse{}
	err := client.Request("ons", "OnsTraceQueryByMsgKey", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsTraceQueryByMsgKeyResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	QueryId   goaliyun.String
}
