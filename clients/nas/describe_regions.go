package nas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRegionsRequest struct {
	PageSize   int64  `position:"Query" name:"PageSize"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeRegionsRequest) Invoke(client goaliyun.Client) (*DescribeRegionsResponse, error) {
	resp := &DescribeRegionsResponse{}
	err := client.Request("nas", "DescribeRegions", "2017-06-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRegionsResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageSize   goaliyun.Integer
	PageNumber goaliyun.Integer
	Regions    DescribeRegionsRegionList
}

type DescribeRegionsRegion struct {
	RegionId  goaliyun.String
	LocalName goaliyun.String
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
