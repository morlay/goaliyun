package emr

import (
	"github.com/morlay/goaliyun"
)

type RerunFlowRequest struct {
	FlowInstanceId  string `position:"Query" name:"FlowInstanceId"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	ReRunFail       string `position:"Query" name:"ReRunFail"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *RerunFlowRequest) Invoke(client goaliyun.Client) (*RerunFlowResponse, error) {
	resp := &RerunFlowResponse{}
	err := client.Request("emr", "RerunFlow", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RerunFlowResponse struct {
	RequestId goaliyun.String
	Data      bool
}
