package ess

import (
	"github.com/morlay/goaliyun"
)

type ModifyScheduledTaskRequest struct {
	LaunchTime           string `position:"Query" name:"LaunchTime"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
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
	ScheduledTaskId      string `position:"Query" name:"ScheduledTaskId"`
	RecurrenceType       string `position:"Query" name:"RecurrenceType"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyScheduledTaskRequest) Invoke(client goaliyun.Client) (*ModifyScheduledTaskResponse, error) {
	resp := &ModifyScheduledTaskResponse{}
	err := client.Request("ess", "ModifyScheduledTask", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyScheduledTaskResponse struct {
	RequestId goaliyun.String
}
