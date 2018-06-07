package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateBatchJobsRequest struct {
	CallingNumbers *CreateBatchJobsCallingNumberList `position:"Query" type:"Repeated" name:"CallingNumber"`
	InstanceId     string                            `position:"Query" name:"InstanceId"`
	Submitted      string                            `position:"Query" name:"Submitted"`
	StrategyJson   string                            `position:"Query" name:"StrategyJson"`
	Name           string                            `position:"Query" name:"Name"`
	Description    string                            `position:"Query" name:"Description"`
	ScenarioId     string                            `position:"Query" name:"ScenarioId"`
	JobFilePath    string                            `position:"Query" name:"JobFilePath"`
	RegionId       string                            `position:"Query" name:"RegionId"`
}

func (req *CreateBatchJobsRequest) Invoke(client goaliyun.Client) (*CreateBatchJobsResponse, error) {
	resp := &CreateBatchJobsResponse{}
	err := client.Request("ccc", "CreateBatchJobs", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateBatchJobsResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	JobGroup       CreateBatchJobsJobGroup
}

type CreateBatchJobsJobGroup struct {
	JobGroupId          goaliyun.String
	JobGroupName        goaliyun.String
	JobGroupDescription goaliyun.String
	ScenarioId          goaliyun.String
	JobFilePath         goaliyun.String
	CreationTime        goaliyun.Integer
	CallingNumbers      CreateBatchJobsCallingNumberList
	Strategy            CreateBatchJobsStrategy
}

type CreateBatchJobsStrategy struct {
	StrategyId          goaliyun.String
	StrategyName        goaliyun.String
	StrategyDescription goaliyun.String
	Type                goaliyun.String
	StartTime           goaliyun.Integer
	EndTime             goaliyun.Integer
	RepeatBy            goaliyun.String
	MaxAttemptsPerDay   goaliyun.Integer
	MinAttemptInterval  goaliyun.Integer
	Customized          goaliyun.String
	RoutingStrategy     goaliyun.String
	FollowUpStrategy    goaliyun.String
	IsTemplate          bool
	WorkingTime         CreateBatchJobsTimeFrameList
	RepeatDays          CreateBatchJobsRepeatDayList
}

type CreateBatchJobsTimeFrame struct {
	BeginTime goaliyun.String
	EndTime   goaliyun.String
}

type CreateBatchJobsCallingNumberList []goaliyun.String

func (list *CreateBatchJobsCallingNumberList) UnmarshalJSON(data []byte) error {
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

type CreateBatchJobsTimeFrameList []CreateBatchJobsTimeFrame

func (list *CreateBatchJobsTimeFrameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateBatchJobsTimeFrame)
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

type CreateBatchJobsRepeatDayList []goaliyun.String

func (list *CreateBatchJobsRepeatDayList) UnmarshalJSON(data []byte) error {
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
