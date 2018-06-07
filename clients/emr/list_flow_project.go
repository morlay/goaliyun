package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListFlowProjectRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListFlowProjectRequest) Invoke(client goaliyun.Client) (*ListFlowProjectResponse, error) {
	resp := &ListFlowProjectResponse{}
	err := client.Request("emr", "ListFlowProject", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListFlowProjectResponse struct {
	RequestId  goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Total      goaliyun.Integer
	Projects   ListFlowProjectProjectList
}

type ListFlowProjectProject struct {
	Id          goaliyun.String
	GmtCreate   goaliyun.Integer
	GmtModified goaliyun.Integer
	UserId      goaliyun.String
	Name        goaliyun.String
	Description goaliyun.String
}

type ListFlowProjectProjectList []ListFlowProjectProject

func (list *ListFlowProjectProjectList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFlowProjectProject)
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
