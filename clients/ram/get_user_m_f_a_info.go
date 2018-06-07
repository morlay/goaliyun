package ram

import (
	"github.com/morlay/goaliyun"
)

type GetUserMFAInfoRequest struct {
	UserName string `position:"Query" name:"UserName"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetUserMFAInfoRequest) Invoke(client goaliyun.Client) (*GetUserMFAInfoResponse, error) {
	resp := &GetUserMFAInfoResponse{}
	err := client.Request("ram", "GetUserMFAInfo", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetUserMFAInfoResponse struct {
	RequestId goaliyun.String
	MFADevice GetUserMFAInfoMFADevice
}

type GetUserMFAInfoMFADevice struct {
	SerialNumber goaliyun.String
}
