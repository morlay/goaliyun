package csb

import (
	"github.com/morlay/goaliyun"
)

type UpdateOrderListRequest struct {
	Data     string `position:"Body" name:"Data"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UpdateOrderListRequest) Invoke(client goaliyun.Client) (*UpdateOrderListResponse, error) {
	resp := &UpdateOrderListResponse{}
	err := client.Request("csb", "UpdateOrderList", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateOrderListResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      UpdateOrderListData
}

type UpdateOrderListData struct {
	UpdateCount goaliyun.Integer
}
