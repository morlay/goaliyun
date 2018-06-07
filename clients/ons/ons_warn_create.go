package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsWarnCreateRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	BlockTime    string `position:"Query" name:"BlockTime"`
	Level        string `position:"Query" name:"Level"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	DelayTime    string `position:"Query" name:"DelayTime"`
	Topic        string `position:"Query" name:"Topic"`
	Threshold    string `position:"Query" name:"Threshold"`
	AlertTime    string `position:"Query" name:"AlertTime"`
	Contacts     string `position:"Query" name:"Contacts"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsWarnCreateRequest) Invoke(client goaliyun.Client) (*OnsWarnCreateResponse, error) {
	resp := &OnsWarnCreateResponse{}
	err := client.Request("ons", "OnsWarnCreate", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsWarnCreateResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
}
