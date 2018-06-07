package emr

import (
	"github.com/morlay/goaliyun"
)

type CreateFlowProjectRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	Description     string `position:"Query" name:"Description"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CreateFlowProjectRequest) Invoke(client goaliyun.Client) (*CreateFlowProjectResponse, error) {
	resp := &CreateFlowProjectResponse{}
	err := client.Request("emr", "CreateFlowProject", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateFlowProjectResponse struct {
	RequestId goaliyun.String
	Id        goaliyun.String
}
