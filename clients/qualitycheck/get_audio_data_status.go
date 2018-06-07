package qualitycheck

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetAudioDataStatusRequest struct {
	JsonStr  string `position:"Query" name:"JsonStr"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetAudioDataStatusRequest) Invoke(client goaliyun.Client) (*GetAudioDataStatusResponse, error) {
	resp := &GetAudioDataStatusResponse{}
	err := client.Request("qualitycheck", "GetAudioDataStatus", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetAudioDataStatusResponse struct {
	RequestId     goaliyun.String
	Success       bool
	Code          goaliyun.String
	Message       goaliyun.String
	Count         goaliyun.Integer
	OverallStatus goaliyun.Integer
	Data          GetAudioDataStatusTaskAsrResultList
}

type GetAudioDataStatusTaskAsrResult struct {
	Tid        goaliyun.String
	StatusCode goaliyun.Integer
	StatusMsg  goaliyun.String
	AsrResult  GetAudioDataStatusAsrResult
}

type GetAudioDataStatusAsrResult struct {
	Asrstatus        goaliyun.String
	AsrStatusCode    goaliyun.String
	ErrorMessage     goaliyun.String
	Duration         goaliyun.Integer
	InteractiveCount goaliyun.Integer
	SentenceResults  GetAudioDataStatusSentenceResultList
	ServiceEvStat    GetAudioDataStatusServiceEvStat
	ClientEvStat     GetAudioDataStatusClientEvStat
	ServiceSrStat    GetAudioDataStatusServiceSrStat
	ClientSrStat     GetAudioDataStatusClientSrStat
}

type GetAudioDataStatusSentenceResult struct {
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

type GetAudioDataStatusServiceEvStat struct {
	Srole            goaliyun.Integer
	SmaxEmotionValue goaliyun.Float
	SminEmotionValue goaliyun.Float
	SavgEmotionValue goaliyun.Float
}

type GetAudioDataStatusClientEvStat struct {
	Crole            goaliyun.Integer
	CmaxEmotionValue goaliyun.Float
	CminEmotionValue goaliyun.Float
	CavgEmotionValue goaliyun.Float
}

type GetAudioDataStatusServiceSrStat struct {
	Srole          goaliyun.Integer
	SmaxSpeechRate goaliyun.Float
	SminSpeechRate goaliyun.Float
	SavgSpeechRate goaliyun.Float
}

type GetAudioDataStatusClientSrStat struct {
	Crole          goaliyun.Integer
	CmaxSpeechRate goaliyun.Float
	CminSpeechRate goaliyun.Float
	CavgSpeechRate goaliyun.Float
}

type GetAudioDataStatusTaskAsrResultList []GetAudioDataStatusTaskAsrResult

func (list *GetAudioDataStatusTaskAsrResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetAudioDataStatusTaskAsrResult)
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

type GetAudioDataStatusSentenceResultList []GetAudioDataStatusSentenceResult

func (list *GetAudioDataStatusSentenceResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetAudioDataStatusSentenceResult)
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
