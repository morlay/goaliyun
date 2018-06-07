package ecs

import (
	"github.com/morlay/goaliyun"
)

type CreateForwardEntryRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	IpProtocol           string `position:"Query" name:"IpProtocol"`
	InternalPort         string `position:"Query" name:"InternalPort"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ForwardTableId       string `position:"Query" name:"ForwardTableId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ExternalIp           string `position:"Query" name:"ExternalIp"`
	ExternalPort         string `position:"Query" name:"ExternalPort"`
	InternalIp           string `position:"Query" name:"InternalIp"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateForwardEntryRequest) Invoke(client goaliyun.Client) (*CreateForwardEntryResponse, error) {
	resp := &CreateForwardEntryResponse{}
	err := client.Request("ecs", "CreateForwardEntry", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateForwardEntryResponse struct {
	RequestId      goaliyun.String
	ForwardEntryId goaliyun.String
}
