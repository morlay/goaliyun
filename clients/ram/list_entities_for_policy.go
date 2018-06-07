package ram

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListEntitiesForPolicyRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListEntitiesForPolicyRequest) Invoke(client goaliyun.Client) (*ListEntitiesForPolicyResponse, error) {
	resp := &ListEntitiesForPolicyResponse{}
	err := client.Request("ram", "ListEntitiesForPolicy", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListEntitiesForPolicyResponse struct {
	RequestId goaliyun.String
	Groups    ListEntitiesForPolicyGroupList
	Users     ListEntitiesForPolicyUserList
	Roles     ListEntitiesForPolicyRoleList
}

type ListEntitiesForPolicyGroup struct {
	GroupName  goaliyun.String
	Comments   goaliyun.String
	AttachDate goaliyun.String
}

type ListEntitiesForPolicyUser struct {
	UserId      goaliyun.String
	UserName    goaliyun.String
	DisplayName goaliyun.String
	AttachDate  goaliyun.String
}

type ListEntitiesForPolicyRole struct {
	RoleId      goaliyun.String
	RoleName    goaliyun.String
	Arn         goaliyun.String
	Description goaliyun.String
	AttachDate  goaliyun.String
}

type ListEntitiesForPolicyGroupList []ListEntitiesForPolicyGroup

func (list *ListEntitiesForPolicyGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListEntitiesForPolicyGroup)
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

type ListEntitiesForPolicyUserList []ListEntitiesForPolicyUser

func (list *ListEntitiesForPolicyUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListEntitiesForPolicyUser)
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

type ListEntitiesForPolicyRoleList []ListEntitiesForPolicyRole

func (list *ListEntitiesForPolicyRoleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListEntitiesForPolicyRole)
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
