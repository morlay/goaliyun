package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetAgentDataRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	StartDay   string `position:"Query" name:"StartDay"`
	EndDay     string `position:"Query" name:"EndDay"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	UserId     string `position:"Query" name:"UserId"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *GetAgentDataRequest) Invoke(client goaliyun.Client) (*GetAgentDataResponse, error) {
	resp := &GetAgentDataResponse{}
	err := client.Request("ccc", "GetAgentData", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetAgentDataResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	DataList       GetAgentDataDataList
}

type GetAgentDataDataList struct {
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	List       GetAgentDataAgentReportDataList
}

type GetAgentDataAgentReportData struct {
	UserId                             goaliyun.String
	LoginName                          goaliyun.String
	DisplayName                        goaliyun.String
	SkillGroupIds                      goaliyun.String
	SkillGroupNames                    goaliyun.String
	InstanceId                         goaliyun.String
	RecordDate                         goaliyun.String
	AgentStatus                        goaliyun.String
	TotalCalls                         goaliyun.String
	CallsAnswered                      goaliyun.String
	CallsAnsweredRate                  goaliyun.String
	CallsAbandoned                     goaliyun.String
	CallsAbandonedRate                 goaliyun.String
	AverageTalkTime                    goaliyun.String
	AverageSpeedOfAnswer               goaliyun.String
	LoggedInTime                       goaliyun.String
	TalkTime                           goaliyun.String
	TalkTimeRate                       goaliyun.String
	BreakTime                          goaliyun.String
	BreakTimeRate                      goaliyun.String
	LoginTime                          goaliyun.String
	LogoutTime                         goaliyun.String
	WorkTime                           goaliyun.String
	InboundCalls                       goaliyun.String
	InboundCallsAnswered               goaliyun.String
	InboundCallsAnsweredRate           goaliyun.String
	InboundCallsAbandoned              goaliyun.String
	InboundCallsAbandonedRate          goaliyun.String
	InboundRingTime                    goaliyun.String
	InboundTalkTime                    goaliyun.String
	OutboundCalls                      goaliyun.String
	OutboundCallsAnswered              goaliyun.String
	OutboundCallsAnsweredRate          goaliyun.String
	OutboundDialingTime                goaliyun.String
	OutboundTalkTime                   goaliyun.String
	InboundSatisfactionSurvey          goaliyun.String
	InboundSatisfactionSurveyRate      goaliyun.String
	InboundTakeSatisfactionSurvey      goaliyun.String
	InboundTakeSatisfactionSurveyRate  goaliyun.String
	InboundSatisfactionRate            goaliyun.String
	InboundFeedbackSumNum              goaliyun.String
	OutboundSatisfactionSurvey         goaliyun.String
	OutboundSatisfactionSurveyRate     goaliyun.String
	OutboundTakeSatisfactionSurvey     goaliyun.String
	OutboundTakeSatisfactionSurveyRate goaliyun.String
	OutboundSatisfactionRate           goaliyun.String
	OutboundFeedbackSumNum             goaliyun.String
	InboundSatisfactionDetail          GetAgentDataAppraiseCountDomainList
	InboundFeedbackNum                 GetAgentDataAppraiseCountDomainList
	OutboundAppraiseNum                GetAgentDataAppraiseCountDomainList
	OutboundFeedbackNum                GetAgentDataAppraiseCountDomainList
}

type GetAgentDataAppraiseCountDomain struct {
	Group    goaliyun.String
	Subgroup goaliyun.String
	Count    goaliyun.String
}

type GetAgentDataAgentReportDataList []GetAgentDataAgentReportData

func (list *GetAgentDataAgentReportDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetAgentDataAgentReportData)
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

type GetAgentDataAppraiseCountDomainList []GetAgentDataAppraiseCountDomain

func (list *GetAgentDataAppraiseCountDomainList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetAgentDataAppraiseCountDomain)
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
