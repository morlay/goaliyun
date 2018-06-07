package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyForwardEntryRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	IpProtocol           string `position:"Query" name:"IpProtocol"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ForwardTableId       string `position:"Query" name:"ForwardTableId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InternalIp           string `position:"Query" name:"InternalIp"`
	ForwardEntryId       string `position:"Query" name:"ForwardEntryId"`
	InternalPort         string `position:"Query" name:"InternalPort"`
	ExternalIp           string `position:"Query" name:"ExternalIp"`
	ExternalPort         string `position:"Query" name:"ExternalPort"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyForwardEntryRequest) Invoke(client goaliyun.Client) (*ModifyForwardEntryResponse, error) {
	resp := &ModifyForwardEntryResponse{}
	err := client.Request("vpc", "ModifyForwardEntry", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyForwardEntryResponse struct {
	RequestId goaliyun.String
}
