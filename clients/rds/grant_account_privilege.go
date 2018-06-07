package rds

import (
	"github.com/morlay/goaliyun"
)

type GrantAccountPrivilegeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AccountPrivilege     string `position:"Query" name:"AccountPrivilege"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *GrantAccountPrivilegeRequest) Invoke(client goaliyun.Client) (*GrantAccountPrivilegeResponse, error) {
	resp := &GrantAccountPrivilegeResponse{}
	err := client.Request("rds", "GrantAccountPrivilege", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GrantAccountPrivilegeResponse struct {
	RequestId goaliyun.String
}
