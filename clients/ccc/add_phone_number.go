package ccc

import (
	"github.com/morlay/goaliyun"
)

type AddPhoneNumberRequest struct {
	ContactFlowId string `position:"Query" name:"ContactFlowId"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	Usage         string `position:"Query" name:"Usage"`
	PhoneNumber   string `position:"Query" name:"PhoneNumber"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *AddPhoneNumberRequest) Invoke(client goaliyun.Client) (*AddPhoneNumberResponse, error) {
	resp := &AddPhoneNumberResponse{}
	err := client.Request("ccc", "AddPhoneNumber", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddPhoneNumberResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	PhoneNumber    AddPhoneNumberPhoneNumber
}

type AddPhoneNumberPhoneNumber struct {
	PhoneNumberId          goaliyun.String
	InstanceId             goaliyun.String
	Number                 goaliyun.String
	PhoneNumberDescription goaliyun.String
	TestOnly               bool
	RemainingTime          goaliyun.Integer
	AllowOutbound          bool
	Usage                  goaliyun.String
	Trunks                 goaliyun.Integer
	ContactFlow            AddPhoneNumberContactFlow
}

type AddPhoneNumberContactFlow struct {
	ContactFlowId          goaliyun.String
	InstanceId             goaliyun.String
	ContactFlowName        goaliyun.String
	ContactFlowDescription goaliyun.String
	Type                   goaliyun.String
}
