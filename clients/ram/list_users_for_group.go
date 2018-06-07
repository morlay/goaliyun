package ram

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListUsersForGroupRequest struct {
	GroupName string `position:"Query" name:"GroupName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ListUsersForGroupRequest) Invoke(client goaliyun.Client) (*ListUsersForGroupResponse, error) {
	resp := &ListUsersForGroupResponse{}
	err := client.Request("ram", "ListUsersForGroup", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListUsersForGroupResponse struct {
	RequestId goaliyun.String
	Users     ListUsersForGroupUserList
}

type ListUsersForGroupUser struct {
	UserName    goaliyun.String
	DisplayName goaliyun.String
	JoinDate    goaliyun.String
}

type ListUsersForGroupUserList []ListUsersForGroupUser

func (list *ListUsersForGroupUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersForGroupUser)
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
