package mts

import (
	"github.com/morlay/goaliyun"
)

type SubmitSnapshotJobRequest struct {
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SnapshotConfig       string `position:"Query" name:"SnapshotConfig"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SubmitSnapshotJobRequest) Invoke(client goaliyun.Client) (*SubmitSnapshotJobResponse, error) {
	resp := &SubmitSnapshotJobResponse{}
	err := client.Request("mts", "SubmitSnapshotJob", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SubmitSnapshotJobResponse struct {
	RequestId   goaliyun.String
	SnapshotJob SubmitSnapshotJobSnapshotJob
}

type SubmitSnapshotJobSnapshotJob struct {
	Id               goaliyun.String
	UserData         goaliyun.String
	PipelineId       goaliyun.String
	State            goaliyun.String
	Code             goaliyun.String
	Count            goaliyun.String
	TileCount        goaliyun.String
	Message          goaliyun.String
	CreationTime     goaliyun.String
	Input            SubmitSnapshotJobInput
	SnapshotConfig   SubmitSnapshotJobSnapshotConfig
	MNSMessageResult SubmitSnapshotJobMNSMessageResult
}

type SubmitSnapshotJobInput struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
	RoleArn  goaliyun.String
}

type SubmitSnapshotJobSnapshotConfig struct {
	Time           goaliyun.String
	Interval       goaliyun.String
	Num            goaliyun.String
	Width          goaliyun.String
	Height         goaliyun.String
	FrameType      goaliyun.String
	OutputFile     SubmitSnapshotJobOutputFile
	TileOutputFile SubmitSnapshotJobTileOutputFile
	TileOut        SubmitSnapshotJobTileOut
}

type SubmitSnapshotJobOutputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
	RoleArn  goaliyun.String
}

type SubmitSnapshotJobTileOutputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
	RoleArn  goaliyun.String
}

type SubmitSnapshotJobTileOut struct {
	Lines         goaliyun.String
	Columns       goaliyun.String
	CellWidth     goaliyun.String
	CellHeight    goaliyun.String
	Margin        goaliyun.String
	Padding       goaliyun.String
	Color         goaliyun.String
	IsKeepCellPic goaliyun.String
}

type SubmitSnapshotJobMNSMessageResult struct {
	MessageId    goaliyun.String
	ErrorMessage goaliyun.String
	ErrorCode    goaliyun.String
}
