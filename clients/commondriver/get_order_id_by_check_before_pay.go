package commondriver

import (
	"github.com/morlay/goaliyun"
)

type GetOrderIdByCheckBeforePayRequest struct {
	OrderId  string `position:"Query" name:"OrderId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetOrderIdByCheckBeforePayRequest) Invoke(client goaliyun.Client) (*GetOrderIdByCheckBeforePayResponse, error) {
	resp := &GetOrderIdByCheckBeforePayResponse{}
	err := client.Request("commondriver", "GetOrderIdByCheckBeforePay", "2015-12-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetOrderIdByCheckBeforePayResponse struct {
	RequestId goaliyun.String
	Success   bool
	Data      bool
}
