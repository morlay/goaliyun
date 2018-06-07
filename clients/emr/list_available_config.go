package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListAvailableConfigRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListAvailableConfigRequest) Invoke(client goaliyun.Client) (*ListAvailableConfigResponse, error) {
	resp := &ListAvailableConfigResponse{}
	err := client.Request("emr", "ListAvailableConfig", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListAvailableConfigResponse struct {
	RequestId                    goaliyun.String
	SecurityGroupTypes           ListAvailableConfigSecurityGroupTypeList
	InstanceTypes                ListAvailableConfigInstanceTypeList
	EmrVerTypes                  ListAvailableConfigEmrVerTypeList
	ZoneTypes                    ListAvailableConfigZoneTypeList
	Vpcs                         ListAvailableConfigVpcList
	EmrSupportedInstanceTypeList ListAvailableConfigEmrSupportInstanceTypeList
}

type ListAvailableConfigSecurityGroupType struct {
	Name  goaliyun.String
	State goaliyun.String
	Id    goaliyun.String
}

type ListAvailableConfigInstanceType struct {
	Classify               goaliyun.String
	Type                   goaliyun.String
	CpuCore                goaliyun.Integer
	MemSize                goaliyun.Integer
	HasCloudDisk           bool
	HasEfficiencyCloudDisk bool
	HasSSDCloudDisk        bool
}

type ListAvailableConfigEmrVerType struct {
	Name       goaliyun.String
	EcmStack   bool
	SubModules ListAvailableConfigSubModuleList
}

type ListAvailableConfigSubModule struct {
	Name         goaliyun.String
	RequiredList ListAvailableConfigRequiredList
	OptionalList ListAvailableConfigOptionalList
}

type ListAvailableConfigRequired struct {
	DisplayName goaliyun.String
	Name        goaliyun.String
	OnlyDisplay bool
	StartTpe    goaliyun.Integer
	Version     goaliyun.String
}

type ListAvailableConfigOptional struct {
	DisplayName goaliyun.String
	Name        goaliyun.String
	OnlyDisplay bool
	StartTpe    goaliyun.Integer
	Version     goaliyun.String
}

type ListAvailableConfigZoneType struct {
	Name                          goaliyun.String
	Id                            goaliyun.String
	AvailableResources            ListAvailableConfigAvailableResourceList
	AvailableResourceCreationList ListAvailableConfigAvailableResourceCreationListList
	AvailableDiskCategoryList     ListAvailableConfigAvailableDiskCategoryListList
}

type ListAvailableConfigAvailableResource struct {
	IoOptimized            bool
	SystemDiskCategories   ListAvailableConfigSystemDiskCategoryList
	DataDiskCategories     ListAvailableConfigDataDiskCategoryList
	NetworkTypes           ListAvailableConfigNetworkTypeList
	SupportedInstanceTypes ListAvailableConfigSupportedInstanceTypeList
	InstanceTypeFamilies   ListAvailableConfigInstanceTypeFamilyList
	InstanceGenerations    ListAvailableConfigInstanceGenerationList
}

type ListAvailableConfigVpc struct {
	Id             goaliyun.String
	VpcName        goaliyun.String
	CidrBlock      goaliyun.String
	VSwitchs       ListAvailableConfigVSwitchList
	SecurityGroups ListAvailableConfigSecurityGroupList
}

type ListAvailableConfigVSwitch struct {
	Id          goaliyun.String
	VswitchName goaliyun.String
	CidrBlock   goaliyun.String
	ZoneId      goaliyun.String
}

type ListAvailableConfigSecurityGroup struct {
	Name goaliyun.String
	Id   goaliyun.String
}

type ListAvailableConfigEmrSupportInstanceType struct {
	ClusterType             goaliyun.String
	NodeTypeSupportInfoList ListAvailableConfigClusterNodeTypeSupportInfoList
}

type ListAvailableConfigClusterNodeTypeSupportInfo struct {
	ClusterNodeType         goaliyun.String
	SupportInstanceTypeList ListAvailableConfigSupportInstanceTypeListList
}

type ListAvailableConfigSecurityGroupTypeList []ListAvailableConfigSecurityGroupType

func (list *ListAvailableConfigSecurityGroupTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigSecurityGroupType)
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

type ListAvailableConfigInstanceTypeList []ListAvailableConfigInstanceType

func (list *ListAvailableConfigInstanceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigInstanceType)
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

type ListAvailableConfigEmrVerTypeList []ListAvailableConfigEmrVerType

func (list *ListAvailableConfigEmrVerTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigEmrVerType)
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

type ListAvailableConfigZoneTypeList []ListAvailableConfigZoneType

func (list *ListAvailableConfigZoneTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigZoneType)
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

type ListAvailableConfigVpcList []ListAvailableConfigVpc

func (list *ListAvailableConfigVpcList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigVpc)
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

type ListAvailableConfigEmrSupportInstanceTypeList []ListAvailableConfigEmrSupportInstanceType

func (list *ListAvailableConfigEmrSupportInstanceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigEmrSupportInstanceType)
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

type ListAvailableConfigSubModuleList []ListAvailableConfigSubModule

func (list *ListAvailableConfigSubModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigSubModule)
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

type ListAvailableConfigRequiredList []ListAvailableConfigRequired

func (list *ListAvailableConfigRequiredList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigRequired)
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

type ListAvailableConfigOptionalList []ListAvailableConfigOptional

func (list *ListAvailableConfigOptionalList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigOptional)
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

type ListAvailableConfigAvailableResourceList []ListAvailableConfigAvailableResource

func (list *ListAvailableConfigAvailableResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigAvailableResource)
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

type ListAvailableConfigAvailableResourceCreationListList []goaliyun.String

func (list *ListAvailableConfigAvailableResourceCreationListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type ListAvailableConfigAvailableDiskCategoryListList []goaliyun.String

func (list *ListAvailableConfigAvailableDiskCategoryListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type ListAvailableConfigSystemDiskCategoryList []goaliyun.String

func (list *ListAvailableConfigSystemDiskCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type ListAvailableConfigDataDiskCategoryList []goaliyun.String

func (list *ListAvailableConfigDataDiskCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type ListAvailableConfigNetworkTypeList []goaliyun.String

func (list *ListAvailableConfigNetworkTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type ListAvailableConfigSupportedInstanceTypeList []goaliyun.String

func (list *ListAvailableConfigSupportedInstanceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type ListAvailableConfigInstanceTypeFamilyList []goaliyun.String

func (list *ListAvailableConfigInstanceTypeFamilyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type ListAvailableConfigInstanceGenerationList []goaliyun.String

func (list *ListAvailableConfigInstanceGenerationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type ListAvailableConfigVSwitchList []ListAvailableConfigVSwitch

func (list *ListAvailableConfigVSwitchList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigVSwitch)
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

type ListAvailableConfigSecurityGroupList []ListAvailableConfigSecurityGroup

func (list *ListAvailableConfigSecurityGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigSecurityGroup)
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

type ListAvailableConfigClusterNodeTypeSupportInfoList []ListAvailableConfigClusterNodeTypeSupportInfo

func (list *ListAvailableConfigClusterNodeTypeSupportInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAvailableConfigClusterNodeTypeSupportInfo)
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

type ListAvailableConfigSupportInstanceTypeListList []goaliyun.String

func (list *ListAvailableConfigSupportInstanceTypeListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
