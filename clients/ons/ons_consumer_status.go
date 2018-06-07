package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsConsumerStatusRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	NeedJstack   string `position:"Query" name:"NeedJstack"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	Detail       string `position:"Query" name:"Detail"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsConsumerStatusRequest) Invoke(client goaliyun.Client) (*OnsConsumerStatusResponse, error) {
	resp := &OnsConsumerStatusResponse{}
	err := client.Request("ons", "OnsConsumerStatus", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsConsumerStatusResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsConsumerStatusData
}

type OnsConsumerStatusData struct {
	Online                     bool
	TotalDiff                  goaliyun.Integer
	ConsumeTps                 goaliyun.Float
	LastTimestamp              goaliyun.Integer
	DelayTime                  goaliyun.Integer
	ConsumeModel               goaliyun.String
	SubscriptionSame           bool
	RebalanceOK                bool
	ConnectionSet              OnsConsumerStatusConnectionDoList
	DetailInTopicList          OnsConsumerStatusDetailInTopicDoList
	ConsumerConnectionInfoList OnsConsumerStatusConsumerConnectionInfoDoList
}

type OnsConsumerStatusConnectionDo struct {
	ClientId   goaliyun.String
	ClientAddr goaliyun.String
	Language   goaliyun.String
	Version    goaliyun.String
}

type OnsConsumerStatusDetailInTopicDo struct {
	Topic         goaliyun.String
	TotalDiff     goaliyun.Integer
	LastTimestamp goaliyun.Integer
	DelayTime     goaliyun.Integer
}

type OnsConsumerStatusConsumerConnectionInfoDo struct {
	ClientId        goaliyun.String
	Connection      goaliyun.String
	Language        goaliyun.String
	Version         goaliyun.String
	ConsumeModel    goaliyun.String
	ConsumeType     goaliyun.String
	ThreadCount     goaliyun.Integer
	StartTimeStamp  goaliyun.Integer
	LastTimeStamp   goaliyun.Integer
	SubscriptionSet OnsConsumerStatusSubscriptionDataList
	RunningDataList OnsConsumerStatusConsumerRunningDataDoList
	Jstack          OnsConsumerStatusThreadTrackDoList
}

type OnsConsumerStatusSubscriptionData struct {
	Topic      goaliyun.String
	SubString  goaliyun.String
	SubVersion goaliyun.Integer
	TagsSet    OnsConsumerStatusTagsSetList
}

type OnsConsumerStatusConsumerRunningDataDo struct {
	ConsumerId         goaliyun.String
	Topic              goaliyun.String
	Rt                 goaliyun.Float
	OkTps              goaliyun.Float
	FailedTps          goaliyun.Float
	FailedCountPerHour goaliyun.Integer
}

type OnsConsumerStatusThreadTrackDo struct {
	Thread    goaliyun.String
	TrackList OnsConsumerStatusTrackListList
}

type OnsConsumerStatusConnectionDoList []OnsConsumerStatusConnectionDo

func (list *OnsConsumerStatusConnectionDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerStatusConnectionDo)
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

type OnsConsumerStatusDetailInTopicDoList []OnsConsumerStatusDetailInTopicDo

func (list *OnsConsumerStatusDetailInTopicDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerStatusDetailInTopicDo)
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

type OnsConsumerStatusConsumerConnectionInfoDoList []OnsConsumerStatusConsumerConnectionInfoDo

func (list *OnsConsumerStatusConsumerConnectionInfoDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerStatusConsumerConnectionInfoDo)
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

type OnsConsumerStatusSubscriptionDataList []OnsConsumerStatusSubscriptionData

func (list *OnsConsumerStatusSubscriptionDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerStatusSubscriptionData)
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

type OnsConsumerStatusConsumerRunningDataDoList []OnsConsumerStatusConsumerRunningDataDo

func (list *OnsConsumerStatusConsumerRunningDataDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerStatusConsumerRunningDataDo)
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

type OnsConsumerStatusThreadTrackDoList []OnsConsumerStatusThreadTrackDo

func (list *OnsConsumerStatusThreadTrackDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerStatusThreadTrackDo)
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

type OnsConsumerStatusTagsSetList []goaliyun.String

func (list *OnsConsumerStatusTagsSetList) UnmarshalJSON(data []byte) error {
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

type OnsConsumerStatusTrackListList []goaliyun.String

func (list *OnsConsumerStatusTrackListList) UnmarshalJSON(data []byte) error {
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
