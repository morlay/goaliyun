package ram

import (
	"github.com/morlay/goaliyun"
)

type DeleteVirtualMFADeviceRequest struct {
	SerialNumber string `position:"Query" name:"SerialNumber"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DeleteVirtualMFADeviceRequest) Invoke(client goaliyun.Client) (*DeleteVirtualMFADeviceResponse, error) {
	resp := &DeleteVirtualMFADeviceResponse{}
	err := client.Request("ram", "DeleteVirtualMFADevice", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteVirtualMFADeviceResponse struct {
	RequestId goaliyun.String
}
