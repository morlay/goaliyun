package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type ModifyGuardDomainModeRequest struct {
	DomainMode           string `position:"Query" name:"DomainMode"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyGuardDomainModeRequest) Invoke(client goaliyun.Client) (*ModifyGuardDomainModeResponse, error) {
	resp := &ModifyGuardDomainModeResponse{}
	err := client.Request("r-kvstore", "ModifyGuardDomainMode", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyGuardDomainModeResponse struct {
	RequestId goaliyun.String
}
