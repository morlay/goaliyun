package aegis

import (
	"github.com/morlay/goaliyun"
)

type GetCrackStatisticsRequest struct {
	EndTime   string `position:"Query" name:"EndTime"`
	StartTime string `position:"Query" name:"StartTime"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetCrackStatisticsRequest) Invoke(client goaliyun.Client) (*GetCrackStatisticsResponse, error) {
	resp := &GetCrackStatisticsResponse{}
	err := client.Request("aegis", "GetCrackStatistics", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetCrackStatisticsResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      GetCrackStatisticsData
}

type GetCrackStatisticsData struct {
	Intercepted goaliyun.Integer
}
