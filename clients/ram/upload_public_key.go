package ram

import (
	"github.com/morlay/goaliyun"
)

type UploadPublicKeyRequest struct {
	PublicKeySpec string `position:"Query" name:"PublicKeySpec"`
	UserName      string `position:"Query" name:"UserName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *UploadPublicKeyRequest) Invoke(client goaliyun.Client) (*UploadPublicKeyResponse, error) {
	resp := &UploadPublicKeyResponse{}
	err := client.Request("ram", "UploadPublicKey", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UploadPublicKeyResponse struct {
	RequestId goaliyun.String
	PublicKey UploadPublicKeyPublicKey
}

type UploadPublicKeyPublicKey struct {
	PublicKeyId   goaliyun.String
	PublicKeySpec goaliyun.String
	Status        goaliyun.String
	CreateDate    goaliyun.String
}
