package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterTemplatesRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	BizId           string `position:"Query" name:"BizId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterTemplatesRequest) Invoke(client goaliyun.Client) (*ListClusterTemplatesResponse, error) {
	resp := &ListClusterTemplatesResponse{}
	err := client.Request("emr", "ListClusterTemplates", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterTemplatesResponse struct {
	RequestId        goaliyun.String
	TotalCount       goaliyun.String
	PageNumber       goaliyun.String
	PageSize         goaliyun.String
	TemplateInfoList ListClusterTemplatesTemplateInfoList
}

type ListClusterTemplatesTemplateInfo struct {
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
	BootstrapActionList    ListClusterTemplatesBootstrapActionList
	HostGroupList          ListClusterTemplatesHostGroupList
	SoftwareInfo           ListClusterTemplatesSoftwareInfo
}

type ListClusterTemplatesBootstrapAction struct {
	Name goaliyun.String
	Path goaliyun.String
	Arg  goaliyun.String
}

type ListClusterTemplatesHostGroup struct {
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

type ListClusterTemplatesSoftwareInfo struct {
	EmrVer      goaliyun.String
	ClusterType goaliyun.String
	Softwares   ListClusterTemplatesSoftwareList
}

type ListClusterTemplatesSoftware struct {
	DisplayName goaliyun.String
	Name        goaliyun.String
	OnlyDisplay bool
	StartTpe    goaliyun.Integer
	Version     goaliyun.String
}

type ListClusterTemplatesTemplateInfoList []ListClusterTemplatesTemplateInfo

func (list *ListClusterTemplatesTemplateInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterTemplatesTemplateInfo)
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

type ListClusterTemplatesBootstrapActionList []ListClusterTemplatesBootstrapAction

func (list *ListClusterTemplatesBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterTemplatesBootstrapAction)
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

type ListClusterTemplatesHostGroupList []ListClusterTemplatesHostGroup

func (list *ListClusterTemplatesHostGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterTemplatesHostGroup)
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

type ListClusterTemplatesSoftwareList []ListClusterTemplatesSoftware

func (list *ListClusterTemplatesSoftwareList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterTemplatesSoftware)
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
