package kms

import (
	"github.com/morlay/goaliyun"
)

type DeleteAliasRequest struct {
	AliasName string `position:"Query" name:"AliasName"`
	STSToken  string `position:"Query" name:"STSToken"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *DeleteAliasRequest) Invoke(client goaliyun.Client) (*DeleteAliasResponse, error) {
	resp := &DeleteAliasResponse{}
	err := client.Request("kms", "DeleteAlias", "2016-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteAliasResponse struct {
	RequestId goaliyun.String
}
