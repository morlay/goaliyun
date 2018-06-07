package ram

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListGroupsForUserRequest struct {
	UserName string `position:"Query" name:"UserName"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListGroupsForUserRequest) Invoke(client goaliyun.Client) (*ListGroupsForUserResponse, error) {
	resp := &ListGroupsForUserResponse{}
	err := client.Request("ram", "ListGroupsForUser", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListGroupsForUserResponse struct {
	RequestId goaliyun.String
	Groups    ListGroupsForUserGroupList
}

type ListGroupsForUserGroup struct {
	GroupName goaliyun.String
	Comments  goaliyun.String
	JoinDate  goaliyun.String
}

type ListGroupsForUserGroupList []ListGroupsForUserGroup

func (list *ListGroupsForUserGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListGroupsForUserGroup)
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
