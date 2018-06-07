package ram

import (
	"github.com/morlay/goaliyun"
)

type CreatePolicyVersionRequest struct {
	SetAsDefault   string `position:"Query" name:"SetAsDefault"`
	PolicyName     string `position:"Query" name:"PolicyName"`
	PolicyDocument string `position:"Query" name:"PolicyDocument"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *CreatePolicyVersionRequest) Invoke(client goaliyun.Client) (*CreatePolicyVersionResponse, error) {
	resp := &CreatePolicyVersionResponse{}
	err := client.Request("ram", "CreatePolicyVersion", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreatePolicyVersionResponse struct {
	RequestId     goaliyun.String
	PolicyVersion CreatePolicyVersionPolicyVersion
}

type CreatePolicyVersionPolicyVersion struct {
	VersionId        goaliyun.String
	IsDefaultVersion bool
	PolicyDocument   goaliyun.String
	CreateDate       goaliyun.String
}
