package ram

import (
	"github.com/morlay/goaliyun"
)

type DeletePublicKeyRequest struct {
	UserPublicKeyId string `position:"Query" name:"UserPublicKeyId"`
	UserName        string `position:"Query" name:"UserName"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeletePublicKeyRequest) Invoke(client goaliyun.Client) (*DeletePublicKeyResponse, error) {
	resp := &DeletePublicKeyResponse{}
	err := client.Request("ram", "DeletePublicKey", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeletePublicKeyResponse struct {
	RequestId goaliyun.String
}
