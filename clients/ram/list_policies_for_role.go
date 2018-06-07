package ram

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListPoliciesForRoleRequest struct {
	RoleName string `position:"Query" name:"RoleName"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListPoliciesForRoleRequest) Invoke(client goaliyun.Client) (*ListPoliciesForRoleResponse, error) {
	resp := &ListPoliciesForRoleResponse{}
	err := client.Request("ram", "ListPoliciesForRole", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListPoliciesForRoleResponse struct {
	RequestId goaliyun.String
	Policies  ListPoliciesForRolePolicyList
}

type ListPoliciesForRolePolicy struct {
	PolicyName     goaliyun.String
	PolicyType     goaliyun.String
	Description    goaliyun.String
	DefaultVersion goaliyun.String
	AttachDate     goaliyun.String
}

type ListPoliciesForRolePolicyList []ListPoliciesForRolePolicy

func (list *ListPoliciesForRolePolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPoliciesForRolePolicy)
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
