package rds

import (
	"github.com/morlay/goaliyun"
)

type GrantOperatorPermissionRequest struct {
	Privileges           string `position:"Query" name:"Privileges"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ExpiredTime          string `position:"Query" name:"ExpiredTime"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GrantOperatorPermissionRequest) Invoke(client goaliyun.Client) (*GrantOperatorPermissionResponse, error) {
	resp := &GrantOperatorPermissionResponse{}
	err := client.Request("rds", "GrantOperatorPermission", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GrantOperatorPermissionResponse struct {
	RequestId goaliyun.String
}
