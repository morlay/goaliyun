package rds

import (
	"github.com/morlay/goaliyun"
)

type DescribeOperatorPermissionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeOperatorPermissionRequest) Invoke(client goaliyun.Client) (*DescribeOperatorPermissionResponse, error) {
	resp := &DescribeOperatorPermissionResponse{}
	err := client.Request("rds", "DescribeOperatorPermission", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeOperatorPermissionResponse struct {
	RequestId   goaliyun.String
	Privileges  goaliyun.String
	CreatedTime goaliyun.String
	ExpiredTime goaliyun.String
}
