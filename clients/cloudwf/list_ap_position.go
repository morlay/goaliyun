package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type ListApPositionRequest struct {
	MapId    int64  `position:"Query" name:"MapId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListApPositionRequest) Invoke(client goaliyun.Client) (*ListApPositionResponse, error) {
	resp := &ListApPositionResponse{}
	err := client.Request("cloudwf", "ListApPosition", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListApPositionResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
