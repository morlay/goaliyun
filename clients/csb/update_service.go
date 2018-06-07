package csb

import (
	"github.com/morlay/goaliyun"
)

type UpdateServiceRequest struct {
	Data     string `position:"Body" name:"Data"`
	CsbId    int64  `position:"Query" name:"CsbId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UpdateServiceRequest) Invoke(client goaliyun.Client) (*UpdateServiceResponse, error) {
	resp := &UpdateServiceResponse{}
	err := client.Request("csb", "UpdateService", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateServiceResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
}
