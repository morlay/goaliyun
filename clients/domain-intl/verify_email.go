package domain_intl

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
	err := client.Request("domain-intl", "VerifyEmail", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type VerifyEmailResponse struct {
	RequestId goaliyun.String
}
