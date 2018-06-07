package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeClusterRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeClusterRequest) Invoke(client goaliyun.Client) (*DescribeClusterResponse, error) {
	resp := &DescribeClusterResponse{}
	err := client.Request("emr", "DescribeCluster", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClusterResponse struct {
	RequestId   goaliyun.String
	ClusterInfo DescribeClusterClusterInfo
}

type DescribeClusterClusterInfo struct {
	Id                     goaliyun.String
	RegionId               goaliyun.String
	ZoneId                 goaliyun.String
	Name                   goaliyun.String
	CreateType             goaliyun.String
	StartTime              goaliyun.Integer
	StopTime               goaliyun.Integer
	LogEnable              bool
	LogPath                goaliyun.String
	UserId                 goaliyun.String
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
	EasEnable              bool
	EcsOrderInfoList       DescribeClusterEcsOrderInfoList
	BootstrapActionList    DescribeClusterBootstrapActionList
	FailReason             DescribeClusterFailReason
	SoftwareInfo           DescribeClusterSoftwareInfo
}

type DescribeClusterEcsOrderInfo struct {
	NodeType       goaliyun.String
	InstanceType   goaliyun.String
	CpuCore        goaliyun.Integer
	MemoryCapacity goaliyun.Integer
	DiskType       goaliyun.String
	DiskCapacity   goaliyun.Integer
	DiskCount      goaliyun.Integer
	BandWidth      goaliyun.String
	Nodes          DescribeClusterNodeList
}

type DescribeClusterNode struct {
	ZoneId         goaliyun.String
	InstanceId     goaliyun.String
	Status         goaliyun.String
	PubIp          goaliyun.String
	InnerIp        goaliyun.String
	ExpiredTime    goaliyun.String
	EmrExpiredTime goaliyun.String
	DaemonInfos    DescribeClusterDaemonInfoList
	DiskInfos      DescribeClusterDiskInfoList
}

type DescribeClusterDaemonInfo struct {
	Name goaliyun.String
}

type DescribeClusterDiskInfo struct {
	Device   goaliyun.String
	DiskName goaliyun.String
	DiskId   goaliyun.String
	Type     goaliyun.String
	Size     goaliyun.Integer
}

type DescribeClusterBootstrapAction struct {
	Name goaliyun.String
	Path goaliyun.String
	Arg  goaliyun.String
}

type DescribeClusterFailReason struct {
	ErrorCode goaliyun.String
	ErrorMsg  goaliyun.String
	RequestId goaliyun.String
}

type DescribeClusterSoftwareInfo struct {
	EmrVer      goaliyun.String
	ClusterType goaliyun.String
	Softwares   DescribeClusterSoftwareList
}

type DescribeClusterSoftware struct {
	DisplayName goaliyun.String
	Name        goaliyun.String
	OnlyDisplay bool
	StartTpe    goaliyun.Integer
	Version     goaliyun.String
}

type DescribeClusterEcsOrderInfoList []DescribeClusterEcsOrderInfo

func (list *DescribeClusterEcsOrderInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterEcsOrderInfo)
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

type DescribeClusterBootstrapActionList []DescribeClusterBootstrapAction

func (list *DescribeClusterBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterBootstrapAction)
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

type DescribeClusterNodeList []DescribeClusterNode

func (list *DescribeClusterNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterNode)
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

type DescribeClusterDaemonInfoList []DescribeClusterDaemonInfo

func (list *DescribeClusterDaemonInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterDaemonInfo)
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

type DescribeClusterDiskInfoList []DescribeClusterDiskInfo

func (list *DescribeClusterDiskInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterDiskInfo)
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

type DescribeClusterSoftwareList []DescribeClusterSoftware

func (list *DescribeClusterSoftwareList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterSoftware)
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
