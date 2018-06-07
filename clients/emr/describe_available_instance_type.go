package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeAvailableInstanceTypeRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeAvailableInstanceTypeRequest) Invoke(client goaliyun.Client) (*DescribeAvailableInstanceTypeResponse, error) {
	resp := &DescribeAvailableInstanceTypeResponse{}
	err := client.Request("emr", "DescribeAvailableInstanceType", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeAvailableInstanceTypeResponse struct {
	RequestId                    goaliyun.String
	EmrSupportedInstanceTypeList DescribeAvailableInstanceTypeEmrSupportInstanceTypeList
}

type DescribeAvailableInstanceTypeEmrSupportInstanceType struct {
	ClusterType             goaliyun.String
	NodeTypeSupportInfoList DescribeAvailableInstanceTypeClusterNodeTypeSupportInfoList
}

type DescribeAvailableInstanceTypeClusterNodeTypeSupportInfo struct {
	ClusterNodeType         goaliyun.String
	SupportInstanceTypeList DescribeAvailableInstanceTypeSupportInstanceTypeListList
}

type DescribeAvailableInstanceTypeEmrSupportInstanceTypeList []DescribeAvailableInstanceTypeEmrSupportInstanceType

func (list *DescribeAvailableInstanceTypeEmrSupportInstanceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAvailableInstanceTypeEmrSupportInstanceType)
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

type DescribeAvailableInstanceTypeClusterNodeTypeSupportInfoList []DescribeAvailableInstanceTypeClusterNodeTypeSupportInfo

func (list *DescribeAvailableInstanceTypeClusterNodeTypeSupportInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAvailableInstanceTypeClusterNodeTypeSupportInfo)
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

type DescribeAvailableInstanceTypeSupportInstanceTypeListList []goaliyun.String

func (list *DescribeAvailableInstanceTypeSupportInstanceTypeListList) UnmarshalJSON(data []byte) error {
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
