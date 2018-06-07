package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryTerrorismJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryTerrorismJobListRequest) Invoke(client goaliyun.Client) (*QueryTerrorismJobListResponse, error) {
	resp := &QueryTerrorismJobListResponse{}
	err := client.Request("mts", "QueryTerrorismJobList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryTerrorismJobListResponse struct {
	RequestId        goaliyun.String
	TerrorismJobList QueryTerrorismJobListTerrorismJobList
	NonExistIds      QueryTerrorismJobListNonExistIdList
}

type QueryTerrorismJobListTerrorismJob struct {
	Id                    goaliyun.String
	UserData              goaliyun.String
	PipelineId            goaliyun.String
	State                 goaliyun.String
	Code                  goaliyun.String
	Message               goaliyun.String
	CreationTime          goaliyun.String
	Input                 QueryTerrorismJobListInput
	TerrorismConfig       QueryTerrorismJobListTerrorismConfig
	CensorTerrorismResult QueryTerrorismJobListCensorTerrorismResult
}

type QueryTerrorismJobListInput struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryTerrorismJobListTerrorismConfig struct {
	Interval   goaliyun.String
	BizType    goaliyun.String
	OutputFile QueryTerrorismJobListOutputFile
}

type QueryTerrorismJobListOutputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryTerrorismJobListCensorTerrorismResult struct {
	Label                goaliyun.String
	Suggestion           goaliyun.String
	MaxScore             goaliyun.String
	AverageScore         goaliyun.String
	TerrorismCounterList QueryTerrorismJobListCounterList
	TerrorismTopList     QueryTerrorismJobListTopList
}

type QueryTerrorismJobListCounter struct {
	Count goaliyun.Integer
	Label goaliyun.String
}

type QueryTerrorismJobListTop struct {
	Label     goaliyun.String
	Score     goaliyun.String
	Timestamp goaliyun.String
	Index     goaliyun.String
	Object    goaliyun.String
}

type QueryTerrorismJobListTerrorismJobList []QueryTerrorismJobListTerrorismJob

func (list *QueryTerrorismJobListTerrorismJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTerrorismJobListTerrorismJob)
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

type QueryTerrorismJobListNonExistIdList []goaliyun.String

func (list *QueryTerrorismJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryTerrorismJobListCounterList []QueryTerrorismJobListCounter

func (list *QueryTerrorismJobListCounterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTerrorismJobListCounter)
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

type QueryTerrorismJobListTopList []QueryTerrorismJobListTop

func (list *QueryTerrorismJobListTopList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTerrorismJobListTop)
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
