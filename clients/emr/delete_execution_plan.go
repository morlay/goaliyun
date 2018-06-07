package emr

import (
	"github.com/morlay/goaliyun"
)

type DeleteExecutionPlanRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteExecutionPlanRequest) Invoke(client goaliyun.Client) (*DeleteExecutionPlanResponse, error) {
	resp := &DeleteExecutionPlanResponse{}
	err := client.Request("emr", "DeleteExecutionPlan", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteExecutionPlanResponse struct {
	RequestId goaliyun.String
}
