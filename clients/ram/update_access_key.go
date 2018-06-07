package ram

import (
	"github.com/morlay/goaliyun"
)

type UpdateAccessKeyRequest struct {
	UserAccessKeyId string `position:"Query" name:"UserAccessKeyId"`
	UserName        string `position:"Query" name:"UserName"`
	Status          string `position:"Query" name:"Status"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *UpdateAccessKeyRequest) Invoke(client goaliyun.Client) (*UpdateAccessKeyResponse, error) {
	resp := &UpdateAccessKeyResponse{}
	err := client.Request("ram", "UpdateAccessKey", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateAccessKeyResponse struct {
	RequestId goaliyun.String
}
