package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterForAdminRequest struct {
	ResourceOwnerId   int64                                    `position:"Query" name:"ResourceOwnerId"`
	StatusLists       *ListClusterForAdminStatusListList       `position:"Query" type:"Repeated" name:"StatusList"`
	FuzzyName         string                                   `position:"Query" name:"FuzzyName"`
	UserId            string                                   `position:"Query" name:"UserId"`
	PageNumber        int64                                    `position:"Query" name:"PageNumber"`
	EcmClusterIdLists *ListClusterForAdminEcmClusterIdListList `position:"Query" type:"Repeated" name:"EcmClusterIdList"`
	ClusterIdLists    *ListClusterForAdminClusterIdListList    `position:"Query" type:"Repeated" name:"ClusterIdList"`
	PayTypeLists      *ListClusterForAdminPayTypeListList      `position:"Query" type:"Repeated" name:"PayTypeList"`
	Name              string                                   `position:"Query" name:"Name"`
	PageSize          int64                                    `position:"Query" name:"PageSize"`
	EmrVersion        string                                   `position:"Query" name:"EmrVersion"`
	Resize            string                                   `position:"Query" name:"Resize"`
	ClusterTypeLists  *ListClusterForAdminClusterTypeListList  `position:"Query" type:"Repeated" name:"ClusterTypeList"`
	RegionId          string                                   `position:"Query" name:"RegionId"`
}

func (req *ListClusterForAdminRequest) Invoke(client goaliyun.Client) (*ListClusterForAdminResponse, error) {
	resp := &ListClusterForAdminResponse{}
	err := client.Request("emr", "ListClusterForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterForAdminResponse struct {
	RequestId   goaliyun.String
	TotalCount  goaliyun.Integer
	PageNumber  goaliyun.Integer
	PageSize    goaliyun.Integer
	ClusterList ListClusterForAdminClusterList
}

type ListClusterForAdminCluster struct {
	Id                     goaliyun.String
	Name                   goaliyun.String
	BizId                  goaliyun.String
	LusterType             goaliyun.String
	CreateTime             goaliyun.Integer
	RunningTime            goaliyun.Integer
	Status                 goaliyun.String
	PayType                goaliyun.String
	RegionId               goaliyun.String
	EasEnable              bool
	DepositType            goaliyun.String
	UseLocalMetadb         bool
	HighAvailabilityEnable bool
	NodeCount              goaliyun.Integer
	ExpiredTime            goaliyun.Integer
	NetType                goaliyun.String
	HasUncompletedOrder    bool
	OrderList              goaliyun.String
	CreateResource         goaliyun.String
	UserId                 goaliyun.String
	EcmClusterId           goaliyun.String
	EmrVersion             goaliyun.String
	VpcId                  goaliyun.String
	VSwitchId              goaliyun.String
	OrderTaskInfo          ListClusterForAdminOrderTaskInfo
	FailReason             ListClusterForAdminFailReason
}

type ListClusterForAdminOrderTaskInfo struct {
	TargetCount  goaliyun.Integer
	CurrentCount goaliyun.Integer
	OrderIdList  goaliyun.String
}

type ListClusterForAdminFailReason struct {
	ErrorCode goaliyun.String
	ErrorMsg  goaliyun.String
	RequestId goaliyun.String
}

type ListClusterForAdminStatusListList []string

func (list *ListClusterForAdminStatusListList) UnmarshalJSON(data []byte) error {
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

type ListClusterForAdminEcmClusterIdListList []int64

func (list *ListClusterForAdminEcmClusterIdListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type ListClusterForAdminClusterIdListList []string

func (list *ListClusterForAdminClusterIdListList) UnmarshalJSON(data []byte) error {
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

type ListClusterForAdminPayTypeListList []string

func (list *ListClusterForAdminPayTypeListList) UnmarshalJSON(data []byte) error {
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

type ListClusterForAdminClusterTypeListList []string

func (list *ListClusterForAdminClusterTypeListList) UnmarshalJSON(data []byte) error {
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

type ListClusterForAdminClusterList []ListClusterForAdminCluster

func (list *ListClusterForAdminClusterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterForAdminCluster)
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
