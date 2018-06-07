package iot

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryDevicePropertyDataRequest struct {
	Asc        int64  `position:"Query" name:"Asc"`
	Identifier string `position:"Query" name:"Identifier"`
	IotId      string `position:"Query" name:"IotId"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	EndTime    int64  `position:"Query" name:"EndTime"`
	DeviceName string `position:"Query" name:"DeviceName"`
	StartTime  int64  `position:"Query" name:"StartTime"`
	ProductKey string `position:"Query" name:"ProductKey"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *QueryDevicePropertyDataRequest) Invoke(client goaliyun.Client) (*QueryDevicePropertyDataResponse, error) {
	resp := &QueryDevicePropertyDataResponse{}
	err := client.Request("iot", "QueryDevicePropertyData", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryDevicePropertyDataResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
	Data         QueryDevicePropertyDataData
}

type QueryDevicePropertyDataData struct {
	NextValid bool
	NextTime  goaliyun.Integer
	List      QueryDevicePropertyDataPropertyInfoList
}

type QueryDevicePropertyDataPropertyInfo struct {
	Time  goaliyun.String
	Value goaliyun.String
}

type QueryDevicePropertyDataPropertyInfoList []QueryDevicePropertyDataPropertyInfo

func (list *QueryDevicePropertyDataPropertyInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryDevicePropertyDataPropertyInfo)
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
