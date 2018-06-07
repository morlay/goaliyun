package domain

import (
	"github.com/morlay/goaliyun"
)

type VerifyEmailRequest struct {
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Token        string `position:"Query" name:"Token"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *VerifyEmailRequest) Invoke(client goaliyun.Client) (*VerifyEmailResponse, error) {
	resp := &VerifyEmailResponse{}
	err := client.Request("domain", "VerifyEmail", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type VerifyEmailResponse struct {
	RequestId goaliyun.String
}
