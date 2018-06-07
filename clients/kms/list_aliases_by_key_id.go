package kms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListAliasesByKeyIdRequest struct {
	PageSize   int64  `position:"Query" name:"PageSize"`
	KeyId      string `position:"Query" name:"KeyId"`
	STSToken   string `position:"Query" name:"STSToken"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListAliasesByKeyIdRequest) Invoke(client goaliyun.Client) (*ListAliasesByKeyIdResponse, error) {
	resp := &ListAliasesByKeyIdResponse{}
	err := client.Request("kms", "ListAliasesByKeyId", "2016-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListAliasesByKeyIdResponse struct {
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	RequestId  goaliyun.String
	Aliases    ListAliasesByKeyIdAliasList
}

type ListAliasesByKeyIdAlias struct {
	KeyId     goaliyun.String
	AliasName goaliyun.String
	AliasArn  goaliyun.String
}

type ListAliasesByKeyIdAliasList []ListAliasesByKeyIdAlias

func (list *ListAliasesByKeyIdAliasList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAliasesByKeyIdAlias)
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
