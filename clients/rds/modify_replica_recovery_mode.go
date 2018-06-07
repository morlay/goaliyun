package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyReplicaRecoveryModeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	RecoveryMode         string `position:"Query" name:"RecoveryMode"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyReplicaRecoveryModeRequest) Invoke(client goaliyun.Client) (*ModifyReplicaRecoveryModeResponse, error) {
	resp := &ModifyReplicaRecoveryModeResponse{}
	err := client.Request("rds", "ModifyReplicaRecoveryMode", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyReplicaRecoveryModeResponse struct {
	RequestId goaliyun.String
}
