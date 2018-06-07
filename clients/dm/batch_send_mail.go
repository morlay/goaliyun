package dm

import (
	"github.com/morlay/goaliyun"
)

type BatchSendMailRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	TemplateName         string `position:"Query" name:"TemplateName"`
	AccountName          string `position:"Query" name:"AccountName"`
	ReceiversName        string `position:"Query" name:"ReceiversName"`
	AddressType          int64  `position:"Query" name:"AddressType"`
	TagName              string `position:"Query" name:"TagName"`
	ReplyAddress         string `position:"Query" name:"ReplyAddress"`
	ReplyAddressAlias    string `position:"Query" name:"ReplyAddressAlias"`
	ClickTrace           string `position:"Query" name:"ClickTrace"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *BatchSendMailRequest) Invoke(client goaliyun.Client) (*BatchSendMailResponse, error) {
	resp := &BatchSendMailResponse{}
	err := client.Request("dm", "BatchSendMail", "2015-11-23", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type BatchSendMailResponse struct {
	RequestId goaliyun.String
	EnvId     goaliyun.String
}
