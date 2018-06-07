package kms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListKeysRequest struct {
	PageSize   int64  `position:"Query" name:"PageSize"`
	STSToken   string `position:"Query" name:"STSToken"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListKeysRequest) Invoke(client goaliyun.Client) (*ListKeysResponse, error) {
	resp := &ListKeysResponse{}
	err := client.Request("kms", "ListKeys", "2016-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListKeysResponse struct {
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	RequestId  goaliyun.String
	Keys       ListKeysKeyList
}

type ListKeysKey struct {
	KeyId  goaliyun.String
	KeyArn goaliyun.String
}

type ListKeysKeyList []ListKeysKey

func (list *ListKeysKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListKeysKey)
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
