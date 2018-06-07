package ram

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListPoliciesRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	Marker     string `position:"Query" name:"Marker"`
	MaxItems   int64  `position:"Query" name:"MaxItems"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListPoliciesRequest) Invoke(client goaliyun.Client) (*ListPoliciesResponse, error) {
	resp := &ListPoliciesResponse{}
	err := client.Request("ram", "ListPolicies", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListPoliciesResponse struct {
	RequestId   goaliyun.String
	IsTruncated bool
	Marker      goaliyun.String
	Policies    ListPoliciesPolicyList
}

type ListPoliciesPolicy struct {
	PolicyName      goaliyun.String
	PolicyType      goaliyun.String
	Description     goaliyun.String
	DefaultVersion  goaliyun.String
	CreateDate      goaliyun.String
	UpdateDate      goaliyun.String
	AttachmentCount goaliyun.Integer
}

type ListPoliciesPolicyList []ListPoliciesPolicy

func (list *ListPoliciesPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPoliciesPolicy)
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
