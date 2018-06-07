package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type DescribeStrategyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeStrategyRequest) Invoke(client goaliyun.Client) (*DescribeStrategyResponse, error) {
	resp := &DescribeStrategyResponse{}
	err := client.Request("r-kvstore", "DescribeStrategy", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeStrategyResponse struct {
	RequestId        goaliyun.String
	ReplicaId        goaliyun.String
	RecoveryMode     goaliyun.String
	VerificationMode goaliyun.String
}
