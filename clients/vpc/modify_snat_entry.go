package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifySnatEntryRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SnatTableId          string `position:"Query" name:"SnatTableId"`
	SnatEntryId          string `position:"Query" name:"SnatEntryId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SnatIp               string `position:"Query" name:"SnatIp"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifySnatEntryRequest) Invoke(client goaliyun.Client) (*ModifySnatEntryResponse, error) {
	resp := &ModifySnatEntryResponse{}
	err := client.Request("vpc", "ModifySnatEntry", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifySnatEntryResponse struct {
	RequestId goaliyun.String
}
