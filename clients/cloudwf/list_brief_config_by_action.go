package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type ListBriefConfigByActionRequest struct {
	AncestorApgroupId int64  `position:"Query" name:"AncestorApgroupId"`
	Limit             int64  `position:"Query" name:"Limit"`
	FuzzySearch       string `position:"Query" name:"FuzzySearch"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *ListBriefConfigByActionRequest) Invoke(client goaliyun.Client) (*ListBriefConfigByActionResponse, error) {
	resp := &ListBriefConfigByActionResponse{}
	err := client.Request("cloudwf", "ListBriefConfigByAction", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListBriefConfigByActionResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
