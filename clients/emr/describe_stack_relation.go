package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeStackRelationRequest struct {
	ResourceOwnerId   int64                                      `position:"Query" name:"ResourceOwnerId"`
	EmrVersion        string                                     `position:"Query" name:"EmrVersion"`
	ClusterId         string                                     `position:"Query" name:"ClusterId"`
	StackVersion      string                                     `position:"Query" name:"StackVersion"`
	StackVersionLists *DescribeStackRelationStackVersionListList `position:"Query" type:"Repeated" name:"StackVersionList"`
	EmrVersionLists   *DescribeStackRelationEmrVersionListList   `position:"Query" type:"Repeated" name:"EmrVersionList"`
	RegionId          string                                     `position:"Query" name:"RegionId"`
}

func (req *DescribeStackRelationRequest) Invoke(client goaliyun.Client) (*DescribeStackRelationResponse, error) {
	resp := &DescribeStackRelationResponse{}
	err := client.Request("emr", "DescribeStackRelation", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeStackRelationResponse struct {
	RequestId            goaliyun.String
	Success              bool
	EmrStackRelationList DescribeStackRelationEmrStackRelationList
}

type DescribeStackRelationEmrStackRelation struct {
	EmrVersion   goaliyun.String
	StackVersion goaliyun.String
}

type DescribeStackRelationStackVersionListList []string

func (list *DescribeStackRelationStackVersionListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeStackRelationEmrVersionListList []string

func (list *DescribeStackRelationEmrVersionListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeStackRelationEmrStackRelationList []DescribeStackRelationEmrStackRelation

func (list *DescribeStackRelationEmrStackRelationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeStackRelationEmrStackRelation)
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
