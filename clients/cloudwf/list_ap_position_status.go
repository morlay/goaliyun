package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type ListApPositionStatusRequest struct {
	JsonData string `position:"Query" name:"JsonData"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListApPositionStatusRequest) Invoke(client goaliyun.Client) (*ListApPositionStatusResponse, error) {
	resp := &ListApPositionStatusResponse{}
	err := client.Request("cloudwf", "ListApPositionStatus", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListApPositionStatusResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
