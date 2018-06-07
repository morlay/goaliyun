package ecs

import (
	"github.com/morlay/goaliyun"
)

type CreateKeyPairRequest struct {
	Tag4Value            string `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Tag2Key              string `position:"Query" name:"Tag.2.Key"`
	Tag5Key              string `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Tag3Key              string `position:"Query" name:"Tag.3.Key"`
	KeyPairName          string `position:"Query" name:"KeyPairName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tag5Value            string `position:"Query" name:"Tag.5.Value"`
	Tag1Key              string `position:"Query" name:"Tag.1.Key"`
	Tag1Value            string `position:"Query" name:"Tag.1.Value"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	Tag2Value            string `position:"Query" name:"Tag.2.Value"`
	Tag4Key              string `position:"Query" name:"Tag.4.Key"`
	Tag3Value            string `position:"Query" name:"Tag.3.Value"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateKeyPairRequest) Invoke(client goaliyun.Client) (*CreateKeyPairResponse, error) {
	resp := &CreateKeyPairResponse{}
	err := client.Request("ecs", "CreateKeyPair", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateKeyPairResponse struct {
	RequestId          goaliyun.String
	KeyPairName        goaliyun.String
	KeyPairFingerPrint goaliyun.String
	PrivateKeyBody     goaliyun.String
}
