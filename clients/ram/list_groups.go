package ram

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListGroupsRequest struct {
	Marker   string `position:"Query" name:"Marker"`
	MaxItems int64  `position:"Query" name:"MaxItems"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListGroupsRequest) Invoke(client goaliyun.Client) (*ListGroupsResponse, error) {
	resp := &ListGroupsResponse{}
	err := client.Request("ram", "ListGroups", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListGroupsResponse struct {
	RequestId   goaliyun.String
	IsTruncated bool
	Marker      goaliyun.String
	Groups      ListGroupsGroupList
}

type ListGroupsGroup struct {
	GroupName  goaliyun.String
	Comments   goaliyun.String
	CreateDate goaliyun.String
	UpdateDate goaliyun.String
}

type ListGroupsGroupList []ListGroupsGroup

func (list *ListGroupsGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListGroupsGroup)
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
