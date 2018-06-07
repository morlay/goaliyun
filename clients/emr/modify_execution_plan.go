package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ModifyExecutionPlanRequest struct {
	ResourceOwnerId        int64                                      `position:"Query" name:"ResourceOwnerId"`
	LogPath                string                                     `position:"Query" name:"LogPath"`
	TimeInterval           int64                                      `position:"Query" name:"TimeInterval"`
	ClusterName            string                                     `position:"Query" name:"ClusterName"`
	Configurations         string                                     `position:"Query" name:"Configurations"`
	IoOptimized            string                                     `position:"Query" name:"IoOptimized"`
	SecurityGroupId        string                                     `position:"Query" name:"SecurityGroupId"`
	EasEnable              string                                     `position:"Query" name:"EasEnable"`
	CreateClusterOnDemand  string                                     `position:"Query" name:"CreateClusterOnDemand"`
	StartTime              int64                                      `position:"Query" name:"StartTime"`
	JobIdLists             *ModifyExecutionPlanJobIdListList          `position:"Query" type:"Repeated" name:"JobIdList"`
	DayOfMonth             string                                     `position:"Query" name:"DayOfMonth"`
	BootstrapActions       *ModifyExecutionPlanBootstrapActionList    `position:"Query" type:"Repeated" name:"BootstrapAction"`
	UseLocalMetaDb         string                                     `position:"Query" name:"UseLocalMetaDb"`
	EmrVer                 string                                     `position:"Query" name:"EmrVer"`
	Id                     string                                     `position:"Query" name:"Id"`
	UserDefinedEmrEcsRole  string                                     `position:"Query" name:"UserDefinedEmrEcsRole"`
	IsOpenPublicIp         string                                     `position:"Query" name:"IsOpenPublicIp"`
	ExecutionPlanVersion   int64                                      `position:"Query" name:"ExecutionPlanVersion"`
	ClusterId              string                                     `position:"Query" name:"ClusterId"`
	TimeUnit               string                                     `position:"Query" name:"TimeUnit"`
	InstanceGeneration     string                                     `position:"Query" name:"InstanceGeneration"`
	ClusterType            string                                     `position:"Query" name:"ClusterType"`
	VSwitchId              string                                     `position:"Query" name:"VSwitchId"`
	OptionSoftWareLists    *ModifyExecutionPlanOptionSoftWareListList `position:"Query" type:"Repeated" name:"OptionSoftWareList"`
	VpcId                  string                                     `position:"Query" name:"VpcId"`
	NetType                string                                     `position:"Query" name:"NetType"`
	WorkflowDefinition     string                                     `position:"Query" name:"WorkflowDefinition"`
	EcsOrders              *ModifyExecutionPlanEcsOrderList           `position:"Query" type:"Repeated" name:"EcsOrder"`
	Name                   string                                     `position:"Query" name:"Name"`
	ZoneId                 string                                     `position:"Query" name:"ZoneId"`
	DayOfWeek              string                                     `position:"Query" name:"DayOfWeek"`
	Strategy               string                                     `position:"Query" name:"Strategy"`
	HighAvailabilityEnable string                                     `position:"Query" name:"HighAvailabilityEnable"`
	LogEnable              string                                     `position:"Query" name:"LogEnable"`
	RegionId               string                                     `position:"Query" name:"RegionId"`
}

func (req *ModifyExecutionPlanRequest) Invoke(client goaliyun.Client) (*ModifyExecutionPlanResponse, error) {
	resp := &ModifyExecutionPlanResponse{}
	err := client.Request("emr", "ModifyExecutionPlan", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyExecutionPlanBootstrapAction struct {
	Name string `name:"Name"`
	Path string `name:"Path"`
	Arg  string `name:"Arg"`
}

type ModifyExecutionPlanEcsOrder struct {
	Index        int64  `name:"Index"`
	NodeCount    int64  `name:"NodeCount"`
	InstanceType string `name:"InstanceType"`
	DiskType     string `name:"DiskType"`
	DiskCapacity int64  `name:"DiskCapacity"`
	NodeType     string `name:"NodeType"`
	DiskCount    int64  `name:"DiskCount"`
}

type ModifyExecutionPlanResponse struct {
	RequestId goaliyun.String
}

type ModifyExecutionPlanJobIdListList []string

func (list *ModifyExecutionPlanJobIdListList) UnmarshalJSON(data []byte) error {
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

type ModifyExecutionPlanBootstrapActionList []ModifyExecutionPlanBootstrapAction

func (list *ModifyExecutionPlanBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyExecutionPlanBootstrapAction)
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

type ModifyExecutionPlanOptionSoftWareListList []string

func (list *ModifyExecutionPlanOptionSoftWareListList) UnmarshalJSON(data []byte) error {
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

type ModifyExecutionPlanEcsOrderList []ModifyExecutionPlanEcsOrder

func (list *ModifyExecutionPlanEcsOrderList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyExecutionPlanEcsOrder)
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
