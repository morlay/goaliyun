package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryFacerecogJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	FacerecogJobIds      string `position:"Query" name:"FacerecogJobIds"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryFacerecogJobListRequest) Invoke(client goaliyun.Client) (*QueryFacerecogJobListResponse, error) {
	resp := &QueryFacerecogJobListResponse{}
	err := client.Request("mts", "QueryFacerecogJobList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryFacerecogJobListResponse struct {
	RequestId        goaliyun.String
	FacerecogJobList QueryFacerecogJobListFacerecogJobList
	NonExistIds      QueryFacerecogJobListNonExistIdList
}

type QueryFacerecogJobListFacerecogJob struct {
	Id                   goaliyun.String
	UserData             goaliyun.String
	PipelineId           goaliyun.String
	State                goaliyun.String
	Code                 goaliyun.String
	Message              goaliyun.String
	CreationTime         goaliyun.String
	Input                QueryFacerecogJobListInput
	VideoFacerecogResult QueryFacerecogJobListVideoFacerecogResult
}

type QueryFacerecogJobListInput struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryFacerecogJobListVideoFacerecogResult struct {
	Facerecogs QueryFacerecogJobListFacerecogList
}

type QueryFacerecogJobListFacerecog struct {
	Time  goaliyun.String
	Faces QueryFacerecogJobListFaceList
}

type QueryFacerecogJobListFace struct {
	Name   goaliyun.String
	Score  goaliyun.String
	Target goaliyun.String
}

type QueryFacerecogJobListFacerecogJobList []QueryFacerecogJobListFacerecogJob

func (list *QueryFacerecogJobListFacerecogJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFacerecogJobListFacerecogJob)
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

type QueryFacerecogJobListNonExistIdList []goaliyun.String

func (list *QueryFacerecogJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryFacerecogJobListFacerecogList []QueryFacerecogJobListFacerecog

func (list *QueryFacerecogJobListFacerecogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFacerecogJobListFacerecog)
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

type QueryFacerecogJobListFaceList []QueryFacerecogJobListFace

func (list *QueryFacerecogJobListFaceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFacerecogJobListFace)
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
