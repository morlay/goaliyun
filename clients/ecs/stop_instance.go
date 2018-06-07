package ecs

import (
	"github.com/morlay/goaliyun"
)

type StopInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ConfirmStop          string `position:"Query" name:"ConfirmStop"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	StoppedMode          string `position:"Query" name:"StoppedMode"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ForceStop            string `position:"Query" name:"ForceStop"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *StopInstanceRequest) Invoke(client goaliyun.Client) (*StopInstanceResponse, error) {
	resp := &StopInstanceResponse{}
	err := client.Request("ecs", "StopInstance", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StopInstanceResponse struct {
	RequestId goaliyun.String
}
