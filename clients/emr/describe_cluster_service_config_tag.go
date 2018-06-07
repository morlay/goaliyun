package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeClusterServiceConfigTagRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ConfigTag       string `position:"Query" name:"ConfigTag"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeClusterServiceConfigTagRequest) Invoke(client goaliyun.Client) (*DescribeClusterServiceConfigTagResponse, error) {
	resp := &DescribeClusterServiceConfigTagResponse{}
	err := client.Request("emr", "DescribeClusterServiceConfigTag", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClusterServiceConfigTagResponse struct {
	RequestId     goaliyun.String
	ConfigTagList DescribeClusterServiceConfigTagConfigTagList
}

type DescribeClusterServiceConfigTagConfigTag struct {
	Tag       goaliyun.String
	TagDesc   goaliyun.String
	ValueList DescribeClusterServiceConfigTagValueList
}

type DescribeClusterServiceConfigTagValue struct {
	Value     goaliyun.String
	ValueDesc goaliyun.String
}

type DescribeClusterServiceConfigTagConfigTagList []DescribeClusterServiceConfigTagConfigTag

func (list *DescribeClusterServiceConfigTagConfigTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceConfigTagConfigTag)
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

type DescribeClusterServiceConfigTagValueList []DescribeClusterServiceConfigTagValue

func (list *DescribeClusterServiceConfigTagValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceConfigTagValue)
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
