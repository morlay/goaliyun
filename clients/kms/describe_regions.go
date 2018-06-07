package kms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRegionsRequest struct {
	STSToken string `position:"Query" name:"STSToken"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DescribeRegionsRequest) Invoke(client goaliyun.Client) (*DescribeRegionsResponse, error) {
	resp := &DescribeRegionsResponse{}
	err := client.Request("kms", "DescribeRegions", "2016-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRegionsResponse struct {
	RequestId goaliyun.String
	Regions   DescribeRegionsRegionList
}

type DescribeRegionsRegion struct {
	RegionId goaliyun.String
}

type DescribeRegionsRegionList []DescribeRegionsRegion

func (list *DescribeRegionsRegionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsRegion)
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
