package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type ModifyCertificationRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NoCertification      string `position:"Query" name:"NoCertification"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyCertificationRequest) Invoke(client goaliyun.Client) (*ModifyCertificationResponse, error) {
	resp := &ModifyCertificationResponse{}
	err := client.Request("r-kvstore", "ModifyCertification", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyCertificationResponse struct {
	RequestId goaliyun.String
}
