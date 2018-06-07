package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBInstancesAsCsvRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstancesAsCsvRequest) Invoke(client goaliyun.Client) (*DescribeDBInstancesAsCsvResponse, error) {
	resp := &DescribeDBInstancesAsCsvResponse{}
	err := client.Request("rds", "DescribeDBInstancesAsCsv", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstancesAsCsvResponse struct {
	RequestId goaliyun.String
	Items     DescribeDBInstancesAsCsvDBInstanceAttributeList
}

type DescribeDBInstancesAsCsvDBInstanceAttribute struct {
	DBInstanceId                goaliyun.String
	PayType                     goaliyun.String
	DBInstanceClassType         goaliyun.String
	DBInstanceType              goaliyun.String
	RegionId                    goaliyun.String
	ConnectionString            goaliyun.String
	Port                        goaliyun.String
	Engine                      goaliyun.String
	EngineVersion               goaliyun.String
	DBInstanceClass             goaliyun.String
	DBInstanceMemory            goaliyun.Integer
	DBInstanceStorage           goaliyun.Integer
	DBInstanceNetType           goaliyun.String
	DBInstanceStatus            goaliyun.String
	DBInstanceDescription       goaliyun.String
	LockMode                    goaliyun.String
	LockReason                  goaliyun.String
	ReadDelayTime               goaliyun.String
	DBMaxQuantity               goaliyun.Integer
	AccountMaxQuantity          goaliyun.Integer
	CreationTime                goaliyun.String
	ExpireTime                  goaliyun.String
	MaintainTime                goaliyun.String
	AvailabilityValue           goaliyun.String
	MaxIOPS                     goaliyun.Integer
	MaxConnections              goaliyun.Integer
	MasterInstanceId            goaliyun.String
	DBInstanceCPU               goaliyun.String
	IncrementSourceDBInstanceId goaliyun.String
	GuardDBInstanceId           goaliyun.String
	TempDBInstanceId            goaliyun.String
	SecurityIPList              goaliyun.String
	ZoneId                      goaliyun.String
	InstanceNetworkType         goaliyun.String
	Category                    goaliyun.String
	AccountType                 goaliyun.String
	SupportUpgradeAccountType   goaliyun.String
	VpcId                       goaliyun.String
	VSwitchId                   goaliyun.String
	ConnectionMode              goaliyun.String
	Tags                        goaliyun.String
}

type DescribeDBInstancesAsCsvDBInstanceAttributeList []DescribeDBInstancesAsCsvDBInstanceAttribute

func (list *DescribeDBInstancesAsCsvDBInstanceAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesAsCsvDBInstanceAttribute)
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
