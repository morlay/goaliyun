package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsConsumerAccumulateRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	Detail       string `position:"Query" name:"Detail"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsConsumerAccumulateRequest) Invoke(client goaliyun.Client) (*OnsConsumerAccumulateResponse, error) {
	resp := &OnsConsumerAccumulateResponse{}
	err := client.Request("ons", "OnsConsumerAccumulate", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsConsumerAccumulateResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsConsumerAccumulateData
}

type OnsConsumerAccumulateData struct {
	Online            bool
	TotalDiff         goaliyun.Integer
	ConsumeTps        goaliyun.Float
	LastTimestamp     goaliyun.Integer
	DelayTime         goaliyun.Integer
	DetailInTopicList OnsConsumerAccumulateDetailInTopicDoList
}

type OnsConsumerAccumulateDetailInTopicDo struct {
	Topic         goaliyun.String
	TotalDiff     goaliyun.Integer
	LastTimestamp goaliyun.Integer
	DelayTime     goaliyun.Integer
}

type OnsConsumerAccumulateDetailInTopicDoList []OnsConsumerAccumulateDetailInTopicDo

func (list *OnsConsumerAccumulateDetailInTopicDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsConsumerAccumulateDetailInTopicDo)
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
