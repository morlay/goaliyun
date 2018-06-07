package ccc

import (
	"github.com/morlay/goaliyun"
)

type LaunchAppraiseRequest struct {
	Acid       string `position:"Query" name:"Acid"`
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *LaunchAppraiseRequest) Invoke(client goaliyun.Client) (*LaunchAppraiseResponse, error) {
	resp := &LaunchAppraiseResponse{}
	err := client.Request("ccc", "LaunchAppraise", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type LaunchAppraiseResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
}
