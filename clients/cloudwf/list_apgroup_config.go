package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type ListApgroupConfigRequest struct {
	OrderCol      string `position:"Query" name:"OrderCol"`
	SearchName    string `position:"Query" name:"SearchName"`
	SearchCompany string `position:"Query" name:"SearchCompany"`
	Length        int64  `position:"Query" name:"Length"`
	PageIndex     int64  `position:"Query" name:"PageIndex"`
	OrderDir      string `position:"Query" name:"OrderDir"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *ListApgroupConfigRequest) Invoke(client goaliyun.Client) (*ListApgroupConfigResponse, error) {
	resp := &ListApgroupConfigResponse{}
	err := client.Request("cloudwf", "ListApgroupConfig", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListApgroupConfigResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
