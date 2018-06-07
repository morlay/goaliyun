package ram

import (
	"github.com/morlay/goaliyun"
)

type GetPolicyRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *GetPolicyRequest) Invoke(client goaliyun.Client) (*GetPolicyResponse, error) {
	resp := &GetPolicyResponse{}
	err := client.Request("ram", "GetPolicy", "2015-05-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetPolicyResponse struct {
	RequestId goaliyun.String
	Policy    GetPolicyPolicy
}

type GetPolicyPolicy struct {
	PolicyName      goaliyun.String
	PolicyType      goaliyun.String
	Description     goaliyun.String
	DefaultVersion  goaliyun.String
	PolicyDocument  goaliyun.String
	CreateDate      goaliyun.String
	UpdateDate      goaliyun.String
	AttachmentCount goaliyun.Integer
}
