package push

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CheckDevicesRequest struct {
	DeviceIds string `position:"Query" name:"DeviceIds"`
	AppKey    int64  `position:"Query" name:"AppKey"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *CheckDevicesRequest) Invoke(client goaliyun.Client) (*CheckDevicesResponse, error) {
	resp := &CheckDevicesResponse{}
	err := client.Request("push", "CheckDevices", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CheckDevicesResponse struct {
	RequestId        goaliyun.String
	DeviceCheckInfos CheckDevicesDeviceCheckInfoList
}

type CheckDevicesDeviceCheckInfo struct {
	DeviceId  goaliyun.String
	Available bool
}

type CheckDevicesDeviceCheckInfoList []CheckDevicesDeviceCheckInfo

func (list *CheckDevicesDeviceCheckInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CheckDevicesDeviceCheckInfo)
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
