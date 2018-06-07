package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type ModifyReplicaDescriptionRequest struct {
	ReplicaDescription   string `position:"Query" name:"ReplicaDescription"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyReplicaDescriptionRequest) Invoke(client goaliyun.Client) (*ModifyReplicaDescriptionResponse, error) {
	resp := &ModifyReplicaDescriptionResponse{}
	err := client.Request("r-kvstore", "ModifyReplicaDescription", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyReplicaDescriptionResponse struct {
	RequestId goaliyun.String
}
