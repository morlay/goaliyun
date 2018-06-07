package emr

import (
	"github.com/morlay/goaliyun"
)

type DeleteFlowProjectRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteFlowProjectRequest) Invoke(client goaliyun.Client) (*DeleteFlowProjectResponse, error) {
	resp := &DeleteFlowProjectResponse{}
	err := client.Request("emr", "DeleteFlowProject", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteFlowProjectResponse struct {
	RequestId goaliyun.String
	Data      bool
}
