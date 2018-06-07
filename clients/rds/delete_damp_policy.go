package rds

import (
	"github.com/morlay/goaliyun"
)

type DeleteDampPolicyRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	PolicyName           string `position:"Query" name:"PolicyName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteDampPolicyRequest) Invoke(client goaliyun.Client) (*DeleteDampPolicyResponse, error) {
	resp := &DeleteDampPolicyResponse{}
	err := client.Request("rds", "DeleteDampPolicy", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteDampPolicyResponse struct {
	RequestId goaliyun.String
}
