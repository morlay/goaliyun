package emr

import (
	"github.com/morlay/goaliyun"
)

type DeleteFlowProjectUserRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	UserName        string `position:"Query" name:"UserName"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DeleteFlowProjectUserRequest) Invoke(client goaliyun.Client) (*DeleteFlowProjectUserResponse, error) {
	resp := &DeleteFlowProjectUserResponse{}
	err := client.Request("emr", "DeleteFlowProjectUser", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteFlowProjectUserResponse struct {
	RequestId goaliyun.String
	Data      bool
}
