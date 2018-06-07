package emr

import (
	"github.com/morlay/goaliyun"
)

type ResumeFlowRequest struct {
	FlowInstanceId  string `position:"Query" name:"FlowInstanceId"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ResumeFlowRequest) Invoke(client goaliyun.Client) (*ResumeFlowResponse, error) {
	resp := &ResumeFlowResponse{}
	err := client.Request("emr", "ResumeFlow", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResumeFlowResponse struct {
	RequestId goaliyun.String
	Data      bool
}
