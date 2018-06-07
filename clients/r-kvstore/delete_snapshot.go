package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type DeleteSnapshotRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteSnapshotRequest) Invoke(client goaliyun.Client) (*DeleteSnapshotResponse, error) {
	resp := &DeleteSnapshotResponse{}
	err := client.Request("r-kvstore", "DeleteSnapshot", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteSnapshotResponse struct {
	RequestId goaliyun.String
}
