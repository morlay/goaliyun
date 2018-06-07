package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GenerateAgentStatisticReportRequest struct {
	AgentId    string `position:"Query" name:"AgentId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	EndDate    string `position:"Query" name:"EndDate"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	StartDate  string `position:"Query" name:"StartDate"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *GenerateAgentStatisticReportRequest) Invoke(client goaliyun.Client) (*GenerateAgentStatisticReportResponse, error) {
	resp := &GenerateAgentStatisticReportResponse{}
	err := client.Request("ccc", "GenerateAgentStatisticReport", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GenerateAgentStatisticReportResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	DataList       GenerateAgentStatisticReportDataList
}

type GenerateAgentStatisticReportDataList struct {
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	List       GenerateAgentStatisticReportGenerateAgentStatisticList
}

type GenerateAgentStatisticReportGenerateAgentStatistic struct {
	AgentId           goaliyun.String
	LoginName         goaliyun.String
	AgentName         goaliyun.String
	SkillGroupIds     goaliyun.String
	SkillGroupNames   goaliyun.String
	InstanceId        goaliyun.String
	RecordDate        goaliyun.String
	TotalLoggedInTime goaliyun.Integer
	TotalBreakTime    goaliyun.Integer
	OccupancyRate     goaliyun.Float
	TotalReadyTime    goaliyun.Integer
	MaxReadyTime      goaliyun.Integer
	AverageReadyTime  goaliyun.Integer
	Inbound           GenerateAgentStatisticReportInbound
	Outbound          GenerateAgentStatisticReportOutbound
	Overall           GenerateAgentStatisticReportOverall
}

type GenerateAgentStatisticReportInbound struct {
	TotalTalkTime                goaliyun.Integer
	MaxTalkTime                  goaliyun.Integer
	AverageTalkTime              goaliyun.Integer
	TotalHoldTime                goaliyun.Integer
	MaxHoldTime                  goaliyun.Integer
	AverageHoldTime              goaliyun.Integer
	TotalWorkTime                goaliyun.Integer
	MaxWorkTime                  goaliyun.Integer
	AverageWorkTime              goaliyun.Integer
	SatisfactionSurveysOffered   goaliyun.Integer
	SatisfactionSurveysResponded goaliyun.Integer
	SatisfactionIndex            goaliyun.Float
	CallsOffered                 goaliyun.Integer
	CallsHandled                 goaliyun.Integer
	HandleRate                   goaliyun.Float
	TotalRingTime                goaliyun.Integer
	MaxRingTime                  goaliyun.Integer
	AverageRingTime              goaliyun.Integer
}

type GenerateAgentStatisticReportOutbound struct {
	TotalTalkTime                goaliyun.Integer
	MaxTalkTime                  goaliyun.Integer
	AverageTalkTime              goaliyun.Integer
	TotalHoldTime                goaliyun.Integer
	MaxHoldTime                  goaliyun.Integer
	AverageHoldTime              goaliyun.Integer
	TotalWorkTime                goaliyun.Integer
	MaxWorkTime                  goaliyun.Integer
	AverageWorkTime              goaliyun.Integer
	SatisfactionSurveysOffered   goaliyun.Integer
	SatisfactionSurveysResponded goaliyun.Integer
	SatisfactionIndex            goaliyun.Float
	CallsDialed                  goaliyun.Integer
	CallsAnswered                goaliyun.Integer
	AnswerRate                   goaliyun.Float
	TotalDialingTime             goaliyun.Integer
	TotalDialingTime1            goaliyun.Integer
	MaxDialingTime               goaliyun.Integer
	AverageDialingTime           goaliyun.Integer
}

type GenerateAgentStatisticReportOverall struct {
	TotalTalkTime                goaliyun.Integer
	MaxTalkTime                  goaliyun.Integer
	AverageTalkTime              goaliyun.Integer
	TotalHoldTime                goaliyun.Integer
	MaxHoldTime                  goaliyun.Integer
	AverageHoldTime              goaliyun.Integer
	TotalWorkTime                goaliyun.Integer
	MaxWorkTime                  goaliyun.Integer
	AverageWorkTime              goaliyun.Integer
	SatisfactionSurveysOffered   goaliyun.Integer
	SatisfactionSurveysResponded goaliyun.Integer
	SatisfactionIndex            goaliyun.Float
	TotalCalls                   goaliyun.Integer
}

type GenerateAgentStatisticReportGenerateAgentStatisticList []GenerateAgentStatisticReportGenerateAgentStatistic

func (list *GenerateAgentStatisticReportGenerateAgentStatisticList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GenerateAgentStatisticReportGenerateAgentStatistic)
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
