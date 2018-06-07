package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type BatchRegisterApAssetRequest struct {
	JsonData string `position:"Query" name:"JsonData"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *BatchRegisterApAssetRequest) Invoke(client goaliyun.Client) (*BatchRegisterApAssetResponse, error) {
	resp := &BatchRegisterApAssetResponse{}
	err := client.Request("cloudwf", "BatchRegisterApAsset", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type BatchRegisterApAssetResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
