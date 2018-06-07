package aegis

import (
	"github.com/morlay/goaliyun"
)

type GetStatisticsRequest struct {
	EndTime   string `position:"Query" name:"EndTime"`
	StartTime string `position:"Query" name:"StartTime"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetStatisticsRequest) Invoke(client goaliyun.Client) (*GetStatisticsResponse, error) {
	resp := &GetStatisticsResponse{}
	err := client.Request("aegis", "GetStatistics", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetStatisticsResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      GetStatisticsData
}

type GetStatisticsData struct {
	Account goaliyun.Integer
	Health  goaliyun.Integer
	Patch   goaliyun.Integer
	Trojan  goaliyun.Integer
}
