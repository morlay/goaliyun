package emr

import (
	"github.com/morlay/goaliyun"
)

type SuspendFlowRequest struct {
	FlowInstanceId  string `position:"Query" name:"FlowInstanceId"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *SuspendFlowRequest) Invoke(client goaliyun.Client) (*SuspendFlowResponse, error) {
	resp := &SuspendFlowResponse{}
	err := client.Request("emr", "SuspendFlow", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SuspendFlowResponse struct {
	RequestId goaliyun.String
	Data      bool
}
