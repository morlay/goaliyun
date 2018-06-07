package csb

import (
	"github.com/morlay/goaliyun"
)

type PublishUnionCasServiceRequest struct {
	CasCsbName string `position:"Query" name:"CasCsbName"`
	Data       string `position:"Body" name:"Data"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *PublishUnionCasServiceRequest) Invoke(client goaliyun.Client) (*PublishUnionCasServiceResponse, error) {
	resp := &PublishUnionCasServiceResponse{}
	err := client.Request("csb", "PublishUnionCasService", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PublishUnionCasServiceResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
}
