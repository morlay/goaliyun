package hsm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRegionsRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeRegionsRequest) Invoke(client goaliyun.Client) (*DescribeRegionsResponse, error) {
	resp := &DescribeRegionsResponse{}
	err := client.Request("hsm", "DescribeRegions", "2018-01-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRegionsResponse struct {
	RequestId goaliyun.String
	Regions   DescribeRegionsRegionList
}

type DescribeRegionsRegion struct {
	RegionId goaliyun.String
	Zones    DescribeRegionsZoneList
}

type DescribeRegionsZone struct {
	ZoneId goaliyun.String
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

type DescribeRegionsZoneList []DescribeRegionsZone

func (list *DescribeRegionsZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsZone)
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
