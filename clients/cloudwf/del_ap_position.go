package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type DelApPositionRequest struct {
	ApAssetId int64  `position:"Query" name:"ApAssetId"`
	MapId     int64  `position:"Query" name:"MapId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *DelApPositionRequest) Invoke(client goaliyun.Client) (*DelApPositionResponse, error) {
	resp := &DelApPositionResponse{}
	err := client.Request("cloudwf", "DelApPosition", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DelApPositionResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
