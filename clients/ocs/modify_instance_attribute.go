package ocs

import (
	"github.com/morlay/goaliyun"
)

type ModifyInstanceAttributeRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	OcsInstanceName      string `position:"Query" name:"OcsInstanceName"`
	OldPassword          string `position:"Query" name:"OldPassword"`
	NewPassword          string `position:"Query" name:"NewPassword"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstanceAttributeRequest) Invoke(client goaliyun.Client) (*ModifyInstanceAttributeResponse, error) {
	resp := &ModifyInstanceAttributeResponse{}
	err := client.Request("ocs", "ModifyInstanceAttribute", "2015-03-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstanceAttributeResponse struct {
	RequestId goaliyun.String
}
