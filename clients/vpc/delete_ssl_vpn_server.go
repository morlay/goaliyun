package vpc

import (
	"github.com/morlay/goaliyun"
)

type DeleteSslVpnServerRequest struct {
	SslVpnServerId       string `position:"Query" name:"SslVpnServerId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteSslVpnServerRequest) Invoke(client goaliyun.Client) (*DeleteSslVpnServerResponse, error) {
	resp := &DeleteSslVpnServerResponse{}
	err := client.Request("vpc", "DeleteSslVpnServer", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteSslVpnServerResponse struct {
	RequestId goaliyun.String
}
