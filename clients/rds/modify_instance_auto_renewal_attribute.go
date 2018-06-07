package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyInstanceAutoRenewalAttributeRequest struct {
	Duration             string `position:"Query" name:"Duration"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AutoRenew            string `position:"Query" name:"AutoRenew"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstanceAutoRenewalAttributeRequest) Invoke(client goaliyun.Client) (*ModifyInstanceAutoRenewalAttributeResponse, error) {
	resp := &ModifyInstanceAutoRenewalAttributeResponse{}
	err := client.Request("rds", "ModifyInstanceAutoRenewalAttribute", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstanceAutoRenewalAttributeResponse struct {
	RequestId goaliyun.String
}
