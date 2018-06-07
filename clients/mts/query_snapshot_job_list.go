package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QuerySnapshotJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	SnapshotJobIds       string `position:"Query" name:"SnapshotJobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QuerySnapshotJobListRequest) Invoke(client goaliyun.Client) (*QuerySnapshotJobListResponse, error) {
	resp := &QuerySnapshotJobListResponse{}
	err := client.Request("mts", "QuerySnapshotJobList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QuerySnapshotJobListResponse struct {
	RequestId              goaliyun.String
	SnapshotJobList        QuerySnapshotJobListSnapshotJobList
	NonExistSnapshotJobIds QuerySnapshotJobListNonExistSnapshotJobIdList
}

type QuerySnapshotJobListSnapshotJob struct {
	Id               goaliyun.String
	UserData         goaliyun.String
	PipelineId       goaliyun.String
	State            goaliyun.String
	Code             goaliyun.String
	Count            goaliyun.String
	TileCount        goaliyun.String
	Message          goaliyun.String
	CreationTime     goaliyun.String
	Input            QuerySnapshotJobListInput
	SnapshotConfig   QuerySnapshotJobListSnapshotConfig
	MNSMessageResult QuerySnapshotJobListMNSMessageResult
}

type QuerySnapshotJobListInput struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
	RoleArn  goaliyun.String
}

type QuerySnapshotJobListSnapshotConfig struct {
	Time           goaliyun.String
	Interval       goaliyun.String
	Num            goaliyun.String
	Width          goaliyun.String
	Height         goaliyun.String
	FrameType      goaliyun.String
	OutputFile     QuerySnapshotJobListOutputFile
	TileOutputFile QuerySnapshotJobListTileOutputFile
	TileOut        QuerySnapshotJobListTileOut
}

type QuerySnapshotJobListOutputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
	RoleArn  goaliyun.String
}

type QuerySnapshotJobListTileOutputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
	RoleArn  goaliyun.String
}

type QuerySnapshotJobListTileOut struct {
	Lines         goaliyun.String
	Columns       goaliyun.String
	CellWidth     goaliyun.String
	CellHeight    goaliyun.String
	Margin        goaliyun.String
	Padding       goaliyun.String
	Color         goaliyun.String
	IsKeepCellPic goaliyun.String
}

type QuerySnapshotJobListMNSMessageResult struct {
	MessageId    goaliyun.String
	ErrorMessage goaliyun.String
	ErrorCode    goaliyun.String
}

type QuerySnapshotJobListSnapshotJobList []QuerySnapshotJobListSnapshotJob

func (list *QuerySnapshotJobListSnapshotJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QuerySnapshotJobListSnapshotJob)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type QuerySnapshotJobListNonExistSnapshotJobIdList []goaliyun.String

func (list *QuerySnapshotJobListNonExistSnapshotJobIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
