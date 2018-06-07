package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeClusterServiceConfigHistoryRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ConfigVersion   string `position:"Query" name:"ConfigVersion"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeClusterServiceConfigHistoryRequest) Invoke(client goaliyun.Client) (*DescribeClusterServiceConfigHistoryResponse, error) {
	resp := &DescribeClusterServiceConfigHistoryResponse{}
	err := client.Request("emr", "DescribeClusterServiceConfigHistory", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClusterServiceConfigHistoryResponse struct {
	RequestId goaliyun.String
	Config    DescribeClusterServiceConfigHistoryConfig
}

type DescribeClusterServiceConfigHistoryConfig struct {
	ServiceName     goaliyun.String
	ConfigVersion   goaliyun.String
	Applied         bool
	CreateTime      goaliyun.String
	Author          goaliyun.String
	Comment         goaliyun.String
	ConfigValueList DescribeClusterServiceConfigHistoryConfigValueList
}

type DescribeClusterServiceConfigHistoryConfigValue struct {
	ConfigName          goaliyun.String
	ConfigItemValueList DescribeClusterServiceConfigHistoryConfigItemValueList
}

type DescribeClusterServiceConfigHistoryConfigItemValue struct {
	ItemName   goaliyun.String
	Value      goaliyun.String
	OldValue   goaliyun.String
	ChangeType goaliyun.String
}

type DescribeClusterServiceConfigHistoryConfigValueList []DescribeClusterServiceConfigHistoryConfigValue

func (list *DescribeClusterServiceConfigHistoryConfigValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceConfigHistoryConfigValue)
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

type DescribeClusterServiceConfigHistoryConfigItemValueList []DescribeClusterServiceConfigHistoryConfigItemValue

func (list *DescribeClusterServiceConfigHistoryConfigItemValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceConfigHistoryConfigItemValue)
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
