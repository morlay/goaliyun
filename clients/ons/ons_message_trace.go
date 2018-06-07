package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsMessageTraceRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	MsgId        string `position:"Query" name:"MsgId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsMessageTraceRequest) Invoke(client goaliyun.Client) (*OnsMessageTraceResponse, error) {
	resp := &OnsMessageTraceResponse{}
	err := client.Request("ons", "OnsMessageTrace", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsMessageTraceResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsMessageTraceMessageTrackList
}

type OnsMessageTraceMessageTrack struct {
	ConsumerGroup goaliyun.String
	TrackType     goaliyun.String
	ExceptionDesc goaliyun.String
}

type OnsMessageTraceMessageTrackList []OnsMessageTraceMessageTrack

func (list *OnsMessageTraceMessageTrackList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMessageTraceMessageTrack)
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
