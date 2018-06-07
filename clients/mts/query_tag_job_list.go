package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryTagJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	TagJobIds            string `position:"Query" name:"TagJobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryTagJobListRequest) Invoke(client goaliyun.Client) (*QueryTagJobListResponse, error) {
	resp := &QueryTagJobListResponse{}
	err := client.Request("mts", "QueryTagJobList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryTagJobListResponse struct {
	RequestId   goaliyun.String
	TagJobList  QueryTagJobListTagJobList
	NonExistIds QueryTagJobListNonExistIdList
}

type QueryTagJobListTagJob struct {
	Id             goaliyun.String
	UserData       goaliyun.String
	PipelineId     goaliyun.String
	State          goaliyun.String
	Code           goaliyun.String
	Message        goaliyun.String
	CreationTime   goaliyun.String
	Input          QueryTagJobListInput
	VideoTagResult QueryTagJobListVideoTagResult
}

type QueryTagJobListInput struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryTagJobListVideoTagResult struct {
	Details      goaliyun.String
	TagAnResults QueryTagJobListTagAnResultList
	TagFrResults QueryTagJobListTagFrResultList
}

type QueryTagJobListTagAnResult struct {
	Label goaliyun.String
	Score goaliyun.String
}

type QueryTagJobListTagFrResult struct {
	Time     goaliyun.String
	TagFaces QueryTagJobListTagFaceList
}

type QueryTagJobListTagFace struct {
	Name   goaliyun.String
	Score  goaliyun.String
	Target goaliyun.String
}

type QueryTagJobListTagJobList []QueryTagJobListTagJob

func (list *QueryTagJobListTagJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTagJobListTagJob)
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

type QueryTagJobListNonExistIdList []goaliyun.String

func (list *QueryTagJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryTagJobListTagAnResultList []QueryTagJobListTagAnResult

func (list *QueryTagJobListTagAnResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTagJobListTagAnResult)
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

type QueryTagJobListTagFrResultList []QueryTagJobListTagFrResult

func (list *QueryTagJobListTagFrResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTagJobListTagFrResult)
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

type QueryTagJobListTagFaceList []QueryTagJobListTagFace

func (list *QueryTagJobListTagFaceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTagJobListTagFace)
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
