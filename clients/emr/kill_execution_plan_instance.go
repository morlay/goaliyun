package emr

import (
	"github.com/morlay/goaliyun"
)

type KillExecutionPlanInstanceRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *KillExecutionPlanInstanceRequest) Invoke(client goaliyun.Client) (*KillExecutionPlanInstanceResponse, error) {
	resp := &KillExecutionPlanInstanceResponse{}
	err := client.Request("emr", "KillExecutionPlanInstance", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type KillExecutionPlanInstanceResponse struct {
	RequestId goaliyun.String
}
