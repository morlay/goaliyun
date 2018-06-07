package emr

import (
	"github.com/morlay/goaliyun"
)

type ModifyExecutionPlanScheduleInfoRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	TimeInterval    int64  `position:"Query" name:"TimeInterval"`
	DayOfWeek       string `position:"Query" name:"DayOfWeek"`
	Id              string `position:"Query" name:"Id"`
	StartTime       int64  `position:"Query" name:"StartTime"`
	Strategy        string `position:"Query" name:"Strategy"`
	TimeUnit        string `position:"Query" name:"TimeUnit"`
	DayOfMonth      string `position:"Query" name:"DayOfMonth"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyExecutionPlanScheduleInfoRequest) Invoke(client goaliyun.Client) (*ModifyExecutionPlanScheduleInfoResponse, error) {
	resp := &ModifyExecutionPlanScheduleInfoResponse{}
	err := client.Request("emr", "ModifyExecutionPlanScheduleInfo", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyExecutionPlanScheduleInfoResponse struct {
	RequestId goaliyun.String
}
