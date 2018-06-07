package domain

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
	err := client.Request("domain", "EmailVerified", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EmailVerifiedResponse struct {
	RequestId goaliyun.String
}
