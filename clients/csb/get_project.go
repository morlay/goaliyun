package csb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetProjectRequest struct {
	ProjectName string `position:"Query" name:"ProjectName"`
	CsbId       int64  `position:"Query" name:"CsbId"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *GetProjectRequest) Invoke(client goaliyun.Client) (*GetProjectResponse, error) {
	resp := &GetProjectResponse{}
	err := client.Request("csb", "GetProject", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetProjectResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      GetProjectData
}

type GetProjectData struct {
	ProjectList GetProjectProjectList
}

type GetProjectProject struct {
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

type GetProjectProjectList []GetProjectProject

func (list *GetProjectProjectList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetProjectProject)
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
