package commondriver

import (
	"github.com/morlay/goaliyun"
)

type GetOrderIdByQueryPurchaseRequest struct {
	OrderId  string `position:"Query" name:"OrderId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetOrderIdByQueryPurchaseRequest) Invoke(client goaliyun.Client) (*GetOrderIdByQueryPurchaseResponse, error) {
	resp := &GetOrderIdByQueryPurchaseResponse{}
	err := client.Request("commondriver", "GetOrderIdByQueryPurchase", "2015-12-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetOrderIdByQueryPurchaseResponse struct {
	RequestId goaliyun.String
	Success   bool
	Data      bool
}
