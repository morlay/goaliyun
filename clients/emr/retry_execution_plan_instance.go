package emr

import (
	"github.com/morlay/goaliyun"
)

type RetryExecutionPlanInstanceRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Arguments       string `position:"Query" name:"Arguments"`
	Id              string `position:"Query" name:"Id"`
	RerunFail       string `position:"Query" name:"RerunFail"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *RetryExecutionPlanInstanceRequest) Invoke(client goaliyun.Client) (*RetryExecutionPlanInstanceResponse, error) {
	resp := &RetryExecutionPlanInstanceResponse{}
	err := client.Request("emr", "RetryExecutionPlanInstance", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RetryExecutionPlanInstanceResponse struct {
	RequestId goaliyun.String
}
