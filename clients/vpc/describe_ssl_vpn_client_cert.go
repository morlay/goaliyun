package vpc

import (
	"github.com/morlay/goaliyun"
)

type DescribeSslVpnClientCertRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SslVpnClientCertId   string `position:"Query" name:"SslVpnClientCertId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSslVpnClientCertRequest) Invoke(client goaliyun.Client) (*DescribeSslVpnClientCertResponse, error) {
	resp := &DescribeSslVpnClientCertResponse{}
	err := client.Request("vpc", "DescribeSslVpnClientCert", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSslVpnClientCertResponse struct {
	RequestId          goaliyun.String
	RegionId           goaliyun.String
	SslVpnClientCertId goaliyun.String
	Name               goaliyun.String
	SslVpnServerId     goaliyun.String
	CaCert             goaliyun.String
	ClientCert         goaliyun.String
	ClientKey          goaliyun.String
	ClientConfig       goaliyun.String
	CreateTime         goaliyun.Integer
	EndTime            goaliyun.Integer
	Status             goaliyun.String
}
