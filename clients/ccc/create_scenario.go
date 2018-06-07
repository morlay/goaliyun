package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateScenarioRequest struct {
	InstanceId   string                         `position:"Query" name:"InstanceId"`
	SurveysJsons *CreateScenarioSurveysJsonList `position:"Query" type:"Repeated" name:"SurveysJson"`
	StrategyJson string                         `position:"Query" name:"StrategyJson"`
	Name         string                         `position:"Query" name:"Name"`
	Description  string                         `position:"Query" name:"Description"`
	RegionId     string                         `position:"Query" name:"RegionId"`
}

func (req *CreateScenarioRequest) Invoke(client goaliyun.Client) (*CreateScenarioResponse, error) {
	resp := &CreateScenarioResponse{}
	err := client.Request("ccc", "CreateScenario", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateScenarioResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	Scenario       CreateScenarioScenario
}

type CreateScenarioScenario struct {
	ScenarioId          goaliyun.String
	ScenarioName        goaliyun.String
	ScenarioDescription goaliyun.String
	Type                goaliyun.String
	IsTemplate          bool
	Surveys             CreateScenarioSurveyList
	Variables           CreateScenarioKeyValuePairList
	Strategy            CreateScenarioStrategy
}

type CreateScenarioSurvey struct {
	SurveyId          goaliyun.String
	SurveyName        goaliyun.String
	SurveyDescription goaliyun.String
	Role              goaliyun.String
	Round             goaliyun.Integer
	BeebotId          goaliyun.String
	Intents           CreateScenarioIntentNodeList
}

type CreateScenarioIntentNode struct {
	NodeId   goaliyun.String
	IntentId goaliyun.String
}

type CreateScenarioKeyValuePair struct {
	Key   goaliyun.String
	Value goaliyun.String
}

type CreateScenarioStrategy struct {
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
	WorkingTime         CreateScenarioTimeFrameList
	RepeatDays          CreateScenarioRepeatDayList
}

type CreateScenarioTimeFrame struct {
	BeginTime goaliyun.String
	EndTime   goaliyun.String
}

type CreateScenarioSurveysJsonList []string

func (list *CreateScenarioSurveysJsonList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type CreateScenarioSurveyList []CreateScenarioSurvey

func (list *CreateScenarioSurveyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateScenarioSurvey)
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

type CreateScenarioKeyValuePairList []CreateScenarioKeyValuePair

func (list *CreateScenarioKeyValuePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateScenarioKeyValuePair)
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

type CreateScenarioIntentNodeList []CreateScenarioIntentNode

func (list *CreateScenarioIntentNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateScenarioIntentNode)
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

type CreateScenarioTimeFrameList []CreateScenarioTimeFrame

func (list *CreateScenarioTimeFrameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateScenarioTimeFrame)
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

type CreateScenarioRepeatDayList []goaliyun.String

func (list *CreateScenarioRepeatDayList) UnmarshalJSON(data []byte) error {
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
