package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type CreateTempInstanceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateTempInstanceRequest) Invoke(client goaliyun.Client) (*CreateTempInstanceResponse, error) {
	resp := &CreateTempInstanceResponse{}
	err := client.Request("r-kvstore", "CreateTempInstance", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateTempInstanceResponse struct {
	RequestId      goaliyun.String
	InstanceId     goaliyun.String
	SnapshotId     goaliyun.String
	TempInstanceId goaliyun.String
	Status         goaliyun.String
}
