package emr

import (
	"github.com/morlay/goaliyun"
)

type ModifyExecutionPlanBasicInfoRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	Id              string `position:"Query" name:"Id"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyExecutionPlanBasicInfoRequest) Invoke(client goaliyun.Client) (*ModifyExecutionPlanBasicInfoResponse, error) {
	resp := &ModifyExecutionPlanBasicInfoResponse{}
	err := client.Request("emr", "ModifyExecutionPlanBasicInfo", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyExecutionPlanBasicInfoResponse struct {
	RequestId goaliyun.String
}
