package emr

import (
	"github.com/morlay/goaliyun"
)

type ResumeExecutionPlanInstanceRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ResumeExecutionPlanInstanceRequest) Invoke(client goaliyun.Client) (*ResumeExecutionPlanInstanceResponse, error) {
	resp := &ResumeExecutionPlanInstanceResponse{}
	err := client.Request("emr", "ResumeExecutionPlanInstance", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResumeExecutionPlanInstanceResponse struct {
	RequestId goaliyun.String
}
