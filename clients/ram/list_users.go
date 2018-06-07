package ram

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListUsersRequest struct {
	Marker   string `position:"Query" name:"Marker"`
	MaxItems int64  `position:"Query" name:"MaxItems"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListUsersRequest) Invoke(client goaliyun.Client) (*ListUsersResponse, error) {
	resp := &ListUsersResponse{}
	err := client.Request("ram", "ListUsers", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListUsersResponse struct {
	RequestId   goaliyun.String
	IsTruncated bool
	Marker      goaliyun.String
	Users       ListUsersUserList
}

type ListUsersUser struct {
	UserId      goaliyun.String
	UserName    goaliyun.String
	DisplayName goaliyun.String
	MobilePhone goaliyun.String
	Email       goaliyun.String
	Comments    goaliyun.String
	CreateDate  goaliyun.String
	UpdateDate  goaliyun.String
}

type ListUsersUserList []ListUsersUser

func (list *ListUsersUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersUser)
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
