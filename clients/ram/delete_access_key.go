package ram

import (
	"github.com/morlay/goaliyun"
)

type DeleteAccessKeyRequest struct {
	UserAccessKeyId string `position:"Query" name:"UserAccessKeyId"`
	UserName        string `position:"Query" name:"UserName"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteAccessKeyRequest) Invoke(client goaliyun.Client) (*DeleteAccessKeyResponse, error) {
	resp := &DeleteAccessKeyResponse{}
	err := client.Request("ram", "DeleteAccessKey", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteAccessKeyResponse struct {
	RequestId goaliyun.String
}
