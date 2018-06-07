package ram

import (
	"github.com/morlay/goaliyun"
)

type CreateLoginProfileRequest struct {
	Password              string `position:"Query" name:"Password"`
	PasswordResetRequired string `position:"Query" name:"PasswordResetRequired"`
	MFABindRequired       string `position:"Query" name:"MFABindRequired"`
	UserName              string `position:"Query" name:"UserName"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *CreateLoginProfileRequest) Invoke(client goaliyun.Client) (*CreateLoginProfileResponse, error) {
	resp := &CreateLoginProfileResponse{}
	err := client.Request("ram", "CreateLoginProfile", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateLoginProfileResponse struct {
	RequestId    goaliyun.String
	LoginProfile CreateLoginProfileLoginProfile
}

type CreateLoginProfileLoginProfile struct {
	UserName              goaliyun.String
	PasswordResetRequired bool
	MFABindRequired       bool
	CreateDate            goaliyun.String
}
