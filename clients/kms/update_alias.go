package kms

import (
	"github.com/morlay/goaliyun"
)

type UpdateAliasRequest struct {
	AliasName string `position:"Query" name:"AliasName"`
	KeyId     string `position:"Query" name:"KeyId"`
	STSToken  string `position:"Query" name:"STSToken"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *UpdateAliasRequest) Invoke(client goaliyun.Client) (*UpdateAliasResponse, error) {
	resp := &UpdateAliasResponse{}
	err := client.Request("kms", "UpdateAlias", "2016-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateAliasResponse struct {
	RequestId goaliyun.String
}
