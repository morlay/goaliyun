package emr

import (
	"github.com/morlay/goaliyun"
)

type RetryExecutionPlanRequest struct {
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	ExecutionPlanWorkNodeIds string `position:"Query" name:"ExecutionPlanWorkNodeIds"`
	Id                       string `position:"Query" name:"Id"`
	RegionId                 string `position:"Query" name:"RegionId"`
}

func (req *RetryExecutionPlanRequest) Invoke(client goaliyun.Client) (*RetryExecutionPlanResponse, error) {
	resp := &RetryExecutionPlanResponse{}
	err := client.Request("emr", "RetryExecutionPlan", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RetryExecutionPlanResponse struct {
	RequestId               goaliyun.String
	ExecutionPlanInstanceId goaliyun.String
}
