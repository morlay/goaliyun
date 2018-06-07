package csb

import (
	"github.com/morlay/goaliyun"
)

type UpdateProjectRequest struct {
	Data     string `position:"Body" name:"Data"`
	CsbId    int64  `position:"Query" name:"CsbId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *UpdateProjectRequest) Invoke(client goaliyun.Client) (*UpdateProjectResponse, error) {
	resp := &UpdateProjectResponse{}
	err := client.Request("csb", "UpdateProject", "2017-11-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateProjectResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
}
