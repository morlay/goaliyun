package ecs

import (
	"github.com/morlay/goaliyun"
)

type DeleteForwardEntryRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ForwardEntryId       string `position:"Query" name:"ForwardEntryId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ForwardTableId       string `position:"Query" name:"ForwardTableId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteForwardEntryRequest) Invoke(client goaliyun.Client) (*DeleteForwardEntryResponse, error) {
	resp := &DeleteForwardEntryResponse{}
	err := client.Request("ecs", "DeleteForwardEntry", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteForwardEntryResponse struct {
	RequestId goaliyun.String
}
