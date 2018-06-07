package csb

import (
	"github.com/morlay/goaliyun"
)

type DeleteProjectRequest struct {
	CsbId     int64  `position:"Query" name:"CsbId"`
	ProjectId int64  `position:"Query" name:"ProjectId"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *DeleteProjectRequest) Invoke(client goaliyun.Client) (*DeleteProjectResponse, error) {
	resp := &DeleteProjectResponse{}
	err := client.Request("csb", "DeleteProject", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteProjectResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
}
