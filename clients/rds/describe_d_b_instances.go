package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBInstancesRequest struct {
	ConnectionMode       string `position:"Query" name:"ConnectionMode"`
	Tag4value            string `position:"Query" name:"Tag.4.value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Tag2key              string `position:"Query" name:"Tag.2.key"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	SearchKey            string `position:"Query" name:"SearchKey"`
	Tag3key              string `position:"Query" name:"Tag.3.key"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	Tag1value            string `position:"Query" name:"Tag.1.value"`
	Engine               string `position:"Query" name:"Engine"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DBInstanceStatus     string `position:"Query" name:"DBInstanceStatus"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	Tag3value            string `position:"Query" name:"Tag.3.value"`
	ProxyId              string `position:"Query" name:"ProxyId"`
	Tag5key              string `position:"Query" name:"Tag.5.key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tag5value            string `position:"Query" name:"Tag.5.value"`
	DBInstanceType       string `position:"Query" name:"DBInstanceType"`
	Tags                 string `position:"Query" name:"Tags"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	Tag1key              string `position:"Query" name:"Tag.1.key"`
	VpcId                string `position:"Query" name:"VpcId"`
	Tag2value            string `position:"Query" name:"Tag.2.value"`
	Tag4key              string `position:"Query" name:"Tag.4.key"`
	InstanceNetworkType  string `position:"Query" name:"InstanceNetworkType"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstancesRequest) Invoke(client goaliyun.Client) (*DescribeDBInstancesResponse, error) {
	resp := &DescribeDBInstancesResponse{}
	err := client.Request("rds", "DescribeDBInstances", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstancesResponse struct {
	RequestId        goaliyun.String
	PageNumber       goaliyun.Integer
	TotalRecordCount goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	Items            DescribeDBInstancesDBInstanceList
}

type DescribeDBInstancesDBInstance struct {
	InsId                 goaliyun.Integer
	DBInstanceId          goaliyun.String
	DBInstanceDescription goaliyun.String
	PayType               goaliyun.String
	DBInstanceType        goaliyun.String
	RegionId              goaliyun.String
	ExpireTime            goaliyun.String
	DBInstanceStatus      goaliyun.String
	Engine                goaliyun.String
	DBInstanceNetType     goaliyun.String
	ConnectionMode        goaliyun.String
	LockMode              goaliyun.String
	DBInstanceClass       goaliyun.String
	InstanceNetworkType   goaliyun.String
	VpcCloudInstanceId    goaliyun.String
	LockReason            goaliyun.String
	ZoneId                goaliyun.String
	MutriORsignle         bool
	CreateTime            goaliyun.String
	EngineVersion         goaliyun.String
	GuardDBInstanceId     goaliyun.String
	TempDBInstanceId      goaliyun.String
	MasterInstanceId      goaliyun.String
	VpcId                 goaliyun.String
	VSwitchId             goaliyun.String
	ReplicateId           goaliyun.String
	ResourceGroupId       goaliyun.String
	ReadOnlyDBInstanceIds DescribeDBInstancesReadOnlyDBInstanceIdList
}

type DescribeDBInstancesReadOnlyDBInstanceId struct {
	DBInstanceId goaliyun.String
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

type DescribeDBInstancesReadOnlyDBInstanceIdList []DescribeDBInstancesReadOnlyDBInstanceId

func (list *DescribeDBInstancesReadOnlyDBInstanceIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstancesReadOnlyDBInstanceId)
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
