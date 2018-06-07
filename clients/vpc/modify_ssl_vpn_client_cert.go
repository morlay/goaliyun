package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifySslVpnClientCertRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SslVpnClientCertId   string `position:"Query" name:"SslVpnClientCertId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifySslVpnClientCertRequest) Invoke(client goaliyun.Client) (*ModifySslVpnClientCertResponse, error) {
	resp := &ModifySslVpnClientCertResponse{}
	err := client.Request("vpc", "ModifySslVpnClientCert", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifySslVpnClientCertResponse struct {
	RequestId          goaliyun.String
	Name               goaliyun.String
	SslVpnClientCertId goaliyun.String
}
