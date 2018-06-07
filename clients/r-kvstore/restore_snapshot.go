package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type RestoreSnapshotRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RestoreSnapshotRequest) Invoke(client goaliyun.Client) (*RestoreSnapshotResponse, error) {
	resp := &RestoreSnapshotResponse{}
	err := client.Request("r-kvstore", "RestoreSnapshot", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RestoreSnapshotResponse struct {
	RequestId goaliyun.String
}
