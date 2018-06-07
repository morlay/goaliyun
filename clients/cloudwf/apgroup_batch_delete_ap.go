package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type ApgroupBatchDeleteApRequest struct {
	ApAssetIds string `position:"Query" name:"ApAssetIds"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ApgroupBatchDeleteApRequest) Invoke(client goaliyun.Client) (*ApgroupBatchDeleteApResponse, error) {
	resp := &ApgroupBatchDeleteApResponse{}
	err := client.Request("cloudwf", "ApgroupBatchDeleteAp", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ApgroupBatchDeleteApResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
