package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DetachInstanceRamRoleRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	RamRoleName          string `position:"Query" name:"RamRoleName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DetachInstanceRamRoleRequest) Invoke(client goaliyun.Client) (*DetachInstanceRamRoleResponse, error) {
	resp := &DetachInstanceRamRoleResponse{}
	err := client.Request("ecs", "DetachInstanceRamRole", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DetachInstanceRamRoleResponse struct {
	RequestId                    goaliyun.String
	TotalCount                   goaliyun.Integer
	FailCount                    goaliyun.Integer
	RamRoleName                  goaliyun.String
	DetachInstanceRamRoleResults DetachInstanceRamRoleDetachInstanceRamRoleResultList
}

type DetachInstanceRamRoleDetachInstanceRamRoleResult struct {
	InstanceId          goaliyun.String
	Success             bool
	Code                goaliyun.String
	Message             goaliyun.String
	InstanceRamRoleSets DetachInstanceRamRoleInstanceRamRoleSetList
}

type DetachInstanceRamRoleInstanceRamRoleSet struct {
	InstanceId  goaliyun.String
	RamRoleName goaliyun.String
}

type DetachInstanceRamRoleDetachInstanceRamRoleResultList []DetachInstanceRamRoleDetachInstanceRamRoleResult

func (list *DetachInstanceRamRoleDetachInstanceRamRoleResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetachInstanceRamRoleDetachInstanceRamRoleResult)
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

type DetachInstanceRamRoleInstanceRamRoleSetList []DetachInstanceRamRoleInstanceRamRoleSet

func (list *DetachInstanceRamRoleInstanceRamRoleSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetachInstanceRamRoleInstanceRamRoleSet)
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
