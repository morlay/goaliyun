package rds

import (
	"github.com/morlay/goaliyun"
)

type DescribeDBInstanceSSLRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstanceSSLRequest) Invoke(client goaliyun.Client) (*DescribeDBInstanceSSLResponse, error) {
	resp := &DescribeDBInstanceSSLResponse{}
	err := client.Request("rds", "DescribeDBInstanceSSL", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstanceSSLResponse struct {
	RequestId           goaliyun.String
	ConnectionString    goaliyun.String
	SSLExpireTime       goaliyun.String
	RequireUpdate       goaliyun.String
	RequireUpdateReason goaliyun.String
}
