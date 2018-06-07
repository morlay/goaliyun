package emr

import (
	"github.com/morlay/goaliyun"
)

type SuspendExecutionPlanSchedulerRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *SuspendExecutionPlanSchedulerRequest) Invoke(client goaliyun.Client) (*SuspendExecutionPlanSchedulerResponse, error) {
	resp := &SuspendExecutionPlanSchedulerResponse{}
	err := client.Request("emr", "SuspendExecutionPlanScheduler", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SuspendExecutionPlanSchedulerResponse struct {
	RequestId goaliyun.String
}
