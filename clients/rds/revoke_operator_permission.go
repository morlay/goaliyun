package rds

import (
	"github.com/morlay/goaliyun"
)

type RevokeOperatorPermissionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RevokeOperatorPermissionRequest) Invoke(client goaliyun.Client) (*RevokeOperatorPermissionResponse, error) {
	resp := &RevokeOperatorPermissionResponse{}
	err := client.Request("rds", "RevokeOperatorPermission", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RevokeOperatorPermissionResponse struct {
	RequestId goaliyun.String
}
