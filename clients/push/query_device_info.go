package push

import (
	"github.com/morlay/goaliyun"
)

type QueryDeviceInfoRequest struct {
	AppKey   int64  `position:"Query" name:"AppKey"`
	DeviceId string `position:"Query" name:"DeviceId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *QueryDeviceInfoRequest) Invoke(client goaliyun.Client) (*QueryDeviceInfoResponse, error) {
	resp := &QueryDeviceInfoResponse{}
	err := client.Request("push", "QueryDeviceInfo", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryDeviceInfoResponse struct {
	RequestId  goaliyun.String
	DeviceInfo QueryDeviceInfoDeviceInfo
}

type QueryDeviceInfoDeviceInfo struct {
	DeviceId       goaliyun.String
	DeviceType     goaliyun.String
	Account        goaliyun.String
	DeviceToken    goaliyun.String
	Tags           goaliyun.String
	Alias          goaliyun.String
	LastOnlineTime goaliyun.String
	Online         bool
	PhoneNumber    goaliyun.String
	PushEnabled    bool
}
