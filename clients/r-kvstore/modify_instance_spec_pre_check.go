package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type ModifyInstanceSpecPreCheckRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TargetInstanceClass  string `position:"Query" name:"TargetInstanceClass"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstanceSpecPreCheckRequest) Invoke(client goaliyun.Client) (*ModifyInstanceSpecPreCheckResponse, error) {
	resp := &ModifyInstanceSpecPreCheckResponse{}
	err := client.Request("r-kvstore", "ModifyInstanceSpecPreCheck", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstanceSpecPreCheckResponse struct {
	RequestId       goaliyun.String
	IsAllowModify   bool
	DisableCommands goaliyun.String
}
