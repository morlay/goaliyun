package aegis

import (
	"github.com/morlay/goaliyun"
)

type GetAccountStatisticsRequest struct {
	EndTime   string `position:"Query" name:"EndTime"`
	StartTime string `position:"Query" name:"StartTime"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetAccountStatisticsRequest) Invoke(client goaliyun.Client) (*GetAccountStatisticsResponse, error) {
	resp := &GetAccountStatisticsResponse{}
	err := client.Request("aegis", "GetAccountStatistics", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetAccountStatisticsResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Success   bool
	Message   goaliyun.String
	Data      GetAccountStatisticsData
}

type GetAccountStatisticsData struct {
	RemoteLogin  goaliyun.Integer
	CrackSuccess goaliyun.Integer
}
