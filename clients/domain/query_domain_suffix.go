package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryDomainSuffixRequest struct {
	Lang     string `position:"Query" name:"Lang"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *QueryDomainSuffixRequest) Invoke(client goaliyun.Client) (*QueryDomainSuffixResponse, error) {
	resp := &QueryDomainSuffixResponse{}
	err := client.Request("domain", "QueryDomainSuffix", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryDomainSuffixResponse struct {
	RequestId  goaliyun.String
	SuffixList QueryDomainSuffixSuffixListList
}

type QueryDomainSuffixSuffixListList []goaliyun.String

func (list *QueryDomainSuffixSuffixListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
