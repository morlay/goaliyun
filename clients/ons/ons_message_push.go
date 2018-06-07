package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsMessagePushRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	ClientId     string `position:"Query" name:"ClientId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	MsgId        string `position:"Query" name:"MsgId"`
	Topic        string `position:"Query" name:"Topic"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsMessagePushRequest) Invoke(client goaliyun.Client) (*OnsMessagePushResponse, error) {
	resp := &OnsMessagePushResponse{}
	err := client.Request("ons", "OnsMessagePush", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsMessagePushResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
}
