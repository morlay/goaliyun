package ecs

import (
	"github.com/morlay/goaliyun"
)

type CancelPhysicalConnectionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	PhysicalConnectionId string `position:"Query" name:"PhysicalConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CancelPhysicalConnectionRequest) Invoke(client goaliyun.Client) (*CancelPhysicalConnectionResponse, error) {
	resp := &CancelPhysicalConnectionResponse{}
	err := client.Request("ecs", "CancelPhysicalConnection", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CancelPhysicalConnectionResponse struct {
	RequestId goaliyun.String
}
