package r_kvstore

import (
	"github.com/morlay/goaliyun"
)

type DescribeInstanceSSLRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstanceSSLRequest) Invoke(client goaliyun.Client) (*DescribeInstanceSSLResponse, error) {
	resp := &DescribeInstanceSSLResponse{}
	err := client.Request("r-kvstore", "DescribeInstanceSSL", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstanceSSLResponse struct {
	RequestId      goaliyun.String
	InstanceId     goaliyun.String
	SSLEnabled     goaliyun.String
	CertCommonName goaliyun.String
	SSLExpiredTime goaliyun.String
}
