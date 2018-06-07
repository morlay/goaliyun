package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryCoverJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CoverJobIds          string `position:"Query" name:"CoverJobIds"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryCoverJobListRequest) Invoke(client goaliyun.Client) (*QueryCoverJobListResponse, error) {
	resp := &QueryCoverJobListResponse{}
	err := client.Request("mts", "QueryCoverJobList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryCoverJobListResponse struct {
	RequestId    goaliyun.String
	CoverJobList QueryCoverJobListCoverJobList
	NonExistIds  QueryCoverJobListNonExistIdList
}

type QueryCoverJobListCoverJob struct {
	Id             goaliyun.String
	UserData       goaliyun.String
	PipelineId     goaliyun.String
	State          goaliyun.String
	Code           goaliyun.String
	Message        goaliyun.String
	CreationTime   goaliyun.String
	CoverImageList QueryCoverJobListCoverImageList
	Input          QueryCoverJobListInput
	CoverConfig    QueryCoverJobListCoverConfig
}

type QueryCoverJobListCoverImage struct {
	Score goaliyun.String
	Url   goaliyun.String
	Time  goaliyun.String
}

type QueryCoverJobListInput struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryCoverJobListCoverConfig struct {
	OutputFile QueryCoverJobListOutputFile
}

type QueryCoverJobListOutputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryCoverJobListCoverJobList []QueryCoverJobListCoverJob

func (list *QueryCoverJobListCoverJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCoverJobListCoverJob)
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

type QueryCoverJobListNonExistIdList []goaliyun.String

func (list *QueryCoverJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryCoverJobListCoverImageList []QueryCoverJobListCoverImage

func (list *QueryCoverJobListCoverImageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCoverJobListCoverImage)
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
