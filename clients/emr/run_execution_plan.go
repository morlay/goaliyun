package emr

import (
	"github.com/morlay/goaliyun"
)

type RunExecutionPlanRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Arguments       string `position:"Query" name:"Arguments"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *RunExecutionPlanRequest) Invoke(client goaliyun.Client) (*RunExecutionPlanResponse, error) {
	resp := &RunExecutionPlanResponse{}
	err := client.Request("emr", "RunExecutionPlan", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RunExecutionPlanResponse struct {
	RequestId               goaliyun.String
	ExecutionPlanInstanceId goaliyun.String
}
