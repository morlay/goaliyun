package mts

import (
	"github.com/morlay/goaliyun"
)

type DecryptKeyRequest struct {
	Rand                 string `position:"Query" name:"Rand"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	CiphertextBlob       string `position:"Query" name:"CiphertextBlob"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DecryptKeyRequest) Invoke(client goaliyun.Client) (*DecryptKeyResponse, error) {
	resp := &DecryptKeyResponse{}
	err := client.Request("mts", "DecryptKey", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DecryptKeyResponse struct {
	RequestId goaliyun.String
	Plaintext goaliyun.String
	Rand      goaliyun.String
}
