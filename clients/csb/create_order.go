package csb

import (
	"github.com/morlay/goaliyun"
)

type CreateOrderRequest struct {
	Data     string `position:"Body" name:"Data"`
	CsbId    int64  `position:"Query" name:"CsbId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *CreateOrderRequest) Invoke(client goaliyun.Client) (*CreateOrderResponse, error) {
	resp := &CreateOrderResponse{}
	err := client.Request("csb", "CreateOrder", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateOrderResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      CreateOrderData
}

type CreateOrderData struct {
	Id goaliyun.Integer
}
