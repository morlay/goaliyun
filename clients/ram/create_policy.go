package ram

import (
	"github.com/morlay/goaliyun"
)

type CreatePolicyRequest struct {
	Description    string `position:"Query" name:"Description"`
	PolicyName     string `position:"Query" name:"PolicyName"`
	PolicyDocument string `position:"Query" name:"PolicyDocument"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *CreatePolicyRequest) Invoke(client goaliyun.Client) (*CreatePolicyResponse, error) {
	resp := &CreatePolicyResponse{}
	err := client.Request("ram", "CreatePolicy", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreatePolicyResponse struct {
	RequestId goaliyun.String
	Policy    CreatePolicyPolicy
}

type CreatePolicyPolicy struct {
	PolicyName     goaliyun.String
	PolicyType     goaliyun.String
	Description    goaliyun.String
	DefaultVersion goaliyun.String
	CreateDate     goaliyun.String
}
