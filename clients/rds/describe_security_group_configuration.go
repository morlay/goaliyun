package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSecurityGroupConfigurationRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSecurityGroupConfigurationRequest) Invoke(client goaliyun.Client) (*DescribeSecurityGroupConfigurationResponse, error) {
	resp := &DescribeSecurityGroupConfigurationResponse{}
	err := client.Request("rds", "DescribeSecurityGroupConfiguration", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSecurityGroupConfigurationResponse struct {
	RequestId      goaliyun.String
	DBInstanceName goaliyun.String
	Items          DescribeSecurityGroupConfigurationEcsSecurityGroupRelationList
}

type DescribeSecurityGroupConfigurationEcsSecurityGroupRelation struct {
	RegionId        goaliyun.String
	SecurityGroupId goaliyun.String
	NetworkType     goaliyun.String
}

type DescribeSecurityGroupConfigurationEcsSecurityGroupRelationList []DescribeSecurityGroupConfigurationEcsSecurityGroupRelation

func (list *DescribeSecurityGroupConfigurationEcsSecurityGroupRelationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityGroupConfigurationEcsSecurityGroupRelation)
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
