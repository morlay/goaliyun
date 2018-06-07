package dm

import (
	"github.com/morlay/goaliyun"
)

type SingleSendMailRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	AddressType          int64  `position:"Query" name:"AddressType"`
	TagName              string `position:"Query" name:"TagName"`
	ReplyToAddress       string `position:"Query" name:"ReplyToAddress"`
	ToAddress            string `position:"Query" name:"ToAddress"`
	Subject              string `position:"Query" name:"Subject"`
	HtmlBody             string `position:"Query" name:"HtmlBody"`
	TextBody             string `position:"Query" name:"TextBody"`
	FromAlias            string `position:"Query" name:"FromAlias"`
	ReplyAddress         string `position:"Query" name:"ReplyAddress"`
	ReplyAddressAlias    string `position:"Query" name:"ReplyAddressAlias"`
	ClickTrace           string `position:"Query" name:"ClickTrace"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SingleSendMailRequest) Invoke(client goaliyun.Client) (*SingleSendMailResponse, error) {
	resp := &SingleSendMailResponse{}
	err := client.Request("dm", "SingleSendMail", "2015-11-23", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SingleSendMailResponse struct {
	RequestId goaliyun.String
	EnvId     goaliyun.String
}
