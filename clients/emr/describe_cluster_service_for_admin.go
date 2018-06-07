package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeClusterServiceForAdminRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeClusterServiceForAdminRequest) Invoke(client goaliyun.Client) (*DescribeClusterServiceForAdminResponse, error) {
	resp := &DescribeClusterServiceForAdminResponse{}
	err := client.Request("emr", "DescribeClusterServiceForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClusterServiceForAdminResponse struct {
	RequestId   goaliyun.String
	ServiceInfo DescribeClusterServiceForAdminServiceInfo
}

type DescribeClusterServiceForAdminServiceInfo struct {
	ServiceName                  goaliyun.String
	ServiceVersion               goaliyun.String
	ServiceStatus                goaliyun.String
	NeedRestartInfo              goaliyun.String
	NeedRestartNum               goaliyun.Integer
	ServiceActionList            DescribeClusterServiceForAdminServiceActionList
	ClusterServiceSummaryList    DescribeClusterServiceForAdminClusterServiceSummaryList
	NeedRestartHostIdList        DescribeClusterServiceForAdminNeedRestartHostIdListList
	NeedRestartComponentNameList DescribeClusterServiceForAdminNeedRestartComponentNameListList
}

type DescribeClusterServiceForAdminServiceAction struct {
	ServiceName   goaliyun.String
	ComponentName goaliyun.String
	ActionName    goaliyun.String
	Command       goaliyun.String
	DisplayName   goaliyun.String
}

type DescribeClusterServiceForAdminClusterServiceSummary struct {
	Key         goaliyun.String
	DisplayName goaliyun.String
	Value       goaliyun.String
	Status      goaliyun.String
	Type        goaliyun.String
	AlertInfo   goaliyun.String
}

type DescribeClusterServiceForAdminServiceActionList []DescribeClusterServiceForAdminServiceAction

func (list *DescribeClusterServiceForAdminServiceActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceForAdminServiceAction)
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

type DescribeClusterServiceForAdminClusterServiceSummaryList []DescribeClusterServiceForAdminClusterServiceSummary

func (list *DescribeClusterServiceForAdminClusterServiceSummaryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceForAdminClusterServiceSummary)
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

type DescribeClusterServiceForAdminNeedRestartHostIdListList []goaliyun.String

func (list *DescribeClusterServiceForAdminNeedRestartHostIdListList) UnmarshalJSON(data []byte) error {
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

type DescribeClusterServiceForAdminNeedRestartComponentNameListList []goaliyun.String

func (list *DescribeClusterServiceForAdminNeedRestartComponentNameListList) UnmarshalJSON(data []byte) error {
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
