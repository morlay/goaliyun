package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeExecutionPlanRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeExecutionPlanRequest) Invoke(client goaliyun.Client) (*DescribeExecutionPlanResponse, error) {
	resp := &DescribeExecutionPlanResponse{}
	err := client.Request("emr", "DescribeExecutionPlan", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeExecutionPlanResponse struct {
	RequestId             goaliyun.String
	Id                    goaliyun.String
	Name                  goaliyun.String
	Status                goaliyun.String
	Strategy              goaliyun.String
	TimeInterval          goaliyun.Integer
	StartTime             goaliyun.Integer
	TimeUnit              goaliyun.String
	DayOfWeek             goaliyun.String
	DayOfMonth            goaliyun.String
	CreateClusterOnDemand bool
	ClusterId             goaliyun.String
	ClusterName           goaliyun.String
	WorkflowApp           goaliyun.String
	ExecutionPlanVersion  goaliyun.Integer
	JobInfoList           DescribeExecutionPlanJobInfoList
	ClusterInfo           DescribeExecutionPlanClusterInfo
}

type DescribeExecutionPlanJobInfo struct {
	Id           goaliyun.String
	Name         goaliyun.String
	Type         goaliyun.String
	RunParameter goaliyun.String
	FailAct      goaliyun.String
}

type DescribeExecutionPlanClusterInfo struct {
	Name                   goaliyun.String
	ZoneId                 goaliyun.String
	LogEnable              bool
	LogPath                goaliyun.String
	SecurityGroupId        goaliyun.String
	EmrVer                 goaliyun.String
	ClusterType            goaliyun.String
	HighAvailabilityEnable bool
	VpcId                  goaliyun.String
	VSwitchId              goaliyun.String
	NetType                goaliyun.String
	IoOptimized            bool
	InstanceGeneration     goaliyun.String
	Configurations         goaliyun.String
	EasEnable              bool
	UserDefinedEmrEcsRole  goaliyun.String
	EcsOrders              DescribeExecutionPlanEcsOrderInfoList
	BootstrapActionList    DescribeExecutionPlanBootstrapActionList
	SoftwareInfo           DescribeExecutionPlanSoftwareInfo
}

type DescribeExecutionPlanEcsOrderInfo struct {
	Index        goaliyun.Integer
	NodeCount    goaliyun.Integer
	InstanceType goaliyun.String
	DiskType     goaliyun.String
	DiskCapacity goaliyun.Integer
	NodeType     goaliyun.String
	DiskCount    goaliyun.Integer
}

type DescribeExecutionPlanBootstrapAction struct {
	Name goaliyun.String
	Path goaliyun.String
	Arg  goaliyun.String
}

type DescribeExecutionPlanSoftwareInfo struct {
	EmrVer      goaliyun.String
	ClusterType goaliyun.String
	Softwares   DescribeExecutionPlanSoftwareList
}

type DescribeExecutionPlanSoftware struct {
	DisplayName goaliyun.String
	Name        goaliyun.String
	OnlyDisplay bool
	StartTpe    goaliyun.Integer
	Version     goaliyun.String
	Optional    bool
}

type DescribeExecutionPlanJobInfoList []DescribeExecutionPlanJobInfo

func (list *DescribeExecutionPlanJobInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeExecutionPlanJobInfo)
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

type DescribeExecutionPlanEcsOrderInfoList []DescribeExecutionPlanEcsOrderInfo

func (list *DescribeExecutionPlanEcsOrderInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeExecutionPlanEcsOrderInfo)
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

type DescribeExecutionPlanBootstrapActionList []DescribeExecutionPlanBootstrapAction

func (list *DescribeExecutionPlanBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeExecutionPlanBootstrapAction)
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

type DescribeExecutionPlanSoftwareList []DescribeExecutionPlanSoftware

func (list *DescribeExecutionPlanSoftwareList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeExecutionPlanSoftware)
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
