package ccc

import (
	"github.com/morlay/goaliyun"
)

type RemovePhoneNumberRequest struct {
	InstanceId    string `position:"Query" name:"InstanceId"`
	PhoneNumberId string `position:"Query" name:"PhoneNumberId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *RemovePhoneNumberRequest) Invoke(client goaliyun.Client) (*RemovePhoneNumberResponse, error) {
	resp := &RemovePhoneNumberResponse{}
	err := client.Request("ccc", "RemovePhoneNumber", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemovePhoneNumberResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
}
