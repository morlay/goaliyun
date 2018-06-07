package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryPornJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryPornJobListRequest) Invoke(client goaliyun.Client) (*QueryPornJobListResponse, error) {
	resp := &QueryPornJobListResponse{}
	err := client.Request("mts", "QueryPornJobList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryPornJobListResponse struct {
	RequestId   goaliyun.String
	PornJobList QueryPornJobListPornJobList
	NonExistIds QueryPornJobListNonExistIdList
}

type QueryPornJobListPornJob struct {
	Id               goaliyun.String
	UserData         goaliyun.String
	PipelineId       goaliyun.String
	State            goaliyun.String
	Code             goaliyun.String
	Message          goaliyun.String
	CreationTime     goaliyun.String
	Input            QueryPornJobListInput
	PornConfig       QueryPornJobListPornConfig
	CensorPornResult QueryPornJobListCensorPornResult
}

type QueryPornJobListInput struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryPornJobListPornConfig struct {
	Interval   goaliyun.String
	BizType    goaliyun.String
	OutputFile QueryPornJobListOutputFile
}

type QueryPornJobListOutputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryPornJobListCensorPornResult struct {
	Label           goaliyun.String
	Suggestion      goaliyun.String
	MaxScore        goaliyun.String
	AverageScore    goaliyun.String
	PornCounterList QueryPornJobListCounterList
	PornTopList     QueryPornJobListTopList
}

type QueryPornJobListCounter struct {
	Count goaliyun.Integer
	Label goaliyun.String
}

type QueryPornJobListTop struct {
	Label     goaliyun.String
	Score     goaliyun.String
	Timestamp goaliyun.String
	Index     goaliyun.String
	Object    goaliyun.String
}

type QueryPornJobListPornJobList []QueryPornJobListPornJob

func (list *QueryPornJobListPornJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPornJobListPornJob)
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

type QueryPornJobListNonExistIdList []goaliyun.String

func (list *QueryPornJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryPornJobListCounterList []QueryPornJobListCounter

func (list *QueryPornJobListCounterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPornJobListCounter)
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

type QueryPornJobListTopList []QueryPornJobListTop

func (list *QueryPornJobListTopList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPornJobListTop)
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
