package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type ModifyInstanceMinorVersionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Minorversion         string `position:"Query" name:"Minorversion"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstanceMinorVersionRequest) Invoke(client goaliyun.Client) (*ModifyInstanceMinorVersionResponse, error) {
	resp := &ModifyInstanceMinorVersionResponse{}
	err := client.Request("r-kvstore", "ModifyInstanceMinorVersion", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstanceMinorVersionResponse struct {
	RequestId goaliyun.String
}
