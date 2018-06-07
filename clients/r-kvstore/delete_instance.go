package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type DeleteInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteInstanceRequest) Invoke(client goaliyun.Client) (*DeleteInstanceResponse, error) {
	resp := &DeleteInstanceResponse{}
	err := client.Request("r-kvstore", "DeleteInstance", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteInstanceResponse struct {
	RequestId goaliyun.String
}
