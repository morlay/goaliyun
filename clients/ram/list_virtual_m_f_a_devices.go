package ram

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListVirtualMFADevicesRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListVirtualMFADevicesRequest) Invoke(client goaliyun.Client) (*ListVirtualMFADevicesResponse, error) {
	resp := &ListVirtualMFADevicesResponse{}
	err := client.Request("ram", "ListVirtualMFADevices", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListVirtualMFADevicesResponse struct {
	RequestId         goaliyun.String
	VirtualMFADevices ListVirtualMFADevicesVirtualMFADeviceList
}

type ListVirtualMFADevicesVirtualMFADevice struct {
	SerialNumber goaliyun.String
	ActivateDate goaliyun.String
	User         ListVirtualMFADevicesUser
}

type ListVirtualMFADevicesUser struct {
	UserId      goaliyun.String
	UserName    goaliyun.String
	DisplayName goaliyun.String
}

type ListVirtualMFADevicesVirtualMFADeviceList []ListVirtualMFADevicesVirtualMFADevice

func (list *ListVirtualMFADevicesVirtualMFADeviceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListVirtualMFADevicesVirtualMFADevice)
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
