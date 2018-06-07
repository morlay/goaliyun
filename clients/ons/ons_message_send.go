package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsMessageSendRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	ProducerId   string `position:"Query" name:"ProducerId"`
	Tag          string `position:"Query" name:"Tag"`
	Message      string `position:"Query" name:"Message"`
	Key          string `position:"Query" name:"Key"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsMessageSendRequest) Invoke(client goaliyun.Client) (*OnsMessageSendResponse, error) {
	resp := &OnsMessageSendResponse{}
	err := client.Request("ons", "OnsMessageSend", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsMessageSendResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      goaliyun.String
}
