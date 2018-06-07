package csb

import (
	"github.com/morlay/goaliyun"
)

type DeleteCasServiceRequest struct {
	LeafOnly     string `position:"Query" name:"LeafOnly"`
	CasCsbName   string `position:"Query" name:"CasCsbName"`
	SrcUserId    string `position:"Query" name:"SrcUserId"`
	CasServiceId string `position:"Query" name:"CasServiceId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DeleteCasServiceRequest) Invoke(client goaliyun.Client) (*DeleteCasServiceResponse, error) {
	resp := &DeleteCasServiceResponse{}
	err := client.Request("csb", "DeleteCasService", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCasServiceResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
}
