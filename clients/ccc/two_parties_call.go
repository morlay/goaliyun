package ccc

import (
	"github.com/morlay/goaliyun"
)

type TwoPartiesCallRequest struct {
	Caller         string `position:"Query" name:"Caller"`
	InstanceId     string `position:"Query" name:"InstanceId"`
	CalleeCustomer string `position:"Query" name:"CalleeCustomer"`
	CalleeAgent    string `position:"Query" name:"CalleeAgent"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *TwoPartiesCallRequest) Invoke(client goaliyun.Client) (*TwoPartiesCallResponse, error) {
	resp := &TwoPartiesCallResponse{}
	err := client.Request("ccc", "TwoPartiesCall", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type TwoPartiesCallResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	StatusCode     goaliyun.String
	StatusDesc     goaliyun.String
	TaskId         goaliyun.String
	TimeStamp      goaliyun.String
}
