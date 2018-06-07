package kms

import (
	"github.com/morlay/goaliyun"
)

type DescribeKeyRequest struct {
	KeyId    string `position:"Query" name:"KeyId"`
	STSToken string `position:"Query" name:"STSToken"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DescribeKeyRequest) Invoke(client goaliyun.Client) (*DescribeKeyResponse, error) {
	resp := &DescribeKeyResponse{}
	err := client.Request("kms", "DescribeKey", "2016-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeKeyResponse struct {
	RequestId   goaliyun.String
	KeyMetadata DescribeKeyKeyMetadata
}

type DescribeKeyKeyMetadata struct {
	CreationDate       goaliyun.String
	Description        goaliyun.String
	KeyId              goaliyun.String
	KeyState           goaliyun.String
	KeyUsage           goaliyun.String
	DeleteDate         goaliyun.String
	Creator            goaliyun.String
	Arn                goaliyun.String
	Origin             goaliyun.String
	MaterialExpireTime goaliyun.String
}
