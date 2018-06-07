package polardb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBClusterAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBClusterAttributeRequest) Invoke(client goaliyun.Client) (*DescribeDBClusterAttributeResponse, error) {
	resp := &DescribeDBClusterAttributeResponse{}
	err := client.Request("polardb", "DescribeDBClusterAttribute", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBClusterAttributeResponse struct {
	RequestId goaliyun.String
	Items     DescribeDBClusterAttributeDBClusterAttributeList
}

type DescribeDBClusterAttributeDBClusterAttribute struct {
	RegionId             goaliyun.String
	DBClusterNetworkType goaliyun.String
	VPCId                goaliyun.String
	VSwitchId            goaliyun.String
	PayType              goaliyun.String
	DBClusterId          goaliyun.String
	DBClusterStatus      goaliyun.String
	DBClusterDescription goaliyun.String
	Engine               goaliyun.String
	DBType               goaliyun.String
	DBVersion            goaliyun.String
	Storage              goaliyun.Integer
	ConnectionString     goaliyun.Integer
	Port                 goaliyun.Integer
	DBClusterNetType     goaliyun.String
	LockMode             goaliyun.String
	LockReason           goaliyun.String
	CreationTime         goaliyun.String
	ExpireTime           goaliyun.String
	DbInstances          DescribeDBClusterAttributeDbInstanceList
}

type DescribeDBClusterAttributeDbInstance struct {
	DBInstanceId          goaliyun.String
	DBInstanceStatus      goaliyun.String
	DBInstanceDescription goaliyun.String
	Engine                goaliyun.String
	DBType                goaliyun.String
	DBVersion             goaliyun.String
	DBInstanceStorage     goaliyun.String
	LockMode              goaliyun.String
	LockReason            goaliyun.String
	MaintainStartTime     goaliyun.String
	MaintainEndTime       goaliyun.String
	CreationTime          goaliyun.String
	DBInstanceClass       goaliyun.String
	SecurityIPList        goaliyun.String
	DBInstanceType        goaliyun.String
}

type DescribeDBClusterAttributeDBClusterAttributeList []DescribeDBClusterAttributeDBClusterAttribute

func (list *DescribeDBClusterAttributeDBClusterAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBClusterAttributeDBClusterAttribute)
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

type DescribeDBClusterAttributeDbInstanceList []DescribeDBClusterAttributeDbInstance

func (list *DescribeDBClusterAttributeDbInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBClusterAttributeDbInstance)
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
