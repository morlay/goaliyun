package rds

import (
	"github.com/morlay/goaliyun"
)

type DescribeDBInstanceUserRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ConnectionString     string `position:"Query" name:"ConnectionString"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstanceUserRequest) Invoke(client goaliyun.Client) (*DescribeDBInstanceUserResponse, error) {
	resp := &DescribeDBInstanceUserResponse{}
	err := client.Request("rds", "DescribeDBInstanceUser", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstanceUserResponse struct {
	RequestId      goaliyun.String
	DBInstanceName goaliyun.String
	InternalDBFlag goaliyun.String
}
