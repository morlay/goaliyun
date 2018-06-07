package csb

import (
	"github.com/morlay/goaliyun"
)

type PublishCasServiceRequest struct {
	CasCsbName string `position:"Query" name:"CasCsbName"`
	Data       string `position:"Body" name:"Data"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *PublishCasServiceRequest) Invoke(client goaliyun.Client) (*PublishCasServiceResponse, error) {
	resp := &PublishCasServiceResponse{}
	err := client.Request("csb", "PublishCasService", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PublishCasServiceResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
}
