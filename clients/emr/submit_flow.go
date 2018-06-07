package emr

import (
	"github.com/morlay/goaliyun"
)

type SubmitFlowRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Conf            string `position:"Query" name:"Conf"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	FlowId          string `position:"Query" name:"FlowId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *SubmitFlowRequest) Invoke(client goaliyun.Client) (*SubmitFlowResponse, error) {
	resp := &SubmitFlowResponse{}
	err := client.Request("emr", "SubmitFlow", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitFlowResponse struct {
	RequestId goaliyun.String
	Id        goaliyun.String
}
