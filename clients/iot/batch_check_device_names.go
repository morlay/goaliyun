package iot

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type BatchCheckDeviceNamesRequest struct {
	DeviceNames *BatchCheckDeviceNamesDeviceNameList `position:"Query" type:"Repeated" name:"DeviceName"`
	ProductKey  string                               `position:"Query" name:"ProductKey"`
	RegionId    string                               `position:"Query" name:"RegionId"`
}

func (req *BatchCheckDeviceNamesRequest) Invoke(client goaliyun.Client) (*BatchCheckDeviceNamesResponse, error) {
	resp := &BatchCheckDeviceNamesResponse{}
	err := client.Request("iot", "BatchCheckDeviceNames", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type BatchCheckDeviceNamesResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
	Data         BatchCheckDeviceNamesData
}

type BatchCheckDeviceNamesData struct {
	ApplyId goaliyun.Integer
}

type BatchCheckDeviceNamesDeviceNameList []string

func (list *BatchCheckDeviceNamesDeviceNameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
