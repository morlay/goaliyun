package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListUserIdInLastTimeForAdminRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	StartTime       int64  `position:"Query" name:"StartTime"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListUserIdInLastTimeForAdminRequest) Invoke(client goaliyun.Client) (*ListUserIdInLastTimeForAdminResponse, error) {
	resp := &ListUserIdInLastTimeForAdminResponse{}
	err := client.Request("emr", "ListUserIdInLastTimeForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListUserIdInLastTimeForAdminResponse struct {
	RequestId  goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalCount goaliyun.Integer
	UserIds    ListUserIdInLastTimeForAdminUserIdList
}

type ListUserIdInLastTimeForAdminUserIdList []goaliyun.String

func (list *ListUserIdInLastTimeForAdminUserIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
