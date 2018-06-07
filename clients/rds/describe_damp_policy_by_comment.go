package rds

import (
	"github.com/morlay/goaliyun"
)

type DescribeDampPolicyByCommentRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	PolicyName           string `position:"Query" name:"PolicyName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDampPolicyByCommentRequest) Invoke(client goaliyun.Client) (*DescribeDampPolicyByCommentResponse, error) {
	resp := &DescribeDampPolicyByCommentResponse{}
	err := client.Request("rds", "DescribeDampPolicyByComment", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDampPolicyByCommentResponse struct {
	RequestId   goaliyun.String
	Policy      goaliyun.String
	TimeRules   goaliyun.String
	ActionRules goaliyun.String
	SourceRules goaliyun.String
	Handler     goaliyun.String
}
