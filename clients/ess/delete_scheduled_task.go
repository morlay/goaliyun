package ess

import (
	"github.com/morlay/goaliyun"
)

type DeleteScheduledTaskRequest struct {
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ScheduledTaskId      string `position:"Query" name:"ScheduledTaskId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteScheduledTaskRequest) Invoke(client goaliyun.Client) (*DeleteScheduledTaskResponse, error) {
	resp := &DeleteScheduledTaskResponse{}
	err := client.Request("ess", "DeleteScheduledTask", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteScheduledTaskResponse struct {
	RequestId goaliyun.String
}
