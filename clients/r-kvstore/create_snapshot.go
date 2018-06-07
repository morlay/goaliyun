package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type CreateSnapshotRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateSnapshotRequest) Invoke(client goaliyun.Client) (*CreateSnapshotResponse, error) {
	resp := &CreateSnapshotResponse{}
	err := client.Request("r-kvstore", "CreateSnapshot", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateSnapshotResponse struct {
	RequestId  goaliyun.String
	SnapshotId goaliyun.String
	Status     goaliyun.String
}
