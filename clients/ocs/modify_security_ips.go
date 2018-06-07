package ocs

import (
	"github.com/morlay/goaliyun"
)

type ModifySecurityIpsRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	SecurityIps          string `position:"Query" name:"SecurityIps"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifySecurityIpsRequest) Invoke(client goaliyun.Client) (*ModifySecurityIpsResponse, error) {
	resp := &ModifySecurityIpsResponse{}
	err := client.Request("ocs", "ModifySecurityIps", "2015-03-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifySecurityIpsResponse struct {
	RequestId goaliyun.String
}
