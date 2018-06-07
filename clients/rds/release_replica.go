package rds

import (
	"github.com/morlay/goaliyun"
)

type ReleaseReplicaRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ReleaseReplicaRequest) Invoke(client goaliyun.Client) (*ReleaseReplicaResponse, error) {
	resp := &ReleaseReplicaResponse{}
	err := client.Request("rds", "ReleaseReplica", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReleaseReplicaResponse struct {
	RequestId goaliyun.String
}
