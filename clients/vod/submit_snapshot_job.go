package vod

import (
	"github.com/morlay/goaliyun"
)

type SubmitSnapshotJobRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SpecifiedOffsetTime  int64  `position:"Query" name:"SpecifiedOffsetTime"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Width                string `position:"Query" name:"Width"`
	Count                int64  `position:"Query" name:"Count"`
	VideoId              string `position:"Query" name:"VideoId"`
	Interval             int64  `position:"Query" name:"Interval"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SpriteSnapshotConfig string `position:"Query" name:"SpriteSnapshotConfig"`
	Height               string `position:"Query" name:"Height"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitSnapshotJobRequest) Invoke(client goaliyun.Client) (*SubmitSnapshotJobResponse, error) {
	resp := &SubmitSnapshotJobResponse{}
	err := client.Request("vod", "SubmitSnapshotJob", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitSnapshotJobResponse struct {
	RequestId   goaliyun.String
	SnapshotJob SubmitSnapshotJobSnapshotJob
}

type SubmitSnapshotJobSnapshotJob struct {
	JobId goaliyun.String
}
