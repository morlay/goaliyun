package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type ListConfigByActionRequest struct {
	SearchName string `position:"Query" name:"SearchName"`
	Limit      int64  `position:"Query" name:"Limit"`
	ActionName string `position:"Query" name:"ActionName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListConfigByActionRequest) Invoke(client goaliyun.Client) (*ListConfigByActionResponse, error) {
	resp := &ListConfigByActionResponse{}
	err := client.Request("cloudwf", "ListConfigByAction", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListConfigByActionResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
