package r_kvstore

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
	err := client.Request("r-kvstore", "DescribeRegions", "2015-01-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRegionsResponse struct {
	RequestId goaliyun.String
	RegionIds DescribeRegionsKVStoreRegionList
}

type DescribeRegionsKVStoreRegion struct {
	RegionId  goaliyun.String
	ZoneIds   goaliyun.String
	LocalName goaliyun.String
}

type DescribeRegionsKVStoreRegionList []DescribeRegionsKVStoreRegion

func (list *DescribeRegionsKVStoreRegionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsKVStoreRegion)
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
