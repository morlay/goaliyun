package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type SetSnapshotSettingsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SetSnapshotSettingsRequest) Invoke(client goaliyun.Client) (*SetSnapshotSettingsResponse, error) {
	resp := &SetSnapshotSettingsResponse{}
	err := client.Request("r-kvstore", "SetSnapshotSettings", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetSnapshotSettingsResponse struct {
	RequestId goaliyun.String
}
