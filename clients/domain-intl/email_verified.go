package domain_intl

import (
	"github.com/morlay/goaliyun"
)

type EmailVerifiedRequest struct {
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Email        string `position:"Query" name:"Email"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *EmailVerifiedRequest) Invoke(client goaliyun.Client) (*EmailVerifiedResponse, error) {
	resp := &EmailVerifiedResponse{}
	err := client.Request("domain-intl", "EmailVerified", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EmailVerifiedResponse struct {
	RequestId goaliyun.String
}
