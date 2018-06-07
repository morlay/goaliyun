package rds

import (
	"github.com/morlay/goaliyun"
)

type ModifySQLCollectorPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	StoragePeriod        int64  `position:"Query" name:"StoragePeriod"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	SQLCollectorStatus   string `position:"Query" name:"SQLCollectorStatus"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifySQLCollectorPolicyRequest) Invoke(client goaliyun.Client) (*ModifySQLCollectorPolicyResponse, error) {
	resp := &ModifySQLCollectorPolicyResponse{}
	err := client.Request("rds", "ModifySQLCollectorPolicy", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifySQLCollectorPolicyResponse struct {
	RequestId goaliyun.String
}
