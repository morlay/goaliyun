package csb

import (
	"github.com/morlay/goaliyun"
)

type DeleteOrderListRequest struct {
	Data     string `position:"Body" name:"Data"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteOrderListRequest) Invoke(client goaliyun.Client) (*DeleteOrderListResponse, error) {
	resp := &DeleteOrderListResponse{}
	err := client.Request("csb", "DeleteOrderList", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteOrderListResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
}
