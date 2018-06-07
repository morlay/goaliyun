package qualitycheck

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetAccAsrResultRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetAccAsrResultRequest) Invoke(client goaliyun.Client) (*GetAccAsrResultResponse, error) {
	resp := &GetAccAsrResultResponse{}
	err := client.Request("qualitycheck", "GetAccAsrResult", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetAccAsrResultResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
	Count     goaliyun.Integer
	Data      GetAccAsrResultAccAsrSentenceResultList
}

type GetAccAsrResultAccAsrSentenceResult struct {
	RecordId         goaliyun.String
	Status           goaliyun.String
	StatusCode       goaliyun.String
	ErrorMessage     goaliyun.String
	Duration         goaliyun.Integer
	InteractiveCount goaliyun.Integer
	Results          GetAccAsrResultSentenceResultList
	ServiceEvStat    GetAccAsrResultServiceEvStat
	ClientEvStat     GetAccAsrResultClientEvStat
	ServiceSrStat    GetAccAsrResultServiceSrStat
	ClientSrStat     GetAccAsrResultClientSrStat
}

type GetAccAsrResultSentenceResult struct {
	BeginTime       goaliyun.Integer
	EndTime         goaliyun.Integer
	ChannelId       goaliyun.Integer
	Text            goaliyun.String
	EmotionValue    goaliyun.Integer
	SilenceDuration goaliyun.Integer
	SpeechRate      goaliyun.Integer
	SpeakerId       goaliyun.String
	AgentId         goaliyun.String
	ChannelKey      goaliyun.String
}

type GetAccAsrResultServiceEvStat struct {
	Srole            goaliyun.Integer
	SmaxEmotionValue goaliyun.Float
	SminEmotionValue goaliyun.Float
	SavgEmotionValue goaliyun.Float
}

type GetAccAsrResultClientEvStat struct {
	Crole            goaliyun.Integer
	CmaxEmotionValue goaliyun.Float
	CminEmotionValue goaliyun.Float
	CavgEmotionValue goaliyun.Float
}

type GetAccAsrResultServiceSrStat struct {
	Srole          goaliyun.Integer
	SmaxSpeechRate goaliyun.Float
	SminSpeechRate goaliyun.Float
	SavgSpeechRate goaliyun.Float
}

type GetAccAsrResultClientSrStat struct {
	Crole          goaliyun.Integer
	CmaxSpeechRate goaliyun.Float
	CminSpeechRate goaliyun.Float
	CavgSpeechRate goaliyun.Float
}

type GetAccAsrResultAccAsrSentenceResultList []GetAccAsrResultAccAsrSentenceResult

func (list *GetAccAsrResultAccAsrSentenceResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetAccAsrResultAccAsrSentenceResult)
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

type GetAccAsrResultSentenceResultList []GetAccAsrResultSentenceResult

func (list *GetAccAsrResultSentenceResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetAccAsrResultSentenceResult)
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
