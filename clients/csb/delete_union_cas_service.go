package csb

import (
	"github.com/morlay/goaliyun"
)

type DeleteUnionCasServiceRequest struct {
	LeafOnly     string `position:"Query" name:"LeafOnly"`
	CasCsbName   string `position:"Query" name:"CasCsbName"`
	SrcUserId    string `position:"Query" name:"SrcUserId"`
	CasServiceId string `position:"Query" name:"CasServiceId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DeleteUnionCasServiceRequest) Invoke(client goaliyun.Client) (*DeleteUnionCasServiceResponse, error) {
	resp := &DeleteUnionCasServiceResponse{}
	err := client.Request("csb", "DeleteUnionCasService", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteUnionCasServiceResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
}
