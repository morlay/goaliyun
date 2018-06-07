package ons

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type OnsMessageGetByKeyRequest struct {
	PreventCache int64  `position:"Query" name:"PreventCache"`
	OnsRegionId  string `position:"Query" name:"OnsRegionId"`
	OnsPlatform  string `position:"Query" name:"OnsPlatform"`
	Topic        string `position:"Query" name:"Topic"`
	Key          string `position:"Query" name:"Key"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OnsMessageGetByKeyRequest) Invoke(client goaliyun.Client) (*OnsMessageGetByKeyResponse, error) {
	resp := &OnsMessageGetByKeyResponse{}
	err := client.Request("ons", "OnsMessageGetByKey", "2017-09-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OnsMessageGetByKeyResponse struct {
	RequestId goaliyun.String
	HelpUrl   goaliyun.String
	Data      OnsMessageGetByKeyOnsRestMessageDoList
}

type OnsMessageGetByKeyOnsRestMessageDo struct {
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
	PropertyList   OnsMessageGetByKeyMessagePropertyList
}

type OnsMessageGetByKeyMessageProperty struct {
	Name  goaliyun.String
	Value goaliyun.String
}

type OnsMessageGetByKeyOnsRestMessageDoList []OnsMessageGetByKeyOnsRestMessageDo

func (list *OnsMessageGetByKeyOnsRestMessageDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMessageGetByKeyOnsRestMessageDo)
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

type OnsMessageGetByKeyMessagePropertyList []OnsMessageGetByKeyMessageProperty

func (list *OnsMessageGetByKeyMessagePropertyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]OnsMessageGetByKeyMessageProperty)
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
