package emr

import (
	"github.com/morlay/goaliyun"
)

type ModifyJobExecutionPlanParamRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ParamName       string `position:"Query" name:"ParamName"`
	ParamValue      string `position:"Query" name:"ParamValue"`
	Id              int64  `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyJobExecutionPlanParamRequest) Invoke(client goaliyun.Client) (*ModifyJobExecutionPlanParamResponse, error) {
	resp := &ModifyJobExecutionPlanParamResponse{}
	err := client.Request("emr", "ModifyJobExecutionPlanParam", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyJobExecutionPlanParamResponse struct {
	RequestId goaliyun.String
	Success   goaliyun.String
	ErrCode   goaliyun.String
	ErrMsg    goaliyun.String
}
