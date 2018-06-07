package polardb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBInstanceAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstanceAttributeRequest) Invoke(client goaliyun.Client) (*DescribeDBInstanceAttributeResponse, error) {
	resp := &DescribeDBInstanceAttributeResponse{}
	err := client.Request("polardb", "DescribeDBInstanceAttribute", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstanceAttributeResponse struct {
	RequestId goaliyun.String
	Items     DescribeDBInstanceAttributeDBInstanceAttributeList
}

type DescribeDBInstanceAttributeDBInstanceAttribute struct {
	DBInstanceId          goaliyun.String
	DBClusterDescription  goaliyun.String
	DBClusterId           goaliyun.String
	PayType               goaliyun.String
	DBInstanceType        goaliyun.String
	RegionId              goaliyun.String
	ZoneId                goaliyun.String
	Engine                goaliyun.String
	DBType                goaliyun.String
	DBVersion             goaliyun.String
	DBInstanceClass       goaliyun.String
	DBInstanceStorage     goaliyun.Integer
	DBInstanceStatus      goaliyun.String
	DBInstanceDescription goaliyun.String
	ConnectionString      goaliyun.Integer
	Port                  goaliyun.Integer
	DBInstanceNetType     goaliyun.String
	LockMode              goaliyun.String
	LockReason            goaliyun.String
	CreationTime          goaliyun.String
	ExpireTime            goaliyun.String
	MaintainStartTime     goaliyun.String
	MaintainEndTime       goaliyun.String
	MaxConnections        goaliyun.Integer
	MaxIOPS               goaliyun.Integer
	SecurityIPList        goaliyun.String
	InstanceNetworkType   goaliyun.String
	VpcId                 goaliyun.String
	VSwitchId             goaliyun.String
	DBInstanceType1       goaliyun.String
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
