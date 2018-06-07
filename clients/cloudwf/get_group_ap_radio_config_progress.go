package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type GetGroupApRadioConfigProgressRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetGroupApRadioConfigProgressRequest) Invoke(client goaliyun.Client) (*GetGroupApRadioConfigProgressResponse, error) {
	resp := &GetGroupApRadioConfigProgressResponse{}
	err := client.Request("cloudwf", "GetGroupApRadioConfigProgress", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetGroupApRadioConfigProgressResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
