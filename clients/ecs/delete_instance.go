package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteInstanceRequest struct {
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId            string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	TerminateSubscription string `position:"Query" name:"TerminateSubscription"`
	Force                 string `position:"Query" name:"Force"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *DeleteInstanceRequest) Invoke(client goaliyun.Client) (*DeleteInstanceResponse, error) {
	resp := &DeleteInstanceResponse{}
	err := client.Request("ecs", "DeleteInstance", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteInstanceResponse struct {
	RequestId goaliyun.String
}
