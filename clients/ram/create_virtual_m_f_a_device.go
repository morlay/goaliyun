package ram

import (
	"github.com/morlay/goaliyun"
)

type CreateVirtualMFADeviceRequest struct {
	VirtualMFADeviceName string `position:"Query" name:"VirtualMFADeviceName"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateVirtualMFADeviceRequest) Invoke(client goaliyun.Client) (*CreateVirtualMFADeviceResponse, error) {
	resp := &CreateVirtualMFADeviceResponse{}
	err := client.Request("ram", "CreateVirtualMFADevice", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateVirtualMFADeviceResponse struct {
	RequestId        goaliyun.String
	VirtualMFADevice CreateVirtualMFADeviceVirtualMFADevice
}

type CreateVirtualMFADeviceVirtualMFADevice struct {
	SerialNumber     goaliyun.String
	Base32StringSeed goaliyun.String
	QRCodePNG        goaliyun.String
}
