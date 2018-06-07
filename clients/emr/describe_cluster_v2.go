package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeClusterV2Request struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeClusterV2Request) Invoke(client goaliyun.Client) (*DescribeClusterV2Response, error) {
	resp := &DescribeClusterV2Response{}
	err := client.Request("emr", "DescribeClusterV2", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClusterV2Response struct {
	RequestId   goaliyun.String
	ClusterInfo DescribeClusterV2ClusterInfo
}

type DescribeClusterV2ClusterInfo struct {
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
	HostGroupList          DescribeClusterV2HostGroupList
	BootstrapActionList    DescribeClusterV2BootstrapActionList
	FailReason             DescribeClusterV2FailReason
	SoftwareInfo           DescribeClusterV2SoftwareInfo
}

type DescribeClusterV2HostGroup struct {
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
	Nodes          DescribeClusterV2NodeList
}

type DescribeClusterV2Node struct {
	ZoneId         goaliyun.String
	InstanceId     goaliyun.String
	Status         goaliyun.String
	PubIp          goaliyun.String
	InnerIp        goaliyun.String
	ExpiredTime    goaliyun.String
	CreateTime     goaliyun.String
	EmrExpiredTime goaliyun.String
	DaemonInfos    DescribeClusterV2DaemonInfoList
	DiskInfos      DescribeClusterV2DiskInfoList
}

type DescribeClusterV2DaemonInfo struct {
	Name goaliyun.String
}

type DescribeClusterV2DiskInfo struct {
	Device   goaliyun.String
	DiskName goaliyun.String
	DiskId   goaliyun.String
	Type     goaliyun.String
	Size     goaliyun.Integer
}

type DescribeClusterV2BootstrapAction struct {
	Name goaliyun.String
	Path goaliyun.String
	Arg  goaliyun.String
}

type DescribeClusterV2FailReason struct {
	ErrorCode goaliyun.String
	ErrorMsg  goaliyun.String
	RequestId goaliyun.String
}

type DescribeClusterV2SoftwareInfo struct {
	EmrVer      goaliyun.String
	ClusterType goaliyun.String
	Softwares   DescribeClusterV2SoftwareList
}

type DescribeClusterV2Software struct {
	DisplayName goaliyun.String
	Name        goaliyun.String
	OnlyDisplay bool
	StartTpe    goaliyun.Integer
	Version     goaliyun.String
}

type DescribeClusterV2HostGroupList []DescribeClusterV2HostGroup

func (list *DescribeClusterV2HostGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterV2HostGroup)
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

type DescribeClusterV2BootstrapActionList []DescribeClusterV2BootstrapAction

func (list *DescribeClusterV2BootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterV2BootstrapAction)
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

type DescribeClusterV2NodeList []DescribeClusterV2Node

func (list *DescribeClusterV2NodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterV2Node)
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

type DescribeClusterV2DaemonInfoList []DescribeClusterV2DaemonInfo

func (list *DescribeClusterV2DaemonInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterV2DaemonInfo)
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

type DescribeClusterV2DiskInfoList []DescribeClusterV2DiskInfo

func (list *DescribeClusterV2DiskInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterV2DiskInfo)
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

type DescribeClusterV2SoftwareList []DescribeClusterV2Software

func (list *DescribeClusterV2SoftwareList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterV2Software)
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
