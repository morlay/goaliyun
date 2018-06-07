package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListMyGroupsRequest struct {
	SelectContactGroups string `position:"Query" name:"SelectContactGroups"`
	InstanceId          string `position:"Query" name:"InstanceId"`
	PageSize            int64  `position:"Query" name:"PageSize"`
	Type                string `position:"Query" name:"Type"`
	Keyword             string `position:"Query" name:"Keyword"`
	GroupName           string `position:"Query" name:"GroupName"`
	PageNumber          int64  `position:"Query" name:"PageNumber"`
	BindUrls            string `position:"Query" name:"BindUrls"`
	RegionId            string `position:"Query" name:"RegionId"`
}

func (req *ListMyGroupsRequest) Invoke(client goaliyun.Client) (*ListMyGroupsResponse, error) {
	resp := &ListMyGroupsResponse{}
	err := client.Request("cms", "ListMyGroups", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListMyGroupsResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
	PageNumber   goaliyun.Integer
	PageSize     goaliyun.Integer
	Total        goaliyun.Integer
	Resources    ListMyGroupsResourceList
}

type ListMyGroupsResource struct {
	GroupId       goaliyun.Integer
	GroupName     goaliyun.String
	ServiceId     goaliyun.String
	BindUrls      goaliyun.String
	Type          goaliyun.String
	GmtModified   goaliyun.Integer
	GmtCreate     goaliyun.Integer
	ContactGroups ListMyGroupsContactGroupList
}

type ListMyGroupsContactGroup struct {
	Name goaliyun.String
}

type ListMyGroupsResourceList []ListMyGroupsResource

func (list *ListMyGroupsResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMyGroupsResource)
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

type ListMyGroupsContactGroupList []ListMyGroupsContactGroup

func (list *ListMyGroupsContactGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMyGroupsContactGroup)
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
