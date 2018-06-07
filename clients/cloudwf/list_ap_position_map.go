package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type ListApPositionMapRequest struct {
	OrderCol          string `position:"Query" name:"OrderCol"`
	SearchName        string `position:"Query" name:"SearchName"`
	TotalItem         int64  `position:"Query" name:"TotalItem"`
	Length            int64  `position:"Query" name:"Length"`
	MapType           int64  `position:"Query" name:"MapType"`
	PageIndex         int64  `position:"Query" name:"PageIndex"`
	SearchApgroupName string `position:"Query" name:"SearchApgroupName"`
	OrderDir          string `position:"Query" name:"OrderDir"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *ListApPositionMapRequest) Invoke(client goaliyun.Client) (*ListApPositionMapResponse, error) {
	resp := &ListApPositionMapResponse{}
	err := client.Request("cloudwf", "ListApPositionMap", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListApPositionMapResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
