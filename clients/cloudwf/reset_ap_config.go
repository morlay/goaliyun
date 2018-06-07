package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type ResetApConfigRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ResetApConfigRequest) Invoke(client goaliyun.Client) (*ResetApConfigResponse, error) {
	resp := &ResetApConfigResponse{}
	err := client.Request("cloudwf", "ResetApConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResetApConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
