package vod

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type SearchEditingProjectRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	Title                string `position:"Query" name:"Title"`
	PageNo               int64  `position:"Query" name:"PageNo"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	SortBy               string `position:"Query" name:"SortBy"`
	Status               string `position:"Query" name:"Status"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *SearchEditingProjectRequest) Invoke(client goaliyun.Client) (*SearchEditingProjectResponse, error) {
	resp := &SearchEditingProjectResponse{}
	err := client.Request("vod", "SearchEditingProject", "2017-03-21", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SearchEditingProjectResponse struct {
	RequestId   goaliyun.String
	Total       goaliyun.Integer
	ProjectList SearchEditingProjectProjectList
}

type SearchEditingProjectProject struct {
	ProjectId    goaliyun.String
	CreationTime goaliyun.String
	ModifiedTime goaliyun.String
	Status       goaliyun.String
	Description  goaliyun.String
	Title        goaliyun.String
	CoverURL     goaliyun.String
}

type SearchEditingProjectProjectList []SearchEditingProjectProject

func (list *SearchEditingProjectProjectList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchEditingProjectProject)
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
