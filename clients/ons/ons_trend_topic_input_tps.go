package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsTrendTopicInputTpsRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	Period       int64  `position:"Query" name:"Period"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	BeginTime    int64  `position:"Query" name:"BeginTime"`
	Type         int64  `position:"Query" name:"Type"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsTrendTopicInputTpsRequest) Invoke(client goaliyun.Client) (*OnsTrendTopicInputTpsResponse, error) {
	resp := &OnsTrendTopicInputTpsResponse{}
	err := client.Request("ons", "OnsTrendTopicInputTps", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsTrendTopicInputTpsResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsTrendTopicInputTpsData
}

type OnsTrendTopicInputTpsData struct {
	Title   goaliyun.String
	XUnit   goaliyun.String
	YUnit   goaliyun.String
	Records OnsTrendTopicInputTpsStatsDataDoList
}

type OnsTrendTopicInputTpsStatsDataDo struct {
	X goaliyun.Integer
	Y goaliyun.Float
}

type OnsTrendTopicInputTpsStatsDataDoList []OnsTrendTopicInputTpsStatsDataDo

func (list *OnsTrendTopicInputTpsStatsDataDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTrendTopicInputTpsStatsDataDo)
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
