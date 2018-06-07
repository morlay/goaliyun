package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBInstanceAttributeRequest struct {
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstanceAttributeRequest) Invoke(client goaliyun.Client) (*DescribeDBInstanceAttributeResponse, error) {
	resp := &DescribeDBInstanceAttributeResponse{}
	err := client.Request("rds", "DescribeDBInstanceAttribute", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstanceAttributeResponse struct {
	RequestId goaliyun.String
	Items     DescribeDBInstanceAttributeDBInstanceAttributeList
}

type DescribeDBInstanceAttributeDBInstanceAttribute struct {
	DBInstanceDiskUsed                goaliyun.String
	GuardDBInstanceName               goaliyun.String
	CanTempUpgrade                    bool
	TempUpgradeTimeStart              goaliyun.String
	TempUpgradeTimeEnd                goaliyun.String
	TempUpgradeRecoveryTime           goaliyun.String
	TempUpgradeRecoveryClass          goaliyun.String
	TempUpgradeRecoveryCpu            goaliyun.Integer
	TempUpgradeRecoveryMemory         goaliyun.Integer
	TempUpgradeRecoveryMaxIOPS        goaliyun.String
	TempUpgradeRecoveryMaxConnections goaliyun.String
	InsId                             goaliyun.Integer
	DBInstanceId                      goaliyun.String
	PayType                           goaliyun.String
	DBInstanceClassType               goaliyun.String
	DBInstanceType                    goaliyun.String
	RegionId                          goaliyun.String
	ConnectionString                  goaliyun.String
	Port                              goaliyun.String
	Engine                            goaliyun.String
	EngineVersion                     goaliyun.String
	DBInstanceClass                   goaliyun.String
	DBInstanceMemory                  goaliyun.Integer
	DBInstanceStorage                 goaliyun.Integer
	VpcCloudInstanceId                goaliyun.String
	DBInstanceNetType                 goaliyun.String
	DBInstanceStatus                  goaliyun.String
	DBInstanceDescription             goaliyun.String
	LockMode                          goaliyun.String
	LockReason                        goaliyun.String
	ReadDelayTime                     goaliyun.String
	DBMaxQuantity                     goaliyun.Integer
	AccountMaxQuantity                goaliyun.Integer
	CreationTime                      goaliyun.String
	ExpireTime                        goaliyun.String
	MaintainTime                      goaliyun.String
	AvailabilityValue                 goaliyun.String
	MaxIOPS                           goaliyun.Integer
	MaxConnections                    goaliyun.Integer
	MasterInstanceId                  goaliyun.String
	DBInstanceCPU                     goaliyun.String
	IncrementSourceDBInstanceId       goaliyun.String
	GuardDBInstanceId                 goaliyun.String
	ReplicateId                       goaliyun.String
	TempDBInstanceId                  goaliyun.String
	SecurityIPList                    goaliyun.String
	ZoneId                            goaliyun.String
	InstanceNetworkType               goaliyun.String
	DBInstanceStorageType             goaliyun.String
	AdvancedFeatures                  goaliyun.String
	Category                          goaliyun.String
	AccountType                       goaliyun.String
	SupportUpgradeAccountType         goaliyun.String
	SupportCreateSuperAccount         goaliyun.String
	VpcId                             goaliyun.String
	VSwitchId                         goaliyun.String
	ConnectionMode                    goaliyun.String
	ResourceGroupId                   goaliyun.String
	ReadOnlyDBInstanceIds             DescribeDBInstanceAttributeReadOnlyDBInstanceIdList
}

type DescribeDBInstanceAttributeReadOnlyDBInstanceId struct {
	DBInstanceId goaliyun.String
}

type DescribeDBInstanceAttributeDBInstanceAttributeList []DescribeDBInstanceAttributeDBInstanceAttribute

func (list *DescribeDBInstanceAttributeDBInstanceAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceAttributeDBInstanceAttribute)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type DescribeDBInstanceAttributeReadOnlyDBInstanceIdList []DescribeDBInstanceAttributeReadOnlyDBInstanceId

func (list *DescribeDBInstanceAttributeReadOnlyDBInstanceIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceAttributeReadOnlyDBInstanceId)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
