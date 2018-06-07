package emr

import (
	"github.com/morlay/goaliyun"
)

type ResumeExecutionPlanSchedulerRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ResumeExecutionPlanSchedulerRequest) Invoke(client goaliyun.Client) (*ResumeExecutionPlanSchedulerResponse, error) {
	resp := &ResumeExecutionPlanSchedulerResponse{}
	err := client.Request("emr", "ResumeExecutionPlanScheduler", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResumeExecutionPlanSchedulerResponse struct {
	RequestId goaliyun.String
}
