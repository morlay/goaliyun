package imm

import (
	"github.com/morlay/goaliyun"
)

type DeleteProjectRequest struct {
	Project  string `position:"Query" name:"Project"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteProjectRequest) Invoke(client goaliyun.Client) (*DeleteProjectResponse, error) {
	resp := &DeleteProjectResponse{}
	err := client.Request("imm", "DeleteProject", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteProjectResponse struct {
	RequestId goaliyun.String
}
