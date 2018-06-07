package saf

import (
	"github.com/morlay/goaliyun"
)

type ExecuteRequestRequest struct {
	ServiceParameters string `position:"Query" name:"ServiceParameters"`
	Service           string `position:"Query" name:"Service"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *ExecuteRequestRequest) Invoke(client goaliyun.Client) (*ExecuteRequestResponse, error) {
	resp := &ExecuteRequestResponse{}
	err := client.Request("saf", "ExecuteRequest", "2017-03-31", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ExecuteRequestResponse struct {
	Code      goaliyun.Integer
	Data      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
}
