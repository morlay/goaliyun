package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBInstanceNetInfoRequest struct {
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	Flag                     string `position:"Query" name:"Flag"`
	DBInstanceNetRWSplitType string `position:"Query" name:"DBInstanceNetRWSplitType"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken              string `position:"Query" name:"ClientToken"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	DBInstanceId             string `position:"Query" name:"DBInstanceId"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
	RegionId                 string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstanceNetInfoRequest) Invoke(client goaliyun.Client) (*DescribeDBInstanceNetInfoResponse, error) {
	resp := &DescribeDBInstanceNetInfoResponse{}
	err := client.Request("rds", "DescribeDBInstanceNetInfo", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstanceNetInfoResponse struct {
	RequestId           goaliyun.String
	InstanceNetworkType goaliyun.String
	DBInstanceNetInfos  DescribeDBInstanceNetInfoDBInstanceNetInfoList
}

type DescribeDBInstanceNetInfoDBInstanceNetInfo struct {
	Upgradeable          goaliyun.String
	ExpiredTime          goaliyun.String
	ConnectionString     goaliyun.String
	IPAddress            goaliyun.String
	IPType               goaliyun.String
	Port                 goaliyun.String
	VPCId                goaliyun.String
	VSwitchId            goaliyun.String
	ConnectionStringType goaliyun.String
	MaxDelayTime         goaliyun.String
	DistributionType     goaliyun.String
	SecurityIPGroups     DescribeDBInstanceNetInfoSecurityIPGroupList
	DBInstanceWeights    DescribeDBInstanceNetInfoDBInstanceWeightList
}

type DescribeDBInstanceNetInfoSecurityIPGroup struct {
	SecurityIPGroupName goaliyun.String
	SecurityIPs         goaliyun.String
}

type DescribeDBInstanceNetInfoDBInstanceWeight struct {
	DBInstanceId   goaliyun.String
	DBInstanceType goaliyun.String
	Availability   goaliyun.String
	Weight         goaliyun.String
}

type DescribeDBInstanceNetInfoDBInstanceNetInfoList []DescribeDBInstanceNetInfoDBInstanceNetInfo

func (list *DescribeDBInstanceNetInfoDBInstanceNetInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceNetInfoDBInstanceNetInfo)
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

type DescribeDBInstanceNetInfoSecurityIPGroupList []DescribeDBInstanceNetInfoSecurityIPGroup

func (list *DescribeDBInstanceNetInfoSecurityIPGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceNetInfoSecurityIPGroup)
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

type DescribeDBInstanceNetInfoDBInstanceWeightList []DescribeDBInstanceNetInfoDBInstanceWeight

func (list *DescribeDBInstanceNetInfoDBInstanceWeightList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceNetInfoDBInstanceWeight)
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
