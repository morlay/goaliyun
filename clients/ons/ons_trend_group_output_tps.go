package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsTrendGroupOutputTpsRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	Period       int64  `position:"Query" name:"Period"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	ConsumerId   string `position:"Query" name:"ConsumerId"`
	Topic        string `position:"Query" name:"Topic"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	BeginTime    int64  `position:"Query" name:"BeginTime"`
	Type         int64  `position:"Query" name:"Type"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsTrendGroupOutputTpsRequest) Invoke(client goaliyun.Client) (*OnsTrendGroupOutputTpsResponse, error) {
	resp := &OnsTrendGroupOutputTpsResponse{}
	err := client.Request("ons", "OnsTrendGroupOutputTps", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsTrendGroupOutputTpsResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsTrendGroupOutputTpsData
}

type OnsTrendGroupOutputTpsData struct {
	Title   goaliyun.String
	XUnit   goaliyun.String
	YUnit   goaliyun.String
	Records OnsTrendGroupOutputTpsStatsDataDoList
}

type OnsTrendGroupOutputTpsStatsDataDo struct {
	X goaliyun.Integer
	Y goaliyun.Float
}

type OnsTrendGroupOutputTpsStatsDataDoList []OnsTrendGroupOutputTpsStatsDataDo

func (list *OnsTrendGroupOutputTpsStatsDataDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsTrendGroupOutputTpsStatsDataDo)
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
