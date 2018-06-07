package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyReplicaVerificationModeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	VerificationMode     string `position:"Query" name:"VerificationMode"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyReplicaVerificationModeRequest) Invoke(client goaliyun.Client) (*ModifyReplicaVerificationModeResponse, error) {
	resp := &ModifyReplicaVerificationModeResponse{}
	err := client.Request("rds", "ModifyReplicaVerificationMode", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyReplicaVerificationModeResponse struct {
	RequestId goaliyun.String
}
