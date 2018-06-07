package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeTagsRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DistinctKey          string `position:"Query" name:"DistinctKey"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeTagsRequest) Invoke(client goaliyun.Client) (*DescribeTagsResponse, error) {
	resp := &DescribeTagsResponse{}
	err := client.Request("slb", "DescribeTags", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeTagsResponse struct {
	RequestId  goaliyun.String
	PageSize   goaliyun.Integer
	PageNumber goaliyun.Integer
	TotalCount goaliyun.Integer
	TagSets    DescribeTagsTagSetList
}

type DescribeTagsTagSet struct {
	TagKey        goaliyun.String
	TagValue      goaliyun.String
	InstanceCount goaliyun.Integer
}

type DescribeTagsTagSetList []DescribeTagsTagSet

func (list *DescribeTagsTagSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTagsTagSet)
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
