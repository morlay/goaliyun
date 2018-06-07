package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type SwitchTempInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SwitchTempInstanceRequest) Invoke(client goaliyun.Client) (*SwitchTempInstanceResponse, error) {
	resp := &SwitchTempInstanceResponse{}
	err := client.Request("r-kvstore", "SwitchTempInstance", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SwitchTempInstanceResponse struct {
	RequestId goaliyun.String
}
