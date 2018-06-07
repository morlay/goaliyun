package emr

import (
	"github.com/morlay/goaliyun"
)

type CreateFlowProjectUserRequest struct {
	UserAccountId string `position:"Query" name:"UserAccountId"`
	ProjectId     string `position:"Query" name:"ProjectId"`
	UserName      string `position:"Query" name:"UserName"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *CreateFlowProjectUserRequest) Invoke(client goaliyun.Client) (*CreateFlowProjectUserResponse, error) {
	resp := &CreateFlowProjectUserResponse{}
	err := client.Request("emr", "CreateFlowProjectUser", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateFlowProjectUserResponse struct {
	RequestId goaliyun.String
	Data      bool
}
