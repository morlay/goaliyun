package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryFpShotJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryFpShotJobListRequest) Invoke(client goaliyun.Client) (*QueryFpShotJobListResponse, error) {
	resp := &QueryFpShotJobListResponse{}
	err := client.Request("mts", "QueryFpShotJobList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryFpShotJobListResponse struct {
	RequestId     goaliyun.String
	FpShotJobList QueryFpShotJobListFpShotJobList
	NonExistIds   QueryFpShotJobListNonExistIdList
}

type QueryFpShotJobListFpShotJob struct {
	Id           goaliyun.String
	UserData     goaliyun.String
	PipelineId   goaliyun.String
	State        goaliyun.String
	Code         goaliyun.String
	Message      goaliyun.String
	CreationTime goaliyun.String
	InputFile    QueryFpShotJobListInputFile
	FpShotResult QueryFpShotJobListFpShotResult
}

type QueryFpShotJobListInputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryFpShotJobListFpShotResult struct {
	FpShots QueryFpShotJobListFpShotList
}

type QueryFpShotJobListFpShot struct {
	PrimaryKey   goaliyun.String
	Similarity   goaliyun.String
	FpShotSlices QueryFpShotJobListFpShotSliceList
}

type QueryFpShotJobListFpShotSlice struct {
	Input       QueryFpShotJobListInput
	Duplication QueryFpShotJobListDuplication
}

type QueryFpShotJobListInput struct {
	Start    goaliyun.String
	Duration goaliyun.String
}

type QueryFpShotJobListDuplication struct {
	Start    goaliyun.String
	Duration goaliyun.String
}

type QueryFpShotJobListFpShotJobList []QueryFpShotJobListFpShotJob

func (list *QueryFpShotJobListFpShotJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFpShotJobListFpShotJob)
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

type QueryFpShotJobListNonExistIdList []goaliyun.String

func (list *QueryFpShotJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryFpShotJobListFpShotList []QueryFpShotJobListFpShot

func (list *QueryFpShotJobListFpShotList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFpShotJobListFpShot)
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

type QueryFpShotJobListFpShotSliceList []QueryFpShotJobListFpShotSlice

func (list *QueryFpShotJobListFpShotSliceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFpShotJobListFpShotSlice)
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
