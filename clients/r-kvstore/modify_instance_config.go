package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type ModifyInstanceConfigRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Config               string `position:"Query" name:"Config"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstanceConfigRequest) Invoke(client goaliyun.Client) (*ModifyInstanceConfigResponse, error) {
	resp := &ModifyInstanceConfigResponse{}
	err := client.Request("r-kvstore", "ModifyInstanceConfig", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstanceConfigResponse struct {
	RequestId goaliyun.String
}
