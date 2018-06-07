package iot

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryDevicePropertyStatusRequest struct {
	IotId      string `position:"Query" name:"IotId"`
	DeviceName string `position:"Query" name:"DeviceName"`
	ProductKey string `position:"Query" name:"ProductKey"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *QueryDevicePropertyStatusRequest) Invoke(client goaliyun.Client) (*QueryDevicePropertyStatusResponse, error) {
	resp := &QueryDevicePropertyStatusResponse{}
	err := client.Request("iot", "QueryDevicePropertyStatus", "2018-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryDevicePropertyStatusResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorMessage goaliyun.String
	Data         QueryDevicePropertyStatusData
}

type QueryDevicePropertyStatusData struct {
	List QueryDevicePropertyStatusPropertyStatusInfoList
}

type QueryDevicePropertyStatusPropertyStatusInfo struct {
	Unit       goaliyun.String
	Identifier goaliyun.String
	DataType   goaliyun.String
	Time       goaliyun.String
	Value      goaliyun.String
	Name       goaliyun.String
}

type QueryDevicePropertyStatusPropertyStatusInfoList []QueryDevicePropertyStatusPropertyStatusInfo

func (list *QueryDevicePropertyStatusPropertyStatusInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryDevicePropertyStatusPropertyStatusInfo)
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
