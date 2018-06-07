package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeClusterTemplateRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeClusterTemplateRequest) Invoke(client goaliyun.Client) (*DescribeClusterTemplateResponse, error) {
	resp := &DescribeClusterTemplateResponse{}
	err := client.Request("emr", "DescribeClusterTemplate", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClusterTemplateResponse struct {
	RequestId    goaliyun.String
	TemplateInfo DescribeClusterTemplateTemplateInfo
}

type DescribeClusterTemplateTemplateInfo struct {
	Id                     goaliyun.String
	TemplateName           goaliyun.String
	LogEnable              bool
	LogPath                goaliyun.String
	UserId                 goaliyun.String
	UserDefinedEmrEcsRole  goaliyun.String
	MasterNodeTotal        goaliyun.Integer
	VpcId                  goaliyun.String
	VSwitchId              goaliyun.String
	NetType                goaliyun.String
	IoOptimized            bool
	InstanceGeneration     goaliyun.String
	HighAvailabilityEnable bool
	EasEnable              bool
	BootstrapActionList    DescribeClusterTemplateBootstrapActionList
	HostGroupList          DescribeClusterTemplateHostGroupList
}

type DescribeClusterTemplateBootstrapAction struct {
	Name goaliyun.String
	Path goaliyun.String
	Arg  goaliyun.String
}

type DescribeClusterTemplateHostGroup struct {
	HostGroupId   goaliyun.String
	HostGroupName goaliyun.String
	HostGroupType goaliyun.String
	ChargeType    goaliyun.String
	Period        goaliyun.String
	NodeCount     goaliyun.Integer
	InstanceType  goaliyun.String
	DiskType      goaliyun.String
	DiskCapacity  goaliyun.Integer
	DiskCount     goaliyun.Integer
}

type DescribeClusterTemplateBootstrapActionList []DescribeClusterTemplateBootstrapAction

func (list *DescribeClusterTemplateBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterTemplateBootstrapAction)
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

type DescribeClusterTemplateHostGroupList []DescribeClusterTemplateHostGroup

func (list *DescribeClusterTemplateHostGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterTemplateHostGroup)
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
