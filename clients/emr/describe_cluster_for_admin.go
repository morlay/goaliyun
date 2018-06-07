package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeClusterForAdminRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeClusterForAdminRequest) Invoke(client goaliyun.Client) (*DescribeClusterForAdminResponse, error) {
	resp := &DescribeClusterForAdminResponse{}
	err := client.Request("emr", "DescribeClusterForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClusterForAdminResponse struct {
	RequestId   goaliyun.String
	ClusterInfo DescribeClusterForAdminClusterInfo
}

type DescribeClusterForAdminClusterInfo struct {
	Id                     goaliyun.String
	RegionId               goaliyun.String
	ZoneId                 goaliyun.String
	Name                   goaliyun.String
	CreateType             goaliyun.String
	StartTime              goaliyun.Integer
	StopTime               goaliyun.Integer
	LogEnable              bool
	LogPath                goaliyun.String
	Status                 goaliyun.String
	HighAvailabilityEnable bool
	ChargeType             goaliyun.String
	ExpiredTime            goaliyun.Integer
	Period                 goaliyun.Integer
	RunningTime            goaliyun.Integer
	MasterNodeTotal        goaliyun.Integer
	MasterNodeInService    goaliyun.Integer
	CoreNodeTotal          goaliyun.Integer
	CoreNodeInService      goaliyun.Integer
	TaskNodeTotal          goaliyun.Integer
	TaskNodeInService      goaliyun.Integer
	ShowSoftwareInterface  bool
	CreateResource         goaliyun.String
	VpcId                  goaliyun.String
	VSwitchId              goaliyun.String
	NetType                goaliyun.String
	UserDefinedEmrEcsRole  goaliyun.String
	IoOptimized            bool
	InstanceGeneration     goaliyun.String
	ImageId                goaliyun.String
	SecurityGroupId        goaliyun.String
	SecurityGroupName      goaliyun.String
	BootstrapFailed        bool
	Configurations         goaliyun.String
	HostGroupList          DescribeClusterForAdminHostGroupList
	EcsOrderInfoList       DescribeClusterForAdminEcsOrderInfoList
	BootstrapActionList    DescribeClusterForAdminBootstrapActionList
	FailReason             DescribeClusterForAdminFailReason
	SoftwareInfo           DescribeClusterForAdminSoftwareInfo
}

type DescribeClusterForAdminHostGroup struct {
	HostGroupId    goaliyun.String
	HostGroupName  goaliyun.String
	HostGroupType  goaliyun.String
	ChargeType     goaliyun.String
	Period         goaliyun.String
	NodeCount      goaliyun.Integer
	InstanceType   goaliyun.String
	CpuCore        goaliyun.Integer
	MemoryCapacity goaliyun.Integer
	DiskType       goaliyun.String
	DiskCapacity   goaliyun.Integer
	DiskCount      goaliyun.Integer
	BandWidth      goaliyun.String
	Nodes          DescribeClusterForAdminNodeList
}

type DescribeClusterForAdminNode struct {
	ZoneId         goaliyun.String
	InstanceId     goaliyun.String
	Status         goaliyun.String
	PubIp          goaliyun.String
	InnerIp        goaliyun.String
	ExpiredTime    goaliyun.String
	EmrExpiredTime goaliyun.String
	DaemonInfos    DescribeClusterForAdminDaemonInfoList
	DiskInfos      DescribeClusterForAdminDiskInfoList
}

type DescribeClusterForAdminDaemonInfo struct {
	Name goaliyun.String
}

type DescribeClusterForAdminDiskInfo struct {
	Device   goaliyun.String
	DiskName goaliyun.String
	DiskId   goaliyun.String
	Type     goaliyun.String
	Size     goaliyun.Integer
}

type DescribeClusterForAdminEcsOrderInfo struct {
	NodeType       goaliyun.String
	InstanceType   goaliyun.String
	CpuCore        goaliyun.Integer
	MemoryCapacity goaliyun.Integer
	DiskType       goaliyun.String
	DiskCapacity   goaliyun.Integer
	DiskCount      goaliyun.Integer
	BandWidth      goaliyun.String
	Nodes1         DescribeClusterForAdminNode2List
}

type DescribeClusterForAdminNode2 struct {
	ZoneId         goaliyun.String
	InstanceId     goaliyun.String
	Status         goaliyun.String
	PubIp          goaliyun.String
	InnerIp        goaliyun.String
	ExpiredTime    goaliyun.String
	EmrExpiredTime goaliyun.String
	DaemonInfos3   DescribeClusterForAdminDaemonInfo5List
	DiskInfos4     DescribeClusterForAdminDiskInfo6List
}

type DescribeClusterForAdminDaemonInfo5 struct {
	Name goaliyun.String
}

type DescribeClusterForAdminDiskInfo6 struct {
	Device   goaliyun.String
	DiskName goaliyun.String
	DiskId   goaliyun.String
	Type     goaliyun.String
	Size     goaliyun.Integer
}

type DescribeClusterForAdminBootstrapAction struct {
	Name goaliyun.String
	Path goaliyun.String
	Arg  goaliyun.String
}

type DescribeClusterForAdminFailReason struct {
	ErrorCode goaliyun.String
	ErrorMsg  goaliyun.String
	RequestId goaliyun.String
}

type DescribeClusterForAdminSoftwareInfo struct {
	EmrVer      goaliyun.String
	ClusterType goaliyun.String
	Softwares   DescribeClusterForAdminSoftwareList
}

type DescribeClusterForAdminSoftware struct {
	DisplayName goaliyun.String
	Name        goaliyun.String
	OnlyDisplay bool
	StartTpe    goaliyun.Integer
	Version     goaliyun.String
}

type DescribeClusterForAdminHostGroupList []DescribeClusterForAdminHostGroup

func (list *DescribeClusterForAdminHostGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterForAdminHostGroup)
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

type DescribeClusterForAdminEcsOrderInfoList []DescribeClusterForAdminEcsOrderInfo

func (list *DescribeClusterForAdminEcsOrderInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterForAdminEcsOrderInfo)
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

type DescribeClusterForAdminBootstrapActionList []DescribeClusterForAdminBootstrapAction

func (list *DescribeClusterForAdminBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterForAdminBootstrapAction)
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

type DescribeClusterForAdminNodeList []DescribeClusterForAdminNode

func (list *DescribeClusterForAdminNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterForAdminNode)
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

type DescribeClusterForAdminDaemonInfoList []DescribeClusterForAdminDaemonInfo

func (list *DescribeClusterForAdminDaemonInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterForAdminDaemonInfo)
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

type DescribeClusterForAdminDiskInfoList []DescribeClusterForAdminDiskInfo

func (list *DescribeClusterForAdminDiskInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterForAdminDiskInfo)
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

type DescribeClusterForAdminNode2List []DescribeClusterForAdminNode2

func (list *DescribeClusterForAdminNode2List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterForAdminNode2)
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

type DescribeClusterForAdminDaemonInfo5List []DescribeClusterForAdminDaemonInfo5

func (list *DescribeClusterForAdminDaemonInfo5List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterForAdminDaemonInfo5)
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

type DescribeClusterForAdminDiskInfo6List []DescribeClusterForAdminDiskInfo6

func (list *DescribeClusterForAdminDiskInfo6List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterForAdminDiskInfo6)
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

type DescribeClusterForAdminSoftwareList []DescribeClusterForAdminSoftware

func (list *DescribeClusterForAdminSoftwareList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterForAdminSoftware)
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
