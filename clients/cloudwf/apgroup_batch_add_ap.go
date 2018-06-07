package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type ApgroupBatchAddApRequest struct {
	ApAssetIds string `position:"Query" name:"ApAssetIds"`
	ApgroupId  int64  `position:"Query" name:"ApgroupId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ApgroupBatchAddApRequest) Invoke(client goaliyun.Client) (*ApgroupBatchAddApResponse, error) {
	resp := &ApgroupBatchAddApResponse{}
	err := client.Request("cloudwf", "ApgroupBatchAddAp", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ApgroupBatchAddApResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
