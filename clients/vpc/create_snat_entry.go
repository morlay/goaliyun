package vpc

import (
	"github.com/morlay/goaliyun"
)

type CreateSnatEntryRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SourceVSwitchId      string `position:"Query" name:"SourceVSwitchId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SourceCIDR           string `position:"Query" name:"SourceCIDR"`
	SnatTableId          string `position:"Query" name:"SnatTableId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SnatIp               string `position:"Query" name:"SnatIp"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateSnatEntryRequest) Invoke(client goaliyun.Client) (*CreateSnatEntryResponse, error) {
	resp := &CreateSnatEntryResponse{}
	err := client.Request("vpc", "CreateSnatEntry", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateSnatEntryResponse struct {
	RequestId   goaliyun.String
	SnatEntryId goaliyun.String
}
