package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListUsersRequest struct {
	PageSize   int64  `position:"Query" name:"PageSize"`
	ClusterId  string `position:"Query" name:"ClusterId"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListUsersRequest) Invoke(client goaliyun.Client) (*ListUsersResponse, error) {
	resp := &ListUsersResponse{}
	err := client.Request("ehpc", "ListUsers", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListUsersResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Users      ListUsersUserInfoList
}

type ListUsersUserInfo struct {
	Name    goaliyun.String
	Group   goaliyun.String
	AddTime goaliyun.String
}

type ListUsersUserInfoList []ListUsersUserInfo

func (list *ListUsersUserInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersUserInfo)
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
