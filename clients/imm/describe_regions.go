package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRegionsRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DescribeRegionsRequest) Invoke(client goaliyun.Client) (*DescribeRegionsResponse, error) {
	resp := &DescribeRegionsResponse{}
	err := client.Request("imm", "DescribeRegions", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRegionsResponse struct {
	RequestId goaliyun.String
	Regions   DescribeRegionsRegionsItemList
}

type DescribeRegionsRegionsItem struct {
	Region   goaliyun.String
	Status   goaliyun.String
	ShowName goaliyun.String
}

type DescribeRegionsRegionsItemList []DescribeRegionsRegionsItem

func (list *DescribeRegionsRegionsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsRegionsItem)
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
