package csb

import (
	"github.com/morlay/goaliyun"
)

type UpdateOrderRequest struct {
	Data     string `position:"Body" name:"Data"`
	CsbId    int64  `position:"Query" name:"CsbId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UpdateOrderRequest) Invoke(client goaliyun.Client) (*UpdateOrderResponse, error) {
	resp := &UpdateOrderResponse{}
	err := client.Request("csb", "UpdateOrder", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateOrderResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
}
