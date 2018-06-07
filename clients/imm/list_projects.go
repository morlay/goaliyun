package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListProjectsRequest struct {
	MaxKeys  int64  `position:"Query" name:"MaxKeys"`
	Marker   string `position:"Query" name:"Marker"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListProjectsRequest) Invoke(client goaliyun.Client) (*ListProjectsResponse, error) {
	resp := &ListProjectsResponse{}
	err := client.Request("imm", "ListProjects", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListProjectsResponse struct {
	RequestId  goaliyun.String
	NextMarker goaliyun.String
	Projects   ListProjectsProjectsItemList
}

type ListProjectsProjectsItem struct {
	Project     goaliyun.String
	Endpoint    goaliyun.String
	ServiceRole goaliyun.String
	CreateTime  goaliyun.String
	ModifyTime  goaliyun.String
	Engines     ListProjectsEnginesItemList
	Indexers    ListProjectsIndexersItemList
}

type ListProjectsEnginesItem struct {
	Name   goaliyun.String
	JobTtl goaliyun.Integer
}

type ListProjectsIndexersItem struct {
	Name   goaliyun.String
	Status goaliyun.String
}

type ListProjectsProjectsItemList []ListProjectsProjectsItem

func (list *ListProjectsProjectsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListProjectsProjectsItem)
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

type ListProjectsEnginesItemList []ListProjectsEnginesItem

func (list *ListProjectsEnginesItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListProjectsEnginesItem)
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

type ListProjectsIndexersItemList []ListProjectsIndexersItem

func (list *ListProjectsIndexersItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListProjectsIndexersItem)
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
