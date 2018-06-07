package emr

import (
	"github.com/morlay/goaliyun"
)

type SuspendExecutionPlanInstanceRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *SuspendExecutionPlanInstanceRequest) Invoke(client goaliyun.Client) (*SuspendExecutionPlanInstanceResponse, error) {
	resp := &SuspendExecutionPlanInstanceResponse{}
	err := client.Request("emr", "SuspendExecutionPlanInstance", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SuspendExecutionPlanInstanceResponse struct {
	RequestId goaliyun.String
}
