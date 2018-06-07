package rds

import (
	"github.com/morlay/goaliyun"
)

type RevokeAccountPrivilegeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RevokeAccountPrivilegeRequest) Invoke(client goaliyun.Client) (*RevokeAccountPrivilegeResponse, error) {
	resp := &RevokeAccountPrivilegeResponse{}
	err := client.Request("rds", "RevokeAccountPrivilege", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RevokeAccountPrivilegeResponse struct {
	RequestId goaliyun.String
}
