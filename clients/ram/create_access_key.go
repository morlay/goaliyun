package ram

import (
	"github.com/morlay/goaliyun"
)

type CreateAccessKeyRequest struct {
	UserName string `position:"Query" name:"UserName"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *CreateAccessKeyRequest) Invoke(client goaliyun.Client) (*CreateAccessKeyResponse, error) {
	resp := &CreateAccessKeyResponse{}
	err := client.Request("ram", "CreateAccessKey", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateAccessKeyResponse struct {
	RequestId goaliyun.String
	AccessKey CreateAccessKeyAccessKey
}

type CreateAccessKeyAccessKey struct {
	AccessKeyId     goaliyun.String
	AccessKeySecret goaliyun.String
	Status          goaliyun.String
	CreateDate      goaliyun.String
}
