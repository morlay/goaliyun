package market_inner

import (
	"github.com/morlay/goaliyun"
)

type InnerVerifyOrderRequest struct {
	Data     string `position:"Query" name:"Data"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *InnerVerifyOrderRequest) Invoke(client goaliyun.Client) (*InnerVerifyOrderResponse, error) {
	resp := &InnerVerifyOrderResponse{}
	err := client.Request("market-inner", "InnerVerifyOrder", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InnerVerifyOrderResponse struct {
	Success   bool
	Data      goaliyun.String
	RequestId goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
}
