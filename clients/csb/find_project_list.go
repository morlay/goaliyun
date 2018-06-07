package csb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type FindProjectListRequest struct {
	ProjectName string `position:"Query" name:"ProjectName"`
	CsbId       int64  `position:"Query" name:"CsbId"`
	PageNum     int64  `position:"Query" name:"PageNum"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *FindProjectListRequest) Invoke(client goaliyun.Client) (*FindProjectListResponse, error) {
	resp := &FindProjectListResponse{}
	err := client.Request("csb", "FindProjectList", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FindProjectListResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      FindProjectListData
}

type FindProjectListData struct {
	CurrentPage goaliyun.Integer
	PageNumber  goaliyun.Integer
	Total       goaliyun.Integer
	ProjectList FindProjectListProjectList
}

type FindProjectListProject struct {
	ApiNum               goaliyun.Integer
	CsbId                goaliyun.Integer
	DeleteFlag           goaliyun.Integer
	Description          goaliyun.String
	GmtCreate            goaliyun.Integer
	GmtModified          goaliyun.Integer
	Id                   goaliyun.Integer
	InterfaceJarLocation goaliyun.String
	InterfaceJarName     goaliyun.String
	JarFileKey           goaliyun.String
	OwnerId              goaliyun.String
	ProjectName          goaliyun.String
	ProjectOwnerEmail    goaliyun.String
	ProjectOwnerName     goaliyun.String
	ProjectOwnerPhoneNum goaliyun.String
	Status               goaliyun.Integer
	UserId               goaliyun.String
}

type FindProjectListProjectList []FindProjectListProject

func (list *FindProjectListProjectList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindProjectListProject)
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
