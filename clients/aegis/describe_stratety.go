package aegis

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeStratetyRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeStratetyRequest) Invoke(client goaliyun.Client) (*DescribeStratetyResponse, error) {
	resp := &DescribeStratetyResponse{}
	err := client.Request("aegis", "DescribeStratety", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeStratetyResponse struct {
	RequestId  goaliyun.String
	Count      goaliyun.Integer
	Strategies DescribeStratetyDataList
}

type DescribeStratetyData struct {
	CycleDays      goaliyun.Integer
	Id             goaliyun.Integer
	CycleStartTime goaliyun.Integer
	Type           goaliyun.Integer
	Name           goaliyun.String
	RiskCount      goaliyun.Integer
	EcsCount       goaliyun.Integer
	ConfigTargets  DescribeStratetyConfigTargetList
}

type DescribeStratetyConfigTarget struct {
	Flag       goaliyun.String
	TargetType goaliyun.String
	Target     goaliyun.String
}

type DescribeStratetyDataList []DescribeStratetyData

func (list *DescribeStratetyDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeStratetyData)
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

type DescribeStratetyConfigTargetList []DescribeStratetyConfigTarget

func (list *DescribeStratetyConfigTargetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeStratetyConfigTarget)
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
