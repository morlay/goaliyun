package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRegionsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeRegionsRequest) Invoke(client goaliyun.Client) (*DescribeRegionsResponse, error) {
	resp := &DescribeRegionsResponse{}
	err := client.Request("rds", "DescribeRegions", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRegionsResponse struct {
	RequestId goaliyun.String
	Regions   DescribeRegionsRDSRegionList
}

type DescribeRegionsRDSRegion struct {
	RegionId goaliyun.String
	ZoneId   goaliyun.String
}

type DescribeRegionsRDSRegionList []DescribeRegionsRDSRegion

func (list *DescribeRegionsRDSRegionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsRDSRegion)
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
