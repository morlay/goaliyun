package ram

import (
	"github.com/morlay/goaliyun"
)

type ChangePasswordRequest struct {
	OldPassword string `position:"Query" name:"OldPassword"`
	NewPassword string `position:"Query" name:"NewPassword"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *ChangePasswordRequest) Invoke(client goaliyun.Client) (*ChangePasswordResponse, error) {
	resp := &ChangePasswordResponse{}
	err := client.Request("ram", "ChangePassword", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ChangePasswordResponse struct {
	RequestId goaliyun.String
}
