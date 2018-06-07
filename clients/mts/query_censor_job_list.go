package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryCensorJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryCensorJobListRequest) Invoke(client goaliyun.Client) (*QueryCensorJobListResponse, error) {
	resp := &QueryCensorJobListResponse{}
	err := client.Request("mts", "QueryCensorJobList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryCensorJobListResponse struct {
	RequestId     goaliyun.String
	CensorJobList QueryCensorJobListCensorJobList
	NonExistIds   QueryCensorJobListNonExistIdList
}

type QueryCensorJobListCensorJob struct {
	Id                    goaliyun.String
	UserData              goaliyun.String
	PipelineId            goaliyun.String
	State                 goaliyun.String
	Code                  goaliyun.String
	Message               goaliyun.String
	CreationTime          goaliyun.String
	Input                 QueryCensorJobListInput
	CensorConfig          QueryCensorJobListCensorConfig
	CensorPornResult      QueryCensorJobListCensorPornResult
	CensorTerrorismResult QueryCensorJobListCensorTerrorismResult
}

type QueryCensorJobListInput struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryCensorJobListCensorConfig struct {
	Interval   goaliyun.String
	BizType    goaliyun.String
	OutputFile QueryCensorJobListOutputFile
}

type QueryCensorJobListOutputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryCensorJobListCensorPornResult struct {
	Label           goaliyun.String
	Suggestion      goaliyun.String
	MaxScore        goaliyun.String
	AverageScore    goaliyun.String
	PornCounterList QueryCensorJobListCounterList
	PornTopList     QueryCensorJobListTopList
}

type QueryCensorJobListCounter struct {
	Count goaliyun.Integer
	Label goaliyun.String
}

type QueryCensorJobListTop struct {
	Label     goaliyun.String
	Score     goaliyun.String
	Timestamp goaliyun.String
	Index     goaliyun.String
	Object    goaliyun.String
}

type QueryCensorJobListCensorTerrorismResult struct {
	Label                goaliyun.String
	Suggestion           goaliyun.String
	MaxScore             goaliyun.String
	AverageScore         goaliyun.String
	TerrorismCounterList QueryCensorJobListCounter1List
	TerrorismTopList     QueryCensorJobListTop2List
}

type QueryCensorJobListCounter1 struct {
	Count goaliyun.Integer
	Label goaliyun.String
}

type QueryCensorJobListTop2 struct {
	Label     goaliyun.String
	Score     goaliyun.String
	Timestamp goaliyun.String
	Index     goaliyun.String
	Object    goaliyun.String
}

type QueryCensorJobListCensorJobList []QueryCensorJobListCensorJob

func (list *QueryCensorJobListCensorJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCensorJobListCensorJob)
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

type QueryCensorJobListNonExistIdList []goaliyun.String

func (list *QueryCensorJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryCensorJobListCounterList []QueryCensorJobListCounter

func (list *QueryCensorJobListCounterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCensorJobListCounter)
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

type QueryCensorJobListTopList []QueryCensorJobListTop

func (list *QueryCensorJobListTopList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCensorJobListTop)
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

type QueryCensorJobListCounter1List []QueryCensorJobListCounter1

func (list *QueryCensorJobListCounter1List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCensorJobListCounter1)
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

type QueryCensorJobListTop2List []QueryCensorJobListTop2

func (list *QueryCensorJobListTop2List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCensorJobListTop2)
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
