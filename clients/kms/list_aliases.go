package kms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListAliasesRequest struct {
	PageSize   int64  `position:"Query" name:"PageSize"`
	STSToken   string `position:"Query" name:"STSToken"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListAliasesRequest) Invoke(client goaliyun.Client) (*ListAliasesResponse, error) {
	resp := &ListAliasesResponse{}
	err := client.Request("kms", "ListAliases", "2016-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListAliasesResponse struct {
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	RequestId  goaliyun.String
	Aliases    ListAliasesAliasList
}

type ListAliasesAlias struct {
	KeyId     goaliyun.String
	AliasName goaliyun.String
	AliasArn  goaliyun.String
}

type ListAliasesAliasList []ListAliasesAlias

func (list *ListAliasesAliasList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAliasesAlias)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
