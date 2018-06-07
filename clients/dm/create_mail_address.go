package dm

import (
	"github.com/morlay/goaliyun"
)

type CreateMailAddressRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	ReplyAddress         string `position:"Query" name:"ReplyAddress"`
	Sendtype             string `position:"Query" name:"Sendtype"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateMailAddressRequest) Invoke(client goaliyun.Client) (*CreateMailAddressResponse, error) {
	resp := &CreateMailAddressResponse{}
	err := client.Request("dm", "CreateMailAddress", "2015-11-23", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateMailAddressResponse struct {
	RequestId goaliyun.String
}
