package ram

import (
	"github.com/morlay/goaliyun"
)

type UpdatePublicKeyRequest struct {
	UserPublicKeyId string `position:"Query" name:"UserPublicKeyId"`
	UserName        string `position:"Query" name:"UserName"`
	Status          string `position:"Query" name:"Status"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *UpdatePublicKeyRequest) Invoke(client goaliyun.Client) (*UpdatePublicKeyResponse, error) {
	resp := &UpdatePublicKeyResponse{}
	err := client.Request("ram", "UpdatePublicKey", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdatePublicKeyResponse struct {
	RequestId goaliyun.String
}
