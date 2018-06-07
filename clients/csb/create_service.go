package csb

import (
	"github.com/morlay/goaliyun"
)

type CreateServiceRequest struct {
	Data     string `position:"Body" name:"Data"`
	CsbId    int64  `position:"Query" name:"CsbId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *CreateServiceRequest) Invoke(client goaliyun.Client) (*CreateServiceResponse, error) {
	resp := &CreateServiceResponse{}
	err := client.Request("csb", "CreateService", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateServiceResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      CreateServiceData
}

type CreateServiceData struct {
	Id goaliyun.Integer
}
