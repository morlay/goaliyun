package vpc

import (
	"github.com/morlay/goaliyun"
)

type DeleteSslVpnClientCertRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SslVpnClientCertId   string `position:"Query" name:"SslVpnClientCertId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteSslVpnClientCertRequest) Invoke(client goaliyun.Client) (*DeleteSslVpnClientCertResponse, error) {
	resp := &DeleteSslVpnClientCertResponse{}
	err := client.Request("vpc", "DeleteSslVpnClientCert", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteSslVpnClientCertResponse struct {
	RequestId goaliyun.String
}
