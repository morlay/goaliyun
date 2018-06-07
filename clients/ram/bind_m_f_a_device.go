package ram

import (
	"github.com/morlay/goaliyun"
)

type BindMFADeviceRequest struct {
	SerialNumber        string `position:"Query" name:"SerialNumber"`
	AuthenticationCode2 string `position:"Query" name:"AuthenticationCode.2"`
	AuthenticationCode1 string `position:"Query" name:"AuthenticationCode.1"`
	UserName            string `position:"Query" name:"UserName"`
	RegionId            string `position:"Query" name:"RegionId"`
}

func (req *BindMFADeviceRequest) Invoke(client goaliyun.Client) (*BindMFADeviceResponse, error) {
	resp := &BindMFADeviceResponse{}
	err := client.Request("ram", "BindMFADevice", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type BindMFADeviceResponse struct {
	RequestId goaliyun.String
}
