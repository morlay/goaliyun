package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type FindApRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *FindApRequest) Invoke(client goaliyun.Client) (*FindApResponse, error) {
	resp := &FindApResponse{}
	err := client.Request("cloudwf", "FindAp", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FindApResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
