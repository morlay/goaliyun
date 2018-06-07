package iot

import (
	"github.com/morlay/goaliyun"
)

type QueryDeviceStatisticsRequest struct {
	ProductKey string `position:"Query" name:"ProductKey"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *QueryDeviceStatisticsRequest) Invoke(client goaliyun.Client) (*QueryDeviceStatisticsResponse, error) {
	resp := &QueryDeviceStatisticsResponse{}
	err := client.Request("iot", "QueryDeviceStatistics", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryDeviceStatisticsResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
	Data         QueryDeviceStatisticsData
}

type QueryDeviceStatisticsData struct {
	DeviceCount goaliyun.Integer
	OnlineCount goaliyun.Integer
	ActiveCount goaliyun.Integer
}
