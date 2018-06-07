package vpc

import (
	"github.com/morlay/goaliyun"
)

type DeleteSnatEntryRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SnatTableId          string `position:"Query" name:"SnatTableId"`
	SnatEntryId          string `position:"Query" name:"SnatEntryId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteSnatEntryRequest) Invoke(client goaliyun.Client) (*DeleteSnatEntryResponse, error) {
	resp := &DeleteSnatEntryResponse{}
	err := client.Request("vpc", "DeleteSnatEntry", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteSnatEntryResponse struct {
	RequestId goaliyun.String
}
