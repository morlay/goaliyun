package rds

import (
	"github.com/morlay/goaliyun"
)

type CreatePolicyWithSpecifiedPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	PolicyId             string `position:"Query" name:"PolicyId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreatePolicyWithSpecifiedPolicyRequest) Invoke(client goaliyun.Client) (*CreatePolicyWithSpecifiedPolicyResponse, error) {
	resp := &CreatePolicyWithSpecifiedPolicyResponse{}
	err := client.Request("rds", "CreatePolicyWithSpecifiedPolicy", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreatePolicyWithSpecifiedPolicyResponse struct {
	RequestId goaliyun.String
}
