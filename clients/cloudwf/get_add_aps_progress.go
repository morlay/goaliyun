package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type GetAddApsProgressRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetAddApsProgressRequest) Invoke(client goaliyun.Client) (*GetAddApsProgressResponse, error) {
	resp := &GetAddApsProgressResponse{}
	err := client.Request("cloudwf", "GetAddApsProgress", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetAddApsProgressResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
