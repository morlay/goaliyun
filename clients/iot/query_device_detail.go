package iot

import (
	"github.com/morlay/goaliyun"
)

type QueryDeviceDetailRequest struct {
	IotId      string `position:"Query" name:"IotId"`
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *QueryDeviceDetailRequest) Invoke(client goaliyun.Client) (*QueryDeviceDetailResponse, error) {
	resp := &QueryDeviceDetailResponse{}
	err := client.Request("iot", "QueryDeviceDetail", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryDeviceDetailResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
	Data         QueryDeviceDetailData
}

type QueryDeviceDetailData struct {
	IotId           goaliyun.String
	ProductKey      goaliyun.String
	ProductName     goaliyun.String
	DeviceName      goaliyun.String
	DeviceSecret    goaliyun.String
	FirmwareVersion goaliyun.String
	GmtCreate       goaliyun.String
	GmtActive       goaliyun.String
	GmtOnline       goaliyun.String
	Status          goaliyun.String
	IpAddress       goaliyun.String
	NodeType        goaliyun.Integer
	Region          goaliyun.String
}
