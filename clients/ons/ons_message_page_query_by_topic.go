package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsMessagePageQueryByTopicRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	Topic        string `position:"Query" name:"Topic"`
	EndTime      int64  `position:"Query" name:"EndTime"`
	BeginTime    int64  `position:"Query" name:"BeginTime"`
	CurrentPage  int64  `position:"Query" name:"CurrentPage"`
	TaskId       string `position:"Query" name:"TaskId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsMessagePageQueryByTopicRequest) Invoke(client goaliyun.Client) (*OnsMessagePageQueryByTopicResponse, error) {
	resp := &OnsMessagePageQueryByTopicResponse{}
	err := client.Request("ons", "OnsMessagePageQueryByTopic", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsMessagePageQueryByTopicResponse struct {
	RequestId  goaliyun.String
	HelpUrl    goaliyun.String
	MsgFoundDo OnsMessagePageQueryByTopicMsgFoundDo
}

type OnsMessagePageQueryByTopicMsgFoundDo struct {
	TaskId       goaliyun.String
	MaxPageCount goaliyun.Integer
	CurrentPage  goaliyun.Integer
	MsgFoundList OnsMessagePageQueryByTopicOnsRestMessageDoList
}

type OnsMessagePageQueryByTopicOnsRestMessageDo struct {
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
	PropertyList   OnsMessagePageQueryByTopicMessagePropertyList
}

type OnsMessagePageQueryByTopicMessageProperty struct {
	Name  goaliyun.String
	Value goaliyun.String
}

type OnsMessagePageQueryByTopicOnsRestMessageDoList []OnsMessagePageQueryByTopicOnsRestMessageDo

func (list *OnsMessagePageQueryByTopicOnsRestMessageDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMessagePageQueryByTopicOnsRestMessageDo)
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

type OnsMessagePageQueryByTopicMessagePropertyList []OnsMessagePageQueryByTopicMessageProperty

func (list *OnsMessagePageQueryByTopicMessagePropertyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMessagePageQueryByTopicMessageProperty)
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
