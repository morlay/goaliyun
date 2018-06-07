package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsMqttQueryMsgTransTrendRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Qos          int64  `position:"Query" name:"Qos"`
	TransType    string `position:"Query" name:"TransType"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	BeginTime    int64  `position:"Query" name:"BeginTime"`
	TpsType      string `position:"Query" name:"TpsType"`
	ParentTopic  string `position:"Query" name:"ParentTopic"`
	MsgType      string `position:"Query" name:"MsgType"`
	SubTopic     string `position:"Query" name:"SubTopic"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsMqttQueryMsgTransTrendRequest) Invoke(client goaliyun.Client) (*OnsMqttQueryMsgTransTrendResponse, error) {
	resp := &OnsMqttQueryMsgTransTrendResponse{}
	err := client.Request("ons", "OnsMqttQueryMsgTransTrend", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsMqttQueryMsgTransTrendResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsMqttQueryMsgTransTrendData
}

type OnsMqttQueryMsgTransTrendData struct {
	Title   goaliyun.String
	XUnit   goaliyun.String
	YUnit   goaliyun.String
	Records OnsMqttQueryMsgTransTrendStatsDataDoList
}

type OnsMqttQueryMsgTransTrendStatsDataDo struct {
	X goaliyun.Integer
	Y goaliyun.Float
}

type OnsMqttQueryMsgTransTrendStatsDataDoList []OnsMqttQueryMsgTransTrendStatsDataDo

func (list *OnsMqttQueryMsgTransTrendStatsDataDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMqttQueryMsgTransTrendStatsDataDo)
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
