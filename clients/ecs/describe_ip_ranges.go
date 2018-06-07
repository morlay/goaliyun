package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeIpRangesRequest struct {
	NicType              string `position:"Query" name:"NicType"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	ClusterId            string `position:"Query" name:"ClusterId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeIpRangesRequest) Invoke(client goaliyun.Client) (*DescribeIpRangesResponse, error) {
	resp := &DescribeIpRangesResponse{}
	err := client.Request("ecs", "DescribeIpRanges", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeIpRangesResponse struct {
	RequestId  goaliyun.String
	RegionId   goaliyun.String
	ClusterId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	IpRanges   DescribeIpRangesIpRangeList
}

type DescribeIpRangesIpRange struct {
	IpAddress goaliyun.String
	NicType   goaliyun.String
}

type DescribeIpRangesIpRangeList []DescribeIpRangesIpRange

func (list *DescribeIpRangesIpRangeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeIpRangesIpRange)
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
