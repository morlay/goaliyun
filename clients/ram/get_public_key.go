package ram

import (
	"github.com/morlay/goaliyun"
)

type GetPublicKeyRequest struct {
	UserPublicKeyId string `position:"Query" name:"UserPublicKeyId"`
	UserName        string `position:"Query" name:"UserName"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GetPublicKeyRequest) Invoke(client goaliyun.Client) (*GetPublicKeyResponse, error) {
	resp := &GetPublicKeyResponse{}
	err := client.Request("ram", "GetPublicKey", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetPublicKeyResponse struct {
	RequestId goaliyun.String
	PublicKey GetPublicKeyPublicKey
}

type GetPublicKeyPublicKey struct {
	PublicKeyId   goaliyun.String
	PublicKeySpec goaliyun.String
	Status        goaliyun.String
	CreateDate    goaliyun.String
}
