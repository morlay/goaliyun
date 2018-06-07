package dds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRegionsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeRegionsRequest) Invoke(client goaliyun.Client) (*DescribeRegionsResponse, error) {
	resp := &DescribeRegionsResponse{}
	err := client.Request("dds", "DescribeRegions", "2015-12-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRegionsResponse struct {
	RequestId goaliyun.String
	Regions   DescribeRegionsDdsRegionList
}

type DescribeRegionsDdsRegion struct {
	RegionId goaliyun.String
	ZoneIds  goaliyun.String
	Zones    DescribeRegionsZoneList
}

type DescribeRegionsZone struct {
	ZoneId     goaliyun.String
	VpcEnabled bool
}

type DescribeRegionsDdsRegionList []DescribeRegionsDdsRegion

func (list *DescribeRegionsDdsRegionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsDdsRegion)
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
