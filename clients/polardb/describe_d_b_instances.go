package polardb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBInstancesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceNetworkType  string `position:"Query" name:"InstanceNetworkType"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstancesRequest) Invoke(client goaliyun.Client) (*DescribeDBInstancesResponse, error) {
	resp := &DescribeDBInstancesResponse{}
	err := client.Request("polardb", "DescribeDBInstances", "2017-08-01", req.RegionId, "", "").Do(req, resp)
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
	DBInstanceId          goaliyun.String
	DBInstanceDescription goaliyun.String
	PayType               goaliyun.String
	DBInstanceType        goaliyun.String
	DBInstanceClass       goaliyun.String
	InstanceNetworkType   goaliyun.String
	RegionId              goaliyun.String
	ZoneId                goaliyun.String
	DBClusterId           goaliyun.String
	ExpireTime            goaliyun.String
	DBInstanceStatus      goaliyun.String
	Engine                goaliyun.String
	DBType                goaliyun.String
	DBVersion             goaliyun.String
	DBInstanceType1       goaliyun.String
	LockMode              goaliyun.String
	LockReason            goaliyun.String
	CreateTime            goaliyun.String
	VpcId                 goaliyun.String
	VSwitchId             goaliyun.String
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
