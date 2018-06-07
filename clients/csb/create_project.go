package csb

import (
	"github.com/morlay/goaliyun"
)

type CreateProjectRequest struct {
	Data     string `position:"Body" name:"Data"`
	CsbId    int64  `position:"Query" name:"CsbId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *CreateProjectRequest) Invoke(client goaliyun.Client) (*CreateProjectResponse, error) {
	resp := &CreateProjectResponse{}
	err := client.Request("csb", "CreateProject", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateProjectResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      CreateProjectData
}

type CreateProjectData struct {
	Id goaliyun.Integer
}
