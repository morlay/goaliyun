package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsMqttQueryHistoryOnlineRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	GroupId      string `position:"Query" name:"GroupId"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	BeginTime    int64  `position:"Query" name:"BeginTime"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsMqttQueryHistoryOnlineRequest) Invoke(client goaliyun.Client) (*OnsMqttQueryHistoryOnlineResponse, error) {
	resp := &OnsMqttQueryHistoryOnlineResponse{}
	err := client.Request("ons", "OnsMqttQueryHistoryOnline", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsMqttQueryHistoryOnlineResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsMqttQueryHistoryOnlineData
}

type OnsMqttQueryHistoryOnlineData struct {
	Title   goaliyun.String
	XUnit   goaliyun.String
	YUnit   goaliyun.String
	Records OnsMqttQueryHistoryOnlineStatsDataDoList
}

type OnsMqttQueryHistoryOnlineStatsDataDo struct {
	X goaliyun.Integer
	Y goaliyun.Float
}

type OnsMqttQueryHistoryOnlineStatsDataDoList []OnsMqttQueryHistoryOnlineStatsDataDo

func (list *OnsMqttQueryHistoryOnlineStatsDataDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMqttQueryHistoryOnlineStatsDataDo)
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
