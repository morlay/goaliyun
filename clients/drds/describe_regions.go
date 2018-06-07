package drds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRegionsRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DescribeRegionsRequest) Invoke(client goaliyun.Client) (*DescribeRegionsResponse, error) {
	resp := &DescribeRegionsResponse{}
	err := client.Request("drds", "DescribeRegions", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRegionsResponse struct {
	RequestId   goaliyun.String
	Success     bool
	DrdsRegions DescribeRegionsDrdsRegionList
}

type DescribeRegionsDrdsRegion struct {
	RegionId           goaliyun.String
	RegionName         goaliyun.String
	ZoneId             goaliyun.String
	ZoneName           goaliyun.String
	InstanceSeriesList DescribeRegionsInstanceSeriesList
}

type DescribeRegionsInstanceSeries struct {
	SeriesId   goaliyun.String
	SeriesName goaliyun.String
	SpecList   DescribeRegionsSpecList
}

type DescribeRegionsSpec struct {
	SpecId   goaliyun.String
	SpecName goaliyun.String
}

type DescribeRegionsDrdsRegionList []DescribeRegionsDrdsRegion

func (list *DescribeRegionsDrdsRegionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsDrdsRegion)
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

type DescribeRegionsInstanceSeriesList []DescribeRegionsInstanceSeries

func (list *DescribeRegionsInstanceSeriesList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsInstanceSeries)
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

type DescribeRegionsSpecList []DescribeRegionsSpec

func (list *DescribeRegionsSpecList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsSpec)
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
