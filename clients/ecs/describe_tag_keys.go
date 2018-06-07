package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeTagKeysRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceType         string `position:"Query" name:"ResourceType"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeTagKeysRequest) Invoke(client goaliyun.Client) (*DescribeTagKeysResponse, error) {
	resp := &DescribeTagKeysResponse{}
	err := client.Request("ecs", "DescribeTagKeys", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeTagKeysResponse struct {
	RequestId  goaliyun.String
	PageSize   goaliyun.Integer
	PageNumber goaliyun.Integer
	TotalCount goaliyun.Integer
	TagKeys    DescribeTagKeysTagKeyList
}

type DescribeTagKeysTagKeyList []goaliyun.String

func (list *DescribeTagKeysTagKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
