package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifyReplicaModeRequest struct {
	DomainMode           string `position:"Query" name:"DomainMode"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	PrimaryInstanceId    string `position:"Query" name:"PrimaryInstanceId"`
	ReplicaMode          string `position:"Query" name:"ReplicaMode"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyReplicaModeRequest) Invoke(client goaliyun.Client) (*ModifyReplicaModeResponse, error) {
	resp := &ModifyReplicaModeResponse{}
	err := client.Request("rds", "ModifyReplicaMode", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyReplicaModeResponse struct {
	RequestId goaliyun.String
}
