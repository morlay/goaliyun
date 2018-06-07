package vpc

import (
	"github.com/morlay/goaliyun"
)

type TerminatePhysicalConnectionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	PhysicalConnectionId string `position:"Query" name:"PhysicalConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *TerminatePhysicalConnectionRequest) Invoke(client goaliyun.Client) (*TerminatePhysicalConnectionResponse, error) {
	resp := &TerminatePhysicalConnectionResponse{}
	err := client.Request("vpc", "TerminatePhysicalConnection", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type TerminatePhysicalConnectionResponse struct {
	RequestId goaliyun.String
}
