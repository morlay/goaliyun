package vpc

import (
	"github.com/morlay/goaliyun"
)

type DeleteVpnConnectionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	VpnConnectionId      string `position:"Query" name:"VpnConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteVpnConnectionRequest) Invoke(client goaliyun.Client) (*DeleteVpnConnectionResponse, error) {
	resp := &DeleteVpnConnectionResponse{}
	err := client.Request("vpc", "DeleteVpnConnection", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteVpnConnectionResponse struct {
	RequestId goaliyun.String
}
