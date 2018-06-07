package ccc

import (
	"github.com/morlay/goaliyun"
)

type ModifyPhoneNumberRequest struct {
	ContactFlowId string `position:"Query" name:"ContactFlowId"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	PhoneNumberId string `position:"Query" name:"PhoneNumberId"`
	Usage         string `position:"Query" name:"Usage"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *ModifyPhoneNumberRequest) Invoke(client goaliyun.Client) (*ModifyPhoneNumberResponse, error) {
	resp := &ModifyPhoneNumberResponse{}
	err := client.Request("ccc", "ModifyPhoneNumber", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyPhoneNumberResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	PhoneNumber    ModifyPhoneNumberPhoneNumber
}

type ModifyPhoneNumberPhoneNumber struct {
	PhoneNumberId          goaliyun.String
	InstanceId             goaliyun.String
	Number                 goaliyun.String
	PhoneNumberDescription goaliyun.String
	TestOnly               bool
	RemainingTime          goaliyun.Integer
	AllowOutbound          bool
	Usage                  goaliyun.String
	Trunks                 goaliyun.Integer
	ContactFlow            ModifyPhoneNumberContactFlow
}

type ModifyPhoneNumberContactFlow struct {
	ContactFlowId          goaliyun.String
	InstanceId             goaliyun.String
	ContactFlowName        goaliyun.String
	ContactFlowDescription goaliyun.String
	Type                   goaliyun.String
}
