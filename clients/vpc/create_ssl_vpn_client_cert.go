package vpc

import (
	"github.com/morlay/goaliyun"
)

type CreateSslVpnClientCertRequest struct {
	SslVpnServerId       string `position:"Query" name:"SslVpnServerId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateSslVpnClientCertRequest) Invoke(client goaliyun.Client) (*CreateSslVpnClientCertResponse, error) {
	resp := &CreateSslVpnClientCertResponse{}
	err := client.Request("vpc", "CreateSslVpnClientCert", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateSslVpnClientCertResponse struct {
	RequestId          goaliyun.String
	Name               goaliyun.String
	SslVpnClientCertId goaliyun.String
}
