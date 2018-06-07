package emr

import (
	"github.com/morlay/goaliyun"
)

type DeleteJobExecutionPlanParamRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              int64  `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteJobExecutionPlanParamRequest) Invoke(client goaliyun.Client) (*DeleteJobExecutionPlanParamResponse, error) {
	resp := &DeleteJobExecutionPlanParamResponse{}
	err := client.Request("emr", "DeleteJobExecutionPlanParam", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteJobExecutionPlanParamResponse struct {
	RequestId goaliyun.String
}
