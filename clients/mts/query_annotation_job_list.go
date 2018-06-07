package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryAnnotationJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AnnotationJobIds     string `position:"Query" name:"AnnotationJobIds"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryAnnotationJobListRequest) Invoke(client goaliyun.Client) (*QueryAnnotationJobListResponse, error) {
	resp := &QueryAnnotationJobListResponse{}
	err := client.Request("mts", "QueryAnnotationJobList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryAnnotationJobListResponse struct {
	RequestId         goaliyun.String
	AnnotationJobList QueryAnnotationJobListAnnotationJobList
	NonExistIds       QueryAnnotationJobListNonExistIdList
}

type QueryAnnotationJobListAnnotationJob struct {
	Id                    goaliyun.String
	UserData              goaliyun.String
	PipelineId            goaliyun.String
	State                 goaliyun.String
	Code                  goaliyun.String
	Message               goaliyun.String
	CreationTime          goaliyun.String
	Input                 QueryAnnotationJobListInput
	VideoAnnotationResult QueryAnnotationJobListVideoAnnotationResult
}

type QueryAnnotationJobListInput struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryAnnotationJobListVideoAnnotationResult struct {
	Details     goaliyun.String
	Annotations QueryAnnotationJobListAnnotationList
}

type QueryAnnotationJobListAnnotation struct {
	Label goaliyun.String
	Score goaliyun.String
}

type QueryAnnotationJobListAnnotationJobList []QueryAnnotationJobListAnnotationJob

func (list *QueryAnnotationJobListAnnotationJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAnnotationJobListAnnotationJob)
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

type QueryAnnotationJobListNonExistIdList []goaliyun.String

func (list *QueryAnnotationJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryAnnotationJobListAnnotationList []QueryAnnotationJobListAnnotation

func (list *QueryAnnotationJobListAnnotationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAnnotationJobListAnnotation)
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
