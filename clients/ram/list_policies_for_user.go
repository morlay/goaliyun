package ram

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListPoliciesForUserRequest struct {
	UserName string `position:"Query" name:"UserName"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListPoliciesForUserRequest) Invoke(client goaliyun.Client) (*ListPoliciesForUserResponse, error) {
	resp := &ListPoliciesForUserResponse{}
	err := client.Request("ram", "ListPoliciesForUser", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListPoliciesForUserResponse struct {
	RequestId goaliyun.String
	Policies  ListPoliciesForUserPolicyList
}

type ListPoliciesForUserPolicy struct {
	PolicyName     goaliyun.String
	PolicyType     goaliyun.String
	Description    goaliyun.String
	DefaultVersion goaliyun.String
	AttachDate     goaliyun.String
}

type ListPoliciesForUserPolicyList []ListPoliciesForUserPolicy

func (list *ListPoliciesForUserPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPoliciesForUserPolicy)
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
