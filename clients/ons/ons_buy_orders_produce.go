package ons

import (
	"github.com/morlay/goaliyun"
)

type OnsBuyOrdersProduceRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Data         string `position:"Query" name:"Data"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsBuyOrdersProduceRequest) Invoke(client goaliyun.Client) (*OnsBuyOrdersProduceResponse, error) {
	resp := &OnsBuyOrdersProduceResponse{}
	err := client.Request("ons", "OnsBuyOrdersProduce", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsBuyOrdersProduceResponse struct {
	Success   bool
	RequestId goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
}
