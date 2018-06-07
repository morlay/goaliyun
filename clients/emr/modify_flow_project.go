package emr

import (
	"github.com/morlay/goaliyun"
)

type ModifyFlowProjectRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	Description     string `position:"Query" name:"Description"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ModifyFlowProjectRequest) Invoke(client goaliyun.Client) (*ModifyFlowProjectResponse, error) {
	resp := &ModifyFlowProjectResponse{}
	err := client.Request("emr", "ModifyFlowProject", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyFlowProjectResponse struct {
	RequestId goaliyun.String
	Data      bool
}
