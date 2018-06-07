package csb

import (
	"github.com/morlay/goaliyun"
)

type UpdateServiceListStatusRequest struct {
	Data     string `position:"Body" name:"Data"`
	CsbId    int64  `position:"Query" name:"CsbId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UpdateServiceListStatusRequest) Invoke(client goaliyun.Client) (*UpdateServiceListStatusResponse, error) {
	resp := &UpdateServiceListStatusResponse{}
	err := client.Request("csb", "UpdateServiceListStatus", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateServiceListStatusResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
}
