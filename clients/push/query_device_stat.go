package push

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryDeviceStatRequest struct {
	EndTime    string `position:"Query" name:"EndTime"`
	AppKey     int64  `position:"Query" name:"AppKey"`
	StartTime  string `position:"Query" name:"StartTime"`
	DeviceType string `position:"Query" name:"DeviceType"`
	QueryType  string `position:"Query" name:"QueryType"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *QueryDeviceStatRequest) Invoke(client goaliyun.Client) (*QueryDeviceStatResponse, error) {
	resp := &QueryDeviceStatResponse{}
	err := client.Request("push", "QueryDeviceStat", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryDeviceStatResponse struct {
	RequestId      goaliyun.String
	AppDeviceStats QueryDeviceStatAppDeviceStatList
}

type QueryDeviceStatAppDeviceStat struct {
	Time       goaliyun.String
	Count      goaliyun.Integer
	DeviceType goaliyun.String
}

type QueryDeviceStatAppDeviceStatList []QueryDeviceStatAppDeviceStat

func (list *QueryDeviceStatAppDeviceStatList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryDeviceStatAppDeviceStat)
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
