package ecs

import (
	"github.com/morlay/goaliyun"
)

type CancelTaskRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TaskId               string `position:"Query" name:"TaskId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CancelTaskRequest) Invoke(client goaliyun.Client) (*CancelTaskResponse, error) {
	resp := &CancelTaskResponse{}
	err := client.Request("ecs", "CancelTask", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CancelTaskResponse struct {
	RequestId goaliyun.String
}
