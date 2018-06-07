package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AttachInstanceRamRoleRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	RamRoleName          string `position:"Query" name:"RamRoleName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AttachInstanceRamRoleRequest) Invoke(client goaliyun.Client) (*AttachInstanceRamRoleResponse, error) {
	resp := &AttachInstanceRamRoleResponse{}
	err := client.Request("ecs", "AttachInstanceRamRole", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AttachInstanceRamRoleResponse struct {
	RequestId                    goaliyun.String
	TotalCount                   goaliyun.Integer
	FailCount                    goaliyun.Integer
	RamRoleName                  goaliyun.String
	AttachInstanceRamRoleResults AttachInstanceRamRoleAttachInstanceRamRoleResultList
}

type AttachInstanceRamRoleAttachInstanceRamRoleResult struct {
	InstanceId goaliyun.String
	Success    bool
	Code       goaliyun.String
	Message    goaliyun.String
}

type AttachInstanceRamRoleAttachInstanceRamRoleResultList []AttachInstanceRamRoleAttachInstanceRamRoleResult

func (list *AttachInstanceRamRoleAttachInstanceRamRoleResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AttachInstanceRamRoleAttachInstanceRamRoleResult)
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
