package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type GetLatestApStatisticRequest struct {
	ApgroupId int64  `position:"Query" name:"ApgroupId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetLatestApStatisticRequest) Invoke(client goaliyun.Client) (*GetLatestApStatisticResponse, error) {
	resp := &GetLatestApStatisticResponse{}
	err := client.Request("cloudwf", "GetLatestApStatistic", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetLatestApStatisticResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
