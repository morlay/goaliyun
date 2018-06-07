package dds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBInstancesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBInstanceIds        string `position:"Query" name:"DBInstanceIds"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	DBInstanceType       string `position:"Query" name:"DBInstanceType"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	Engine               string `position:"Query" name:"Engine"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstancesRequest) Invoke(client goaliyun.Client) (*DescribeDBInstancesResponse, error) {
	resp := &DescribeDBInstancesResponse{}
	err := client.Request("dds", "DescribeDBInstances", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstancesResponse struct {
	RequestId   goaliyun.String
	PageNumber  goaliyun.Integer
	PageSize    goaliyun.Integer
	TotalCount  goaliyun.Integer
	DBInstances DescribeDBInstancesDBInstanceList
}

type DescribeDBInstancesDBInstance struct {
	DBInstanceId          goaliyun.String
	DBInstanceDescription goaliyun.String
	RegionId              goaliyun.String
	ZoneId                goaliyun.String
	Engine                goaliyun.String
	EngineVersion         goaliyun.String
	DBInstanceClass       goaliyun.String
	DBInstanceStorage     goaliyun.Integer
	DBInstanceStatus      goaliyun.String
	LockMode              goaliyun.String
	ChargeType            goaliyun.String
	NetworkType           goaliyun.String
	CreationTime          goaliyun.String
	ExpireTime            goaliyun.String
	DBInstanceType        goaliyun.String
	LastDowngradeTime     goaliyun.Integer
	MongosList            DescribeDBInstancesMongosAttributeList
	ShardList             DescribeDBInstancesShardAttributeList
}

type DescribeDBInstancesMongosAttribute struct {
	NodeId          goaliyun.String
	NodeDescription goaliyun.String
	NodeClass       goaliyun.String
	ConnectSting    goaliyun.String
	Port            goaliyun.Integer
}

type DescribeDBInstancesShardAttribute struct {
	NodeId          goaliyun.String
	NodeDescription goaliyun.String
	NodeClass       goaliyun.String
	NodeStorage     goaliyun.Integer
}

type DescribeDBInstancesDBInstanceList []DescribeDBInstancesDBInstance

func (list *DescribeDBInstancesDBInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesDBInstance)
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

type DescribeDBInstancesMongosAttributeList []DescribeDBInstancesMongosAttribute

func (list *DescribeDBInstancesMongosAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesMongosAttribute)
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

type DescribeDBInstancesShardAttributeList []DescribeDBInstancesShardAttribute

func (list *DescribeDBInstancesShardAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesShardAttribute)
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
