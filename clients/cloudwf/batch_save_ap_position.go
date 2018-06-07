package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type BatchSaveApPositionRequest struct {
	JsonData string `position:"Query" name:"JsonData"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *BatchSaveApPositionRequest) Invoke(client goaliyun.Client) (*BatchSaveApPositionResponse, error) {
	resp := &BatchSaveApPositionResponse{}
	err := client.Request("cloudwf", "BatchSaveApPosition", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type BatchSaveApPositionResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
