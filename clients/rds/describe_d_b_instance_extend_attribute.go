package rds

import (
	"github.com/morlay/goaliyun"
)

type DescribeDBInstanceExtendAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstanceExtendAttributeRequest) Invoke(client goaliyun.Client) (*DescribeDBInstanceExtendAttributeResponse, error) {
	resp := &DescribeDBInstanceExtendAttributeResponse{}
	err := client.Request("rds", "DescribeDBInstanceExtendAttribute", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstanceExtendAttributeResponse struct {
	RequestId                         goaliyun.String
	CanTempUpgrade                    bool
	TempUpgradeTimeStart              goaliyun.String
	TempUpgradeTimeEnd                goaliyun.String
	TempUpgradeRecoveryTime           goaliyun.String
	TempUpgradeRecoveryClass          goaliyun.String
	TempUpgradeRecoveryCpu            goaliyun.Integer
	TempUpgradeRecoveryMemory         goaliyun.Integer
	TempUpgradeRecoveryMaxIOPS        goaliyun.String
	TempUpgradeRecoveryMaxConnections goaliyun.String
}
