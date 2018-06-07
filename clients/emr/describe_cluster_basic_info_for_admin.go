package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeClusterBasicInfoForAdminRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeClusterBasicInfoForAdminRequest) Invoke(client goaliyun.Client) (*DescribeClusterBasicInfoForAdminResponse, error) {
	resp := &DescribeClusterBasicInfoForAdminResponse{}
	err := client.Request("emr", "DescribeClusterBasicInfoForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClusterBasicInfoForAdminResponse struct {
	RequestId   goaliyun.String
	ClusterInfo DescribeClusterBasicInfoForAdminClusterInfo
}

type DescribeClusterBasicInfoForAdminClusterInfo struct {
	Id                     goaliyun.String
	BizId                  goaliyun.String
	RegionId               goaliyun.String
	ZoneId                 goaliyun.String
	Name                   goaliyun.String
	CreateType             goaliyun.String
	ClusterType            goaliyun.String
	StartTime              goaliyun.Integer
	StopTime               goaliyun.Integer
	LogEnable              bool
	LogPath                goaliyun.String
	UserId                 goaliyun.String
	Status                 goaliyun.String
	HighAvailabilityEnable bool
	PayType                goaliyun.String
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
	EcmClusterId           goaliyun.String
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
	EcsOrderInfoList       DescribeClusterBasicInfoForAdminEcsOrderInfoList
	BootstrapActionList    DescribeClusterBasicInfoForAdminBootstrapActionList
	FailReason             DescribeClusterBasicInfoForAdminFailReason
	SoftwareInfo           DescribeClusterBasicInfoForAdminSoftwareInfo
}

type DescribeClusterBasicInfoForAdminEcsOrderInfo struct {
	NodeType       goaliyun.String
	InstanceType   goaliyun.String
	CpuCore        goaliyun.Integer
	MemoryCapacity goaliyun.Integer
	DiskType       goaliyun.String
	DiskCapacity   goaliyun.Integer
	DiskCount      goaliyun.Integer
	BandWidth      goaliyun.String
	Nodes          DescribeClusterBasicInfoForAdminNodeList
}

type DescribeClusterBasicInfoForAdminNode struct {
	ZoneId         goaliyun.String
	InstanceId     goaliyun.String
	Status         goaliyun.String
	PubIp          goaliyun.String
	InnerIp        goaliyun.String
	ExpiredTime    goaliyun.String
	EmrExpiredTime goaliyun.String
	DiskInfos      DescribeClusterBasicInfoForAdminDiskInfoList
}

type DescribeClusterBasicInfoForAdminDiskInfo struct {
	Device   goaliyun.String
	DiskName goaliyun.String
	DiskId   goaliyun.String
	Type     goaliyun.String
	Size     goaliyun.Integer
}

type DescribeClusterBasicInfoForAdminBootstrapAction struct {
	Name goaliyun.String
	Path goaliyun.String
	Arg  goaliyun.String
}

type DescribeClusterBasicInfoForAdminFailReason struct {
	ErrorCode goaliyun.String
	ErrorMsg  goaliyun.String
	RequestId goaliyun.String
}

type DescribeClusterBasicInfoForAdminSoftwareInfo struct {
	EmrVer      goaliyun.String
	ClusterType goaliyun.String
	Softwares   DescribeClusterBasicInfoForAdminSoftwareList
}

type DescribeClusterBasicInfoForAdminSoftware struct {
	DisplayName goaliyun.String
	Name        goaliyun.String
	OnlyDisplay bool
	StartTpe    goaliyun.Integer
	Version     goaliyun.String
}

type DescribeClusterBasicInfoForAdminEcsOrderInfoList []DescribeClusterBasicInfoForAdminEcsOrderInfo

func (list *DescribeClusterBasicInfoForAdminEcsOrderInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterBasicInfoForAdminEcsOrderInfo)
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

type DescribeClusterBasicInfoForAdminBootstrapActionList []DescribeClusterBasicInfoForAdminBootstrapAction

func (list *DescribeClusterBasicInfoForAdminBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterBasicInfoForAdminBootstrapAction)
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

type DescribeClusterBasicInfoForAdminNodeList []DescribeClusterBasicInfoForAdminNode

func (list *DescribeClusterBasicInfoForAdminNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterBasicInfoForAdminNode)
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

type DescribeClusterBasicInfoForAdminDiskInfoList []DescribeClusterBasicInfoForAdminDiskInfo

func (list *DescribeClusterBasicInfoForAdminDiskInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterBasicInfoForAdminDiskInfo)
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

type DescribeClusterBasicInfoForAdminSoftwareList []DescribeClusterBasicInfoForAdminSoftware

func (list *DescribeClusterBasicInfoForAdminSoftwareList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterBasicInfoForAdminSoftware)
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
