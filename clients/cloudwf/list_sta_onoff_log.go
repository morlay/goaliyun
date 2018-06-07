package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type ListStaOnoffLogRequest struct {
	OrderCol       string `position:"Query" name:"OrderCol"`
	SearchSsid     string `position:"Query" name:"SearchSsid"`
	SearchApName   string `position:"Query" name:"SearchApName"`
	Length         int64  `position:"Query" name:"Length"`
	SearchUsername string `position:"Query" name:"SearchUsername"`
	PageIndex      int64  `position:"Query" name:"PageIndex"`
	Id             int64  `position:"Query" name:"Id"`
	OrderDir       string `position:"Query" name:"OrderDir"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *ListStaOnoffLogRequest) Invoke(client goaliyun.Client) (*ListStaOnoffLogResponse, error) {
	resp := &ListStaOnoffLogResponse{}
	err := client.Request("cloudwf", "ListStaOnoffLog", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListStaOnoffLogResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
