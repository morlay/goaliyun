package kms

import (
	"github.com/morlay/goaliyun"
)

type ImportKeyMaterialRequest struct {
	ImportToken           string `position:"Query" name:"ImportToken"`
	EncryptedKeyMaterial  string `position:"Query" name:"EncryptedKeyMaterial"`
	KeyMaterialExpireUnix int64  `position:"Query" name:"KeyMaterialExpireUnix"`
	KeyId                 string `position:"Query" name:"KeyId"`
	STSToken              string `position:"Query" name:"STSToken"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *ImportKeyMaterialRequest) Invoke(client goaliyun.Client) (*ImportKeyMaterialResponse, error) {
	resp := &ImportKeyMaterialResponse{}
	err := client.Request("kms", "ImportKeyMaterial", "2016-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ImportKeyMaterialResponse struct {
	RequestId goaliyun.String
}
