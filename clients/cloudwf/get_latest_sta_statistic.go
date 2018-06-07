package cloudwf

import (
	"github.com/morlay/goaliyun"
)

type GetLatestStaStatisticRequest struct {
	ApgroupId int64  `position:"Query" name:"ApgroupId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetLatestStaStatisticRequest) Invoke(client goaliyun.Client) (*GetLatestStaStatisticResponse, error) {
	resp := &GetLatestStaStatisticResponse{}
	err := client.Request("cloudwf", "GetLatestStaStatistic", "2017-03-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetLatestStaStatisticResponse struct {
	RequestId goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
}
