package ecs

import (
	"github.com/morlay/goaliyun"
)

type CreateSnapshotRequest struct {
	Tag4Value            string `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Tag2Key              string `position:"Query" name:"Tag.2.Key"`
	Tag5Key              string `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	SnapshotName         string `position:"Query" name:"SnapshotName"`
	Tag3Key              string `position:"Query" name:"Tag.3.Key"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tag5Value            string `position:"Query" name:"Tag.5.Value"`
	Tag1Key              string `position:"Query" name:"Tag.1.Key"`
	Tag1Value            string `position:"Query" name:"Tag.1.Value"`
	Tag2Value            string `position:"Query" name:"Tag.2.Value"`
	Tag4Key              string `position:"Query" name:"Tag.4.Key"`
	DiskId               string `position:"Query" name:"DiskId"`
	Tag3Value            string `position:"Query" name:"Tag.3.Value"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateSnapshotRequest) Invoke(client goaliyun.Client) (*CreateSnapshotResponse, error) {
	resp := &CreateSnapshotResponse{}
	err := client.Request("ecs", "CreateSnapshot", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateSnapshotResponse struct {
	RequestId  goaliyun.String
	SnapshotId goaliyun.String
}
