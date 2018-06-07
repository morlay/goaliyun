package kms

import (
	"github.com/morlay/goaliyun"
)

type CreateAliasRequest struct {
	AliasName string `position:"Query" name:"AliasName"`
	KeyId     string `position:"Query" name:"KeyId"`
	STSToken  string `position:"Query" name:"STSToken"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *CreateAliasRequest) Invoke(client goaliyun.Client) (*CreateAliasResponse, error) {
	resp := &CreateAliasResponse{}
	err := client.Request("kms", "CreateAlias", "2016-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateAliasResponse struct {
	RequestId goaliyun.String
}
