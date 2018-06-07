package ocs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRegionsRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeRegionsRequest) Invoke(client goaliyun.Client) (*DescribeRegionsResponse, error) {
	resp := &DescribeRegionsResponse{}
	err := client.Request("ocs", "DescribeRegions", "2015-03-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRegionsResponse struct {
	RequestId                  goaliyun.String
	GetDescribeRegionsResponse DescribeRegionsGetDescribeRegionsResponse
}

type DescribeRegionsGetDescribeRegionsResponse struct {
	Total    goaliyun.Integer
	PageNo   goaliyun.Integer
	PageSize goaliyun.Integer
	Regions  DescribeRegionsRegionList
}

type DescribeRegionsRegion struct {
	RegionId  goaliyun.String
	IzNumber  goaliyun.String
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
