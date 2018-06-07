package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ModifySecurityGroupConfigurationRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifySecurityGroupConfigurationRequest) Invoke(client goaliyun.Client) (*ModifySecurityGroupConfigurationResponse, error) {
	resp := &ModifySecurityGroupConfigurationResponse{}
	err := client.Request("rds", "ModifySecurityGroupConfiguration", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifySecurityGroupConfigurationResponse struct {
	RequestId      goaliyun.String
	DBInstanceName goaliyun.String
	Items          ModifySecurityGroupConfigurationEcsSecurityGroupRelationList
}

type ModifySecurityGroupConfigurationEcsSecurityGroupRelation struct {
	RegionId        goaliyun.String
	SecurityGroupId goaliyun.String
	NetworkType     goaliyun.String
}

type ModifySecurityGroupConfigurationEcsSecurityGroupRelationList []ModifySecurityGroupConfigurationEcsSecurityGroupRelation

func (list *ModifySecurityGroupConfigurationEcsSecurityGroupRelationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifySecurityGroupConfigurationEcsSecurityGroupRelation)
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
