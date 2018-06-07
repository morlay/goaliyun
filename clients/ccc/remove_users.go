package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type RemoveUsersRequest struct {
	InstanceId string                 `position:"Query" name:"InstanceId"`
	UserIds    *RemoveUsersUserIdList `position:"Query" type:"Repeated" name:"UserId"`
	RegionId   string                 `position:"Query" name:"RegionId"`
}

func (req *RemoveUsersRequest) Invoke(client goaliyun.Client) (*RemoveUsersResponse, error) {
	resp := &RemoveUsersResponse{}
	err := client.Request("ccc", "RemoveUsers", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveUsersResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
}

type RemoveUsersUserIdList []string

func (list *RemoveUsersUserIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
