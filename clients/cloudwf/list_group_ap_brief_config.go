package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type ListGroupApBriefConfigRequest struct {
	OrderCol   string `position:"Query" name:"OrderCol"`
	SearchName string `position:"Query" name:"SearchName"`
	ApgroupId  int64  `position:"Query" name:"ApgroupId"`
	ColCnt     int64  `position:"Query" name:"ColCnt"`
	Length     int64  `position:"Query" name:"Length"`
	PageIndex  int64  `position:"Query" name:"PageIndex"`
	SearchMac  string `position:"Query" name:"SearchMac"`
	OrderDir   string `position:"Query" name:"OrderDir"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListGroupApBriefConfigRequest) Invoke(client goaliyun.Client) (*ListGroupApBriefConfigResponse, error) {
	resp := &ListGroupApBriefConfigResponse{}
	err := client.Request("cloudwf", "ListGroupApBriefConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListGroupApBriefConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
