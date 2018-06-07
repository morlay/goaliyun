package csb

import (
	"github.com/morlay/goaliyun"
)

type ApproveOrderListRequest struct {
	Data     string `position:"Body" name:"Data"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ApproveOrderListRequest) Invoke(client goaliyun.Client) (*ApproveOrderListResponse, error) {
	resp := &ApproveOrderListResponse{}
	err := client.Request("csb", "ApproveOrderList", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ApproveOrderListResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
}
