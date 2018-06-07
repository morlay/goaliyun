package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsMessageGetByMsgIdRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	MsgId        string `position:"Query" name:"MsgId"`
	Topic        string `position:"Query" name:"Topic"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsMessageGetByMsgIdRequest) Invoke(client goaliyun.Client) (*OnsMessageGetByMsgIdResponse, error) {
	resp := &OnsMessageGetByMsgIdResponse{}
	err := client.Request("ons", "OnsMessageGetByMsgId", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsMessageGetByMsgIdResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsMessageGetByMsgIdData
}

type OnsMessageGetByMsgIdData struct {
	Topic          goaliyun.String
	Flag           goaliyun.Integer
	Body           goaliyun.String
	StoreSize      goaliyun.Integer
	BornTimestamp  goaliyun.Integer
	BornHost       goaliyun.String
	StoreTimestamp goaliyun.Integer
	StoreHost      goaliyun.String
	MsgId          goaliyun.String
	OffsetId       goaliyun.String
	BodyCRC        goaliyun.Integer
	ReconsumeTimes goaliyun.Integer
	PropertyList   OnsMessageGetByMsgIdMessagePropertyList
}

type OnsMessageGetByMsgIdMessageProperty struct {
	Name  goaliyun.String
	Value goaliyun.String
}

type OnsMessageGetByMsgIdMessagePropertyList []OnsMessageGetByMsgIdMessageProperty

func (list *OnsMessageGetByMsgIdMessagePropertyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMessageGetByMsgIdMessageProperty)
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
