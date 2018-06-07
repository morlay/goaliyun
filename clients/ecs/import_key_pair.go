package ecs

import (
	"github.com/morlay/goaliyun"
)

type ImportKeyPairRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PublicKeyBody        string `position:"Query" name:"PublicKeyBody"`
	KeyPairName          string `position:"Query" name:"KeyPairName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ImportKeyPairRequest) Invoke(client goaliyun.Client) (*ImportKeyPairResponse, error) {
	resp := &ImportKeyPairResponse{}
	err := client.Request("ecs", "ImportKeyPair", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ImportKeyPairResponse struct {
	RequestId          goaliyun.String
	KeyPairName        goaliyun.String
	KeyPairFingerPrint goaliyun.String
}
