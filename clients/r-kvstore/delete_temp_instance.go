package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type DeleteTempInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteTempInstanceRequest) Invoke(client goaliyun.Client) (*DeleteTempInstanceResponse, error) {
	resp := &DeleteTempInstanceResponse{}
	err := client.Request("r-kvstore", "DeleteTempInstance", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteTempInstanceResponse struct {
	RequestId goaliyun.String
}
