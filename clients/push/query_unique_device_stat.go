package push

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryUniqueDeviceStatRequest struct {
	Granularity string `position:"Query" name:"Granularity"`
	EndTime     string `position:"Query" name:"EndTime"`
	AppKey      int64  `position:"Query" name:"AppKey"`
	StartTime   string `position:"Query" name:"StartTime"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *QueryUniqueDeviceStatRequest) Invoke(client goaliyun.Client) (*QueryUniqueDeviceStatResponse, error) {
	resp := &QueryUniqueDeviceStatResponse{}
	err := client.Request("push", "QueryUniqueDeviceStat", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryUniqueDeviceStatResponse struct {
	RequestId      goaliyun.String
	AppDeviceStats QueryUniqueDeviceStatAppDeviceStatList
}

type QueryUniqueDeviceStatAppDeviceStat struct {
	Time  goaliyun.String
	Count goaliyun.Integer
}

type QueryUniqueDeviceStatAppDeviceStatList []QueryUniqueDeviceStatAppDeviceStat

func (list *QueryUniqueDeviceStatAppDeviceStatList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryUniqueDeviceStatAppDeviceStat)
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
