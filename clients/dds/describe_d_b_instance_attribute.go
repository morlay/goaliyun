package dds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBInstanceAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Engine               string `position:"Query" name:"Engine"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstanceAttributeRequest) Invoke(client goaliyun.Client) (*DescribeDBInstanceAttributeResponse, error) {
	resp := &DescribeDBInstanceAttributeResponse{}
	err := client.Request("dds", "DescribeDBInstanceAttribute", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstanceAttributeResponse struct {
	RequestId   goaliyun.String
	DBInstances DescribeDBInstanceAttributeDBInstanceList
}

type DescribeDBInstanceAttributeDBInstance struct {
	DBInstanceId          goaliyun.String
	DBInstanceDescription goaliyun.String
	RegionId              goaliyun.String
	ZoneId                goaliyun.String
	Engine                goaliyun.String
	EngineVersion         goaliyun.String
	StorageEngine         goaliyun.String
	DBInstanceClass       goaliyun.String
	DBInstanceStorage     goaliyun.Integer
	DBInstanceStatus      goaliyun.String
	LockMode              goaliyun.String
	ChargeType            goaliyun.String
	CreationTime          goaliyun.String
	ReplicaSetName        goaliyun.String
	NetworkType           goaliyun.String
	ExpireTime            goaliyun.String
	MaintainStartTime     goaliyun.String
	MaintainEndTime       goaliyun.String
	DBInstanceType        goaliyun.String
	LastDowngradeTime     goaliyun.Integer
	MongosList            DescribeDBInstanceAttributeMongosAttributeList
	ShardList             DescribeDBInstanceAttributeShardAttributeList
}

type DescribeDBInstanceAttributeMongosAttribute struct {
	NodeId          goaliyun.String
	NodeDescription goaliyun.String
	NodeClass       goaliyun.String
	ConnectSting    goaliyun.String
	Port            goaliyun.Integer
}

type DescribeDBInstanceAttributeShardAttribute struct {
	NodeId          goaliyun.String
	NodeDescription goaliyun.String
	NodeClass       goaliyun.String
	NodeStorage     goaliyun.Integer
}

type DescribeDBInstanceAttributeDBInstanceList []DescribeDBInstanceAttributeDBInstance

func (list *DescribeDBInstanceAttributeDBInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceAttributeDBInstance)
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

type DescribeDBInstanceAttributeMongosAttributeList []DescribeDBInstanceAttributeMongosAttribute

func (list *DescribeDBInstanceAttributeMongosAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceAttributeMongosAttribute)
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

type DescribeDBInstanceAttributeShardAttributeList []DescribeDBInstanceAttributeShardAttribute

func (list *DescribeDBInstanceAttributeShardAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceAttributeShardAttribute)
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
