package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateScenarioFromTemplateRequest struct {
	Variables   string `position:"Query" name:"Variables"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	Name        string `position:"Query" name:"Name"`
	Description string `position:"Query" name:"Description"`
	TemplateId  string `position:"Query" name:"TemplateId"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *CreateScenarioFromTemplateRequest) Invoke(client goaliyun.Client) (*CreateScenarioFromTemplateResponse, error) {
	resp := &CreateScenarioFromTemplateResponse{}
	err := client.Request("ccc", "CreateScenarioFromTemplate", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateScenarioFromTemplateResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	Scenario       CreateScenarioFromTemplateScenario
}

type CreateScenarioFromTemplateScenario struct {
	ScenarioId          goaliyun.String
	ScenarioName        goaliyun.String
	ScenarioDescription goaliyun.String
	Type                goaliyun.String
	IsTemplate          bool
	Surveys             CreateScenarioFromTemplateSurveyList
	Variables           CreateScenarioFromTemplateKeyValuePairList
	Strategy            CreateScenarioFromTemplateStrategy
}

type CreateScenarioFromTemplateSurvey struct {
	SurveyId          goaliyun.String
	SurveyName        goaliyun.String
	SurveyDescription goaliyun.String
	Role              goaliyun.String
	Round             goaliyun.Integer
	BeebotId          goaliyun.String
	Intents           CreateScenarioFromTemplateIntentNodeList
}

type CreateScenarioFromTemplateIntentNode struct {
	NodeId   goaliyun.String
	IntentId goaliyun.String
}

type CreateScenarioFromTemplateKeyValuePair struct {
	Key   goaliyun.String
	Value goaliyun.String
}

type CreateScenarioFromTemplateStrategy struct {
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
	WorkingTime         CreateScenarioFromTemplateTimeFrameList
	RepeatDays          CreateScenarioFromTemplateRepeatDayList
}

type CreateScenarioFromTemplateTimeFrame struct {
	BeginTime goaliyun.String
	EndTime   goaliyun.String
}

type CreateScenarioFromTemplateSurveyList []CreateScenarioFromTemplateSurvey

func (list *CreateScenarioFromTemplateSurveyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateScenarioFromTemplateSurvey)
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

type CreateScenarioFromTemplateKeyValuePairList []CreateScenarioFromTemplateKeyValuePair

func (list *CreateScenarioFromTemplateKeyValuePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateScenarioFromTemplateKeyValuePair)
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

type CreateScenarioFromTemplateIntentNodeList []CreateScenarioFromTemplateIntentNode

func (list *CreateScenarioFromTemplateIntentNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateScenarioFromTemplateIntentNode)
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

type CreateScenarioFromTemplateTimeFrameList []CreateScenarioFromTemplateTimeFrame

func (list *CreateScenarioFromTemplateTimeFrameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateScenarioFromTemplateTimeFrame)
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

type CreateScenarioFromTemplateRepeatDayList []goaliyun.String

func (list *CreateScenarioFromTemplateRepeatDayList) UnmarshalJSON(data []byte) error {
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
