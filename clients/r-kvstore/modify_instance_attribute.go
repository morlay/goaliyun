package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type ModifyInstanceAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	InstanceName         string `position:"Query" name:"InstanceName"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NewPassword          string `position:"Query" name:"NewPassword"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstanceAttributeRequest) Invoke(client goaliyun.Client) (*ModifyInstanceAttributeResponse, error) {
	resp := &ModifyInstanceAttributeResponse{}
	err := client.Request("r-kvstore", "ModifyInstanceAttribute", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstanceAttributeResponse struct {
	RequestId goaliyun.String
}
