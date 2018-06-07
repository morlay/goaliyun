package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type GetApgroupSsidConfigProgressRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetApgroupSsidConfigProgressRequest) Invoke(client goaliyun.Client) (*GetApgroupSsidConfigProgressResponse, error) {
	resp := &GetApgroupSsidConfigProgressResponse{}
	err := client.Request("cloudwf", "GetApgroupSsidConfigProgress", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetApgroupSsidConfigProgressResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
