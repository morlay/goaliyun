package kms

import (
	"github.com/morlay/goaliyun"
)

type CreateKeyRequest struct {
	KeyUsage    string `position:"Query" name:"KeyUsage"`
	Origin      string `position:"Query" name:"Origin"`
	Description string `position:"Query" name:"Description"`
	STSToken    string `position:"Query" name:"STSToken"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *CreateKeyRequest) Invoke(client goaliyun.Client) (*CreateKeyResponse, error) {
	resp := &CreateKeyResponse{}
	err := client.Request("kms", "CreateKey", "2016-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateKeyResponse struct {
	RequestId   goaliyun.String
	KeyMetadata CreateKeyKeyMetadata
}

type CreateKeyKeyMetadata struct {
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
