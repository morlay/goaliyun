package ess

import (
	"github.com/morlay/goaliyun"
)

type CreateScheduledTaskRequest struct {
	LaunchTime           string `position:"Query" name:"LaunchTime"`
	ScheduledAction      string `position:"Query" name:"ScheduledAction"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RecurrenceValue      string `position:"Query" name:"RecurrenceValue"`
	LaunchExpirationTime int64  `position:"Query" name:"LaunchExpirationTime"`
	RecurrenceEndTime    string `position:"Query" name:"RecurrenceEndTime"`
	ScheduledTaskName    string `position:"Query" name:"ScheduledTaskName"`
	TaskEnabled          string `position:"Query" name:"TaskEnabled"`
	RecurrenceType       string `position:"Query" name:"RecurrenceType"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateScheduledTaskRequest) Invoke(client goaliyun.Client) (*CreateScheduledTaskResponse, error) {
	resp := &CreateScheduledTaskResponse{}
	err := client.Request("ess", "CreateScheduledTask", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateScheduledTaskResponse struct {
	ScheduledTaskId goaliyun.String
	RequestId       goaliyun.String
}
