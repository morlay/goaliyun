package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeClusterBasicInfoRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeClusterBasicInfoRequest) Invoke(client goaliyun.Client) (*DescribeClusterBasicInfoResponse, error) {
	resp := &DescribeClusterBasicInfoResponse{}
	err := client.Request("emr", "DescribeClusterBasicInfo", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClusterBasicInfoResponse struct {
	RequestId   goaliyun.String
	ClusterInfo DescribeClusterBasicInfoClusterInfo
}

type DescribeClusterBasicInfoClusterInfo struct {
	Id                     goaliyun.String
	BizId                  goaliyun.String
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
	VpcId                  goaliyun.String
	VSwitchId              goaliyun.String
	NetType                goaliyun.String
	UserDefinedEmrEcsRole  goaliyun.String
	IoOptimized            bool
	InstanceGeneration     goaliyun.String
	BootstrapFailed        bool
	Configurations         goaliyun.String
	ImageId                goaliyun.String
	SecurityGroupId        goaliyun.String
	SecurityGroupName      goaliyun.String
	EasEnable              bool
	BootstrapActionList    DescribeClusterBasicInfoBootstrapActionList
	SoftwareInfo           DescribeClusterBasicInfoSoftwareInfo
	FailReason             DescribeClusterBasicInfoFailReason
}

type DescribeClusterBasicInfoBootstrapAction struct {
	Name goaliyun.String
	Path goaliyun.String
	Arg  goaliyun.String
}

type DescribeClusterBasicInfoSoftwareInfo struct {
	EmrVer      goaliyun.String
	ClusterType goaliyun.String
	Softwares   DescribeClusterBasicInfoSoftwareList
}

type DescribeClusterBasicInfoSoftware struct {
	DisplayName goaliyun.String
	Name        goaliyun.String
	OnlyDisplay bool
	StartTpe    goaliyun.Integer
	Version     goaliyun.String
}

type DescribeClusterBasicInfoFailReason struct {
	ErrorCode goaliyun.String
	ErrorMsg  goaliyun.String
	RequestId goaliyun.String
}

type DescribeClusterBasicInfoBootstrapActionList []DescribeClusterBasicInfoBootstrapAction

func (list *DescribeClusterBasicInfoBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterBasicInfoBootstrapAction)
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

type DescribeClusterBasicInfoSoftwareList []DescribeClusterBasicInfoSoftware

func (list *DescribeClusterBasicInfoSoftwareList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterBasicInfoSoftware)
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
