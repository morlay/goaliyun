package ram

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListPolicyVersionsRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListPolicyVersionsRequest) Invoke(client goaliyun.Client) (*ListPolicyVersionsResponse, error) {
	resp := &ListPolicyVersionsResponse{}
	err := client.Request("ram", "ListPolicyVersions", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListPolicyVersionsResponse struct {
	RequestId      goaliyun.String
	PolicyVersions ListPolicyVersionsPolicyVersionList
}

type ListPolicyVersionsPolicyVersion struct {
	VersionId        goaliyun.String
	IsDefaultVersion bool
	PolicyDocument   goaliyun.String
	CreateDate       goaliyun.String
}

type ListPolicyVersionsPolicyVersionList []ListPolicyVersionsPolicyVersion

func (list *ListPolicyVersionsPolicyVersionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPolicyVersionsPolicyVersion)
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
