package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListScenariosRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListScenariosRequest) Invoke(client goaliyun.Client) (*ListScenariosResponse, error) {
	resp := &ListScenariosResponse{}
	err := client.Request("ccc", "ListScenarios", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListScenariosResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	Scenarios      ListScenariosScenarioList
}

type ListScenariosScenario struct {
	ScenarioId          goaliyun.String
	ScenarioName        goaliyun.String
	ScenarioDescription goaliyun.String
	Type                goaliyun.String
	IsTemplate          bool
	Surveys             ListScenariosSurveyList
	Variables           ListScenariosKeyValuePairList
	Strategy            ListScenariosStrategy
}

type ListScenariosSurvey struct {
	SurveyId          goaliyun.String
	SurveyName        goaliyun.String
	SurveyDescription goaliyun.String
	Role              goaliyun.String
	Round             goaliyun.Integer
	BeebotId          goaliyun.String
	Intents           ListScenariosIntentNodeList
}

type ListScenariosIntentNode struct {
	NodeId   goaliyun.String
	IntentId goaliyun.String
}

type ListScenariosKeyValuePair struct {
	Key   goaliyun.String
	Value goaliyun.String
}

type ListScenariosStrategy struct {
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
	WorkingTime         ListScenariosTimeFrameList
	RepeatDays          ListScenariosRepeatDayList
}

type ListScenariosTimeFrame struct {
	BeginTime goaliyun.String
	EndTime   goaliyun.String
}

type ListScenariosScenarioList []ListScenariosScenario

func (list *ListScenariosScenarioList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListScenariosScenario)
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

type ListScenariosSurveyList []ListScenariosSurvey

func (list *ListScenariosSurveyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListScenariosSurvey)
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

type ListScenariosKeyValuePairList []ListScenariosKeyValuePair

func (list *ListScenariosKeyValuePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListScenariosKeyValuePair)
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

type ListScenariosIntentNodeList []ListScenariosIntentNode

func (list *ListScenariosIntentNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListScenariosIntentNode)
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

type ListScenariosTimeFrameList []ListScenariosTimeFrame

func (list *ListScenariosTimeFrameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListScenariosTimeFrame)
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

type ListScenariosRepeatDayList []goaliyun.String

func (list *ListScenariosRepeatDayList) UnmarshalJSON(data []byte) error {
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
