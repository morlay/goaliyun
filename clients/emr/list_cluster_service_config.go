package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterServiceConfigRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	TagValue        string `position:"Query" name:"TagValue"`
	GroupId         int64  `position:"Query" name:"GroupId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ConfigVersion   string `position:"Query" name:"ConfigVersion"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterServiceConfigRequest) Invoke(client goaliyun.Client) (*ListClusterServiceConfigResponse, error) {
	resp := &ListClusterServiceConfigResponse{}
	err := client.Request("emr", "ListClusterServiceConfig", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterServiceConfigResponse struct {
	RequestId goaliyun.String
	Configs   ListClusterServiceConfigConfigList
}

type ListClusterServiceConfigConfig struct {
	ServiceName      goaliyun.String
	ConfigVersion    goaliyun.String
	Applied          goaliyun.String
	CreateTime       goaliyun.String
	Author           goaliyun.String
	Comment          goaliyun.String
	ConfigValueList  ListClusterServiceConfigConfigValueList
	PropertyInfoList ListClusterServiceConfigPropertyInfoList
}

type ListClusterServiceConfigConfigValue struct {
	ConfigName          goaliyun.String
	AllowCustom         bool
	ConfigItemValueList ListClusterServiceConfigConfigItemValueList
}

type ListClusterServiceConfigConfigItemValue struct {
	ItemName    goaliyun.String
	Value       goaliyun.String
	IsCustom    bool
	Description goaliyun.String
}

type ListClusterServiceConfigPropertyInfo struct {
	Name                    goaliyun.String
	Value                   goaliyun.String
	Description             goaliyun.String
	FileName                goaliyun.String
	DisplayName             goaliyun.String
	ServiceName             goaliyun.String
	Component               goaliyun.String
	PropertyTypes           ListClusterServiceConfigPropertyTypeList
	PropertyValueAttributes ListClusterServiceConfigPropertyValueAttributes
	EffectWay               ListClusterServiceConfigEffectWay
}

type ListClusterServiceConfigPropertyValueAttributes struct {
	Type          goaliyun.String
	Maximum       goaliyun.String
	Mimimum       goaliyun.String
	Unit          goaliyun.String
	ReadOnly      bool
	Hidden        bool
	IncrememtStep goaliyun.String
	Entries       ListClusterServiceConfigValueEntryInfoList
}

type ListClusterServiceConfigValueEntryInfo struct {
	Value       goaliyun.String
	Label       goaliyun.String
	Description goaliyun.String
}

type ListClusterServiceConfigEffectWay struct {
	EffectType        goaliyun.String
	InvokeServiceName goaliyun.String
}

type ListClusterServiceConfigConfigList []ListClusterServiceConfigConfig

func (list *ListClusterServiceConfigConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceConfigConfig)
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

type ListClusterServiceConfigConfigValueList []ListClusterServiceConfigConfigValue

func (list *ListClusterServiceConfigConfigValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceConfigConfigValue)
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

type ListClusterServiceConfigPropertyInfoList []ListClusterServiceConfigPropertyInfo

func (list *ListClusterServiceConfigPropertyInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceConfigPropertyInfo)
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

type ListClusterServiceConfigConfigItemValueList []ListClusterServiceConfigConfigItemValue

func (list *ListClusterServiceConfigConfigItemValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceConfigConfigItemValue)
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

type ListClusterServiceConfigPropertyTypeList []goaliyun.String

func (list *ListClusterServiceConfigPropertyTypeList) UnmarshalJSON(data []byte) error {
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

type ListClusterServiceConfigValueEntryInfoList []ListClusterServiceConfigValueEntryInfo

func (list *ListClusterServiceConfigValueEntryInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceConfigValueEntryInfo)
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
