package ram

import (
	"github.com/morlay/goaliyun"
)

type UpdateLoginProfileRequest struct {
	Password              string `position:"Query" name:"Password"`
	PasswordResetRequired string `position:"Query" name:"PasswordResetRequired"`
	MFABindRequired       string `position:"Query" name:"MFABindRequired"`
	UserName              string `position:"Query" name:"UserName"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *UpdateLoginProfileRequest) Invoke(client goaliyun.Client) (*UpdateLoginProfileResponse, error) {
	resp := &UpdateLoginProfileResponse{}
	err := client.Request("ram", "UpdateLoginProfile", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateLoginProfileResponse struct {
	RequestId goaliyun.String
}
