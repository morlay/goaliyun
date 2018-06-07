package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type ResetApRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ResetApRequest) Invoke(client goaliyun.Client) (*ResetApResponse, error) {
	resp := &ResetApResponse{}
	err := client.Request("cloudwf", "ResetAp", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResetApResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
