package iot

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryDeviceServiceDataRequest struct {
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

func (req *QueryDeviceServiceDataRequest) Invoke(client goaliyun.Client) (*QueryDeviceServiceDataResponse, error) {
	resp := &QueryDeviceServiceDataResponse{}
	err := client.Request("iot", "QueryDeviceServiceData", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryDeviceServiceDataResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
	Data         QueryDeviceServiceDataData
}

type QueryDeviceServiceDataData struct {
	NextTime  goaliyun.Integer
	NextValid bool
	List      QueryDeviceServiceDataServiceInfoList
}

type QueryDeviceServiceDataServiceInfo struct {
	Time       goaliyun.String
	Identifier goaliyun.String
	Name       goaliyun.String
	InputData  goaliyun.String
	OutputData goaliyun.String
}

type QueryDeviceServiceDataServiceInfoList []QueryDeviceServiceDataServiceInfo

func (list *QueryDeviceServiceDataServiceInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryDeviceServiceDataServiceInfo)
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
