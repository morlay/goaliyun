package ram

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListPoliciesForGroupRequest struct {
	GroupName string `position:"Query" name:"GroupName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ListPoliciesForGroupRequest) Invoke(client goaliyun.Client) (*ListPoliciesForGroupResponse, error) {
	resp := &ListPoliciesForGroupResponse{}
	err := client.Request("ram", "ListPoliciesForGroup", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListPoliciesForGroupResponse struct {
	RequestId goaliyun.String
	Policies  ListPoliciesForGroupPolicyList
}

type ListPoliciesForGroupPolicy struct {
	PolicyName     goaliyun.String
	PolicyType     goaliyun.String
	Description    goaliyun.String
	DefaultVersion goaliyun.String
	AttachDate     goaliyun.String
}

type ListPoliciesForGroupPolicyList []ListPoliciesForGroupPolicy

func (list *ListPoliciesForGroupPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPoliciesForGroupPolicy)
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
