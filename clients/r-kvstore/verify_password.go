package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type VerifyPasswordRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Password             string `position:"Query" name:"Password"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *VerifyPasswordRequest) Invoke(client goaliyun.Client) (*VerifyPasswordResponse, error) {
	resp := &VerifyPasswordResponse{}
	err := client.Request("r-kvstore", "VerifyPassword", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type VerifyPasswordResponse struct {
	RequestId goaliyun.String
}
