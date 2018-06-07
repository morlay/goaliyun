package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSslVpnClientCertsRequest struct {
	SslVpnServerId       string `position:"Query" name:"SslVpnServerId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SslVpnClientCertId   string `position:"Query" name:"SslVpnClientCertId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSslVpnClientCertsRequest) Invoke(client goaliyun.Client) (*DescribeSslVpnClientCertsResponse, error) {
	resp := &DescribeSslVpnClientCertsResponse{}
	err := client.Request("vpc", "DescribeSslVpnClientCerts", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSslVpnClientCertsResponse struct {
	RequestId            goaliyun.String
	TotalCount           goaliyun.Integer
	PageNumber           goaliyun.Integer
	PageSize             goaliyun.Integer
	SslVpnClientCertKeys DescribeSslVpnClientCertsSslVpnClientCertKeyList
}

type DescribeSslVpnClientCertsSslVpnClientCertKey struct {
	RegionId           goaliyun.String
	SslVpnClientCertId goaliyun.String
	Name               goaliyun.String
	SslVpnServerId     goaliyun.String
	CreateTime         goaliyun.Integer
	EndTime            goaliyun.Integer
	Status             goaliyun.String
}

type DescribeSslVpnClientCertsSslVpnClientCertKeyList []DescribeSslVpnClientCertsSslVpnClientCertKey

func (list *DescribeSslVpnClientCertsSslVpnClientCertKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSslVpnClientCertsSslVpnClientCertKey)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
