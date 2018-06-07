package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type GetGroupApRepairProgressRequest struct {
	Id       int64  `position:"Query" name:"Id"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetGroupApRepairProgressRequest) Invoke(client goaliyun.Client) (*GetGroupApRepairProgressResponse, error) {
	resp := &GetGroupApRepairProgressResponse{}
	err := client.Request("cloudwf", "GetGroupApRepairProgress", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetGroupApRepairProgressResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
