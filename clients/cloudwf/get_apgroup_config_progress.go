package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type GetApgroupConfigProgressRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetApgroupConfigProgressRequest) Invoke(client goaliyun.Client) (*GetApgroupConfigProgressResponse, error) {
	resp := &GetApgroupConfigProgressResponse{}
	err := client.Request("cloudwf", "GetApgroupConfigProgress", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetApgroupConfigProgressResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
