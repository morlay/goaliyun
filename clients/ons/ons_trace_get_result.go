package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsTraceGetResultRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	QueryId      string `position:"Query" name:"QueryId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsTraceGetResultRequest) Invoke(client goaliyun.Client) (*OnsTraceGetResultResponse, error) {
	resp := &OnsTraceGetResultResponse{}
	err := client.Request("ons", "OnsTraceGetResult", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsTraceGetResultResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	TraceData OnsTraceGetResultTraceData
}

type OnsTraceGetResultTraceData struct {
	QueryId    goaliyun.String
	UserId     goaliyun.String
	Topic      goaliyun.String
	MsgId      goaliyun.String
	MsgKey     goaliyun.String
	Status     goaliyun.String
	CreateTime goaliyun.Integer
	UpdateTime goaliyun.Integer
	TraceList  OnsTraceGetResultTraceMapDoList
}

type OnsTraceGetResultTraceMapDo struct {
	PubTime      goaliyun.Integer
	Topic        goaliyun.String
	PubGroupName goaliyun.String
	MsgId        goaliyun.String
	Tag          goaliyun.String
	MsgKey       goaliyun.String
	BornHost     goaliyun.String
	CostTime     goaliyun.Integer
	Status       goaliyun.String
	SubList      OnsTraceGetResultSubMapDoList
}

type OnsTraceGetResultSubMapDo struct {
	SubGroupName goaliyun.String
	SuccessCount goaliyun.Integer
	FailCount    goaliyun.Integer
	ClientList   OnsTraceGetResultSubClientInfoDoList
}

type OnsTraceGetResultSubClientInfoDo struct {
	SubGroupName   goaliyun.String
	SubTime        goaliyun.Integer
	ClientHost     goaliyun.String
	ReconsumeTimes goaliyun.Integer
	CostTime       goaliyun.Integer
	Status         goaliyun.String
}

type OnsTraceGetResultTraceMapDoList []OnsTraceGetResultTraceMapDo

func (list *OnsTraceGetResultTraceMapDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTraceGetResultTraceMapDo)
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

type OnsTraceGetResultSubMapDoList []OnsTraceGetResultSubMapDo

func (list *OnsTraceGetResultSubMapDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTraceGetResultSubMapDo)
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

type OnsTraceGetResultSubClientInfoDoList []OnsTraceGetResultSubClientInfoDo

func (list *OnsTraceGetResultSubClientInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTraceGetResultSubClientInfoDo)
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
