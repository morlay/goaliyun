package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListScenarioTemplatesRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListScenarioTemplatesRequest) Invoke(client goaliyun.Client) (*ListScenarioTemplatesResponse, error) {
	resp := &ListScenarioTemplatesResponse{}
	err := client.Request("ccc", "ListScenarioTemplates", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListScenarioTemplatesResponse struct {
	RequestId         goaliyun.String
	Success           bool
	Code              goaliyun.String
	Message           goaliyun.String
	HttpStatusCode    goaliyun.Integer
	ScenarioTemplates ListScenarioTemplatesScenarioList
}

type ListScenarioTemplatesScenario struct {
	ScenarioId          goaliyun.String
	ScenarioName        goaliyun.String
	ScenarioDescription goaliyun.String
	Type                goaliyun.String
	IsTemplate          bool
	Surveys             ListScenarioTemplatesSurveyList
	Variables           ListScenarioTemplatesKeyValuePairList
	Strategy            ListScenarioTemplatesStrategy
}

type ListScenarioTemplatesSurvey struct {
	SurveyId          goaliyun.String
	SurveyName        goaliyun.String
	SurveyDescription goaliyun.String
	Role              goaliyun.String
	Round             goaliyun.Integer
	BeebotId          goaliyun.String
	Intents           ListScenarioTemplatesIntentNodeList
}

type ListScenarioTemplatesIntentNode struct {
	NodeId   goaliyun.String
	IntentId goaliyun.String
}

type ListScenarioTemplatesKeyValuePair struct {
	Key   goaliyun.String
	Value goaliyun.String
}

type ListScenarioTemplatesStrategy struct {
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
	WorkingTime         ListScenarioTemplatesTimeFrameList
	RepeatDays          ListScenarioTemplatesRepeatDayList
}

type ListScenarioTemplatesTimeFrame struct {
	BeginTime goaliyun.String
	EndTime   goaliyun.String
}

type ListScenarioTemplatesScenarioList []ListScenarioTemplatesScenario

func (list *ListScenarioTemplatesScenarioList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListScenarioTemplatesScenario)
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

type ListScenarioTemplatesSurveyList []ListScenarioTemplatesSurvey

func (list *ListScenarioTemplatesSurveyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListScenarioTemplatesSurvey)
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

type ListScenarioTemplatesKeyValuePairList []ListScenarioTemplatesKeyValuePair

func (list *ListScenarioTemplatesKeyValuePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListScenarioTemplatesKeyValuePair)
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

type ListScenarioTemplatesIntentNodeList []ListScenarioTemplatesIntentNode

func (list *ListScenarioTemplatesIntentNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListScenarioTemplatesIntentNode)
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

type ListScenarioTemplatesTimeFrameList []ListScenarioTemplatesTimeFrame

func (list *ListScenarioTemplatesTimeFrameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListScenarioTemplatesTimeFrame)
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

type ListScenarioTemplatesRepeatDayList []goaliyun.String

func (list *ListScenarioTemplatesRepeatDayList) UnmarshalJSON(data []byte) error {
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
