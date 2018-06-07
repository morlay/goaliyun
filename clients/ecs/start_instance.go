package ecs

import (
	"github.com/morlay/goaliyun"
)

type StartInstanceRequest struct {
	InitLocalDisk        string `position:"Query" name:"InitLocalDisk"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *StartInstanceRequest) Invoke(client goaliyun.Client) (*StartInstanceResponse, error) {
	resp := &StartInstanceResponse{}
	err := client.Request("ecs", "StartInstance", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type StartInstanceResponse struct {
	RequestId goaliyun.String
}
