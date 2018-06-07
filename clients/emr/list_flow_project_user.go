package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListFlowProjectUserRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListFlowProjectUserRequest) Invoke(client goaliyun.Client) (*ListFlowProjectUserResponse, error) {
	resp := &ListFlowProjectUserResponse{}
	err := client.Request("emr", "ListFlowProjectUser", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListFlowProjectUserResponse struct {
	RequestId  goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Total      goaliyun.Integer
	Users      ListFlowProjectUserUserList
}

type ListFlowProjectUserUser struct {
	GmtCreate   goaliyun.Integer
	GmtModified goaliyun.Integer
	ProjectId   goaliyun.String
	OwnerId     goaliyun.String
	UserName    goaliyun.String
}

type ListFlowProjectUserUserList []ListFlowProjectUserUser

func (list *ListFlowProjectUserUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFlowProjectUserUser)
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
