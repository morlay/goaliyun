package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteKeyPairsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	KeyPairNames         string `position:"Query" name:"KeyPairNames"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteKeyPairsRequest) Invoke(client goaliyun.Client) (*DeleteKeyPairsResponse, error) {
	resp := &DeleteKeyPairsResponse{}
	err := client.Request("ecs", "DeleteKeyPairs", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteKeyPairsResponse struct {
	RequestId goaliyun.String
}
