package aegis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeStratetyDetailRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeStratetyDetailRequest) Invoke(client goaliyun.Client) (*DescribeStratetyDetailResponse, error) {
	resp := &DescribeStratetyDetailResponse{}
	err := client.Request("aegis", "DescribeStratetyDetail", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeStratetyDetailResponse struct {
	RequestId goaliyun.String
	Strategy  DescribeStratetyDetailStrategy
}

type DescribeStratetyDetailStrategy struct {
	CycleDays                        goaliyun.Integer
	Name                             goaliyun.String
	Id                               goaliyun.Integer
	CycleStartTime                   goaliyun.Integer
	Type                             goaliyun.Integer
	RiskTypeWhiteListQueryResultList DescribeStratetyDetailRiskTypeWhiteListQueryResultList
}

type DescribeStratetyDetailRiskTypeWhiteListQueryResult struct {
	TypeName goaliyun.String
	Alias    goaliyun.String
	On       bool
	SubTypes DescribeStratetyDetailSubTypList
}

type DescribeStratetyDetailSubTyp struct {
	TypeName goaliyun.String
	Alias    goaliyun.String
	On       bool
}

type DescribeStratetyDetailRiskTypeWhiteListQueryResultList []DescribeStratetyDetailRiskTypeWhiteListQueryResult

func (list *DescribeStratetyDetailRiskTypeWhiteListQueryResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeStratetyDetailRiskTypeWhiteListQueryResult)
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

type DescribeStratetyDetailSubTypList []DescribeStratetyDetailSubTyp

func (list *DescribeStratetyDetailSubTypList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeStratetyDetailSubTyp)
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
