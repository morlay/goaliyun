package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListPhoneNumbersRequest struct {
	OutboundOnly string `position:"Query" name:"OutboundOnly"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *ListPhoneNumbersRequest) Invoke(client goaliyun.Client) (*ListPhoneNumbersResponse, error) {
	resp := &ListPhoneNumbersResponse{}
	err := client.Request("ccc", "ListPhoneNumbers", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListPhoneNumbersResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	PhoneNumbers   ListPhoneNumbersPhoneNumberList
}

type ListPhoneNumbersPhoneNumber struct {
	PhoneNumberId          goaliyun.String
	InstanceId             goaliyun.String
	Number                 goaliyun.String
	PhoneNumberDescription goaliyun.String
	TestOnly               bool
	RemainingTime          goaliyun.Integer
	AllowOutbound          bool
	Usage                  goaliyun.String
	Trunks                 goaliyun.Integer
	Province               goaliyun.String
	City                   goaliyun.String
	ContactFlow            ListPhoneNumbersContactFlow
}

type ListPhoneNumbersContactFlow struct {
	ContactFlowId          goaliyun.String
	InstanceId             goaliyun.String
	ContactFlowName        goaliyun.String
	ContactFlowDescription goaliyun.String
	Type                   goaliyun.String
}

type ListPhoneNumbersPhoneNumberList []ListPhoneNumbersPhoneNumber

func (list *ListPhoneNumbersPhoneNumberList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPhoneNumbersPhoneNumber)
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
