package iot

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryDeviceEventDataRequest struct {
	Asc        int64  `position:"Query" name:"Asc"`
	Identifier string `position:"Query" name:"Identifier"`
	IotId      string `position:"Query" name:"IotId"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	EndTime    int64  `position:"Query" name:"EndTime"`
	EventType  string `position:"Query" name:"EventType"`
	DeviceName string `position:"Query" name:"DeviceName"`
	StartTime  int64  `position:"Query" name:"StartTime"`
	ProductKey string `position:"Query" name:"ProductKey"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *QueryDeviceEventDataRequest) Invoke(client goaliyun.Client) (*QueryDeviceEventDataResponse, error) {
	resp := &QueryDeviceEventDataResponse{}
	err := client.Request("iot", "QueryDeviceEventData", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryDeviceEventDataResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
	Data         QueryDeviceEventDataData
}

type QueryDeviceEventDataData struct {
	NextTime  goaliyun.Integer
	NextValid bool
	List      QueryDeviceEventDataEventInfoList
}

type QueryDeviceEventDataEventInfo struct {
	Time       goaliyun.String
	Identifier goaliyun.String
	Name       goaliyun.String
	EventType  goaliyun.String
	OutputData goaliyun.String
}

type QueryDeviceEventDataEventInfoList []QueryDeviceEventDataEventInfo

func (list *QueryDeviceEventDataEventInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryDeviceEventDataEventInfo)
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
