package kms

import (
	"github.com/morlay/goaliyun"
)

type ScheduleKeyDeletionRequest struct {
	PendingWindowInDays int64  `position:"Query" name:"PendingWindowInDays"`
	KeyId               string `position:"Query" name:"KeyId"`
	STSToken            string `position:"Query" name:"STSToken"`
	RegionId            string `position:"Query" name:"RegionId"`
}

func (req *ScheduleKeyDeletionRequest) Invoke(client goaliyun.Client) (*ScheduleKeyDeletionResponse, error) {
	resp := &ScheduleKeyDeletionResponse{}
	err := client.Request("kms", "ScheduleKeyDeletion", "2016-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ScheduleKeyDeletionResponse struct {
	RequestId goaliyun.String
}
