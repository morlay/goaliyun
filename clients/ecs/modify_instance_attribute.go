package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyInstanceAttributeRequest struct {
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Password             string `position:"Query" name:"Password"`
	HostName             string `position:"Query" name:"HostName"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	InstanceName         string `position:"Query" name:"InstanceName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Recyclable           string `position:"Query" name:"Recyclable"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstanceAttributeRequest) Invoke(client goaliyun.Client) (*ModifyInstanceAttributeResponse, error) {
	resp := &ModifyInstanceAttributeResponse{}
	err := client.Request("ecs", "ModifyInstanceAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstanceAttributeResponse struct {
	RequestId goaliyun.String
}
