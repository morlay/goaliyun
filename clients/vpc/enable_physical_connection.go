package vpc

import (
	"github.com/morlay/goaliyun"
)

type EnablePhysicalConnectionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	PhysicalConnectionId string `position:"Query" name:"PhysicalConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *EnablePhysicalConnectionRequest) Invoke(client goaliyun.Client) (*EnablePhysicalConnectionResponse, error) {
	resp := &EnablePhysicalConnectionResponse{}
	err := client.Request("vpc", "EnablePhysicalConnection", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EnablePhysicalConnectionResponse struct {
	RequestId goaliyun.String
}
