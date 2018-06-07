package ram

import (
	"github.com/morlay/goaliyun"
)

type GetPolicyVersionRequest struct {
	VersionId  string `position:"Query" name:"VersionId"`
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *GetPolicyVersionRequest) Invoke(client goaliyun.Client) (*GetPolicyVersionResponse, error) {
	resp := &GetPolicyVersionResponse{}
	err := client.Request("ram", "GetPolicyVersion", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetPolicyVersionResponse struct {
	RequestId     goaliyun.String
	PolicyVersion GetPolicyVersionPolicyVersion
}

type GetPolicyVersionPolicyVersion struct {
	VersionId        goaliyun.String
	IsDefaultVersion bool
	PolicyDocument   goaliyun.String
	CreateDate       goaliyun.String
}
