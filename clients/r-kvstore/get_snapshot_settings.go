package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type GetSnapshotSettingsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GetSnapshotSettingsRequest) Invoke(client goaliyun.Client) (*GetSnapshotSettingsResponse, error) {
	resp := &GetSnapshotSettingsResponse{}
	err := client.Request("r-kvstore", "GetSnapshotSettings", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetSnapshotSettingsResponse struct {
	RequestId          goaliyun.String
	InstanceId         goaliyun.String
	BeginHour          goaliyun.Integer
	EndHour            goaliyun.Integer
	RetentionDay       goaliyun.Integer
	MaxAutoSnapshots   goaliyun.Integer
	MaxManualSnapshots goaliyun.Integer
	DayList            goaliyun.Integer
	NextTime           goaliyun.String
}
