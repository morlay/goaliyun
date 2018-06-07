package rds

import (
	"github.com/morlay/goaliyun"
)

type CreateDampPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int64  `position:"Query" name:"Priority"`
	TimeRules            string `position:"Query" name:"TimeRules"`
	ActionRules          string `position:"Query" name:"ActionRules"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	Handlers             string `position:"Query" name:"Handlers"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	PolicyName           string `position:"Query" name:"PolicyName"`
	SourceRules          string `position:"Query" name:"SourceRules"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateDampPolicyRequest) Invoke(client goaliyun.Client) (*CreateDampPolicyResponse, error) {
	resp := &CreateDampPolicyResponse{}
	err := client.Request("rds", "CreateDampPolicy", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateDampPolicyResponse struct {
	RequestId  goaliyun.String
	PolicyId   goaliyun.String
	PolicyName goaliyun.String
}
