package ram

import (
	"github.com/morlay/goaliyun"
)

type GetLoginProfileRequest struct {
	UserName string `position:"Query" name:"UserName"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetLoginProfileRequest) Invoke(client goaliyun.Client) (*GetLoginProfileResponse, error) {
	resp := &GetLoginProfileResponse{}
	err := client.Request("ram", "GetLoginProfile", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetLoginProfileResponse struct {
	RequestId    goaliyun.String
	LoginProfile GetLoginProfileLoginProfile
}

type GetLoginProfileLoginProfile struct {
	UserName              goaliyun.String
	PasswordResetRequired bool
	MFABindRequired       bool
	CreateDate            goaliyun.String
}
