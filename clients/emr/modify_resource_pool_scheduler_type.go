package emr

import (
	"github.com/morlay/goaliyun"
)

type ModifyResourcePoolSchedulerTypeRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SchedulerType   string `position:"Query" name:"SchedulerType"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyResourcePoolSchedulerTypeRequest) Invoke(client goaliyun.Client) (*ModifyResourcePoolSchedulerTypeResponse, error) {
	resp := &ModifyResourcePoolSchedulerTypeResponse{}
	err := client.Request("emr", "ModifyResourcePoolSchedulerType", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyResourcePoolSchedulerTypeResponse struct {
	RequestId goaliyun.String
}
