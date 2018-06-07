package ccc

import (
	"github.com/morlay/goaliyun"
)

type SubmitBatchJobsRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	JobGroupId string `position:"Query" name:"JobGroupId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *SubmitBatchJobsRequest) Invoke(client goaliyun.Client) (*SubmitBatchJobsResponse, error) {
	resp := &SubmitBatchJobsResponse{}
	err := client.Request("ccc", "SubmitBatchJobs", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitBatchJobsResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
}
