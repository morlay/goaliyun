package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeClusterServiceRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeClusterServiceRequest) Invoke(client goaliyun.Client) (*DescribeClusterServiceResponse, error) {
	resp := &DescribeClusterServiceResponse{}
	err := client.Request("emr", "DescribeClusterService", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClusterServiceResponse struct {
	RequestId   goaliyun.String
	ServiceInfo DescribeClusterServiceServiceInfo
}

type DescribeClusterServiceServiceInfo struct {
	ServiceName                  goaliyun.String
	ServiceVersion               goaliyun.String
	ServiceStatus                goaliyun.String
	NeedRestartInfo              goaliyun.String
	NeedRestartNum               goaliyun.Integer
	ServiceActionList            DescribeClusterServiceServiceActionList
	ClusterServiceSummaryList    DescribeClusterServiceClusterServiceSummaryList
	NeedRestartHostIdList        DescribeClusterServiceNeedRestartHostIdListList
	NeedRestartComponentNameList DescribeClusterServiceNeedRestartComponentNameListList
}

type DescribeClusterServiceServiceAction struct {
	ServiceName   goaliyun.String
	ComponentName goaliyun.String
	ActionName    goaliyun.String
	Command       goaliyun.String
	DisplayName   goaliyun.String
}

type DescribeClusterServiceClusterServiceSummary struct {
	Key         goaliyun.String
	DisplayName goaliyun.String
	Value       goaliyun.String
	Status      goaliyun.String
	Type        goaliyun.String
	AlertInfo   goaliyun.String
}

type DescribeClusterServiceServiceActionList []DescribeClusterServiceServiceAction

func (list *DescribeClusterServiceServiceActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceServiceAction)
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

type DescribeClusterServiceClusterServiceSummaryList []DescribeClusterServiceClusterServiceSummary

func (list *DescribeClusterServiceClusterServiceSummaryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceClusterServiceSummary)
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

type DescribeClusterServiceNeedRestartHostIdListList []goaliyun.String

func (list *DescribeClusterServiceNeedRestartHostIdListList) UnmarshalJSON(data []byte) error {
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

type DescribeClusterServiceNeedRestartComponentNameListList []goaliyun.String

func (list *DescribeClusterServiceNeedRestartComponentNameListList) UnmarshalJSON(data []byte) error {
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
