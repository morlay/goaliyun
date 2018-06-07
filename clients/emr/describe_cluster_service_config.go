package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeClusterServiceConfigRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	TagValue        string `position:"Query" name:"TagValue"`
	GroupId         int64  `position:"Query" name:"GroupId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ConfigVersion   string `position:"Query" name:"ConfigVersion"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeClusterServiceConfigRequest) Invoke(client goaliyun.Client) (*DescribeClusterServiceConfigResponse, error) {
	resp := &DescribeClusterServiceConfigResponse{}
	err := client.Request("emr", "DescribeClusterServiceConfig", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClusterServiceConfigResponse struct {
	RequestId goaliyun.String
	Config    DescribeClusterServiceConfigConfig
}

type DescribeClusterServiceConfigConfig struct {
	ServiceName      goaliyun.String
	ConfigVersion    goaliyun.String
	Applied          goaliyun.String
	CreateTime       goaliyun.String
	Author           goaliyun.String
	Comment          goaliyun.String
	ConfigValueList  DescribeClusterServiceConfigConfigValueList
	PropertyInfoList DescribeClusterServiceConfigPropertyInfoList
}

type DescribeClusterServiceConfigConfigValue struct {
	ConfigName          goaliyun.String
	AllowCustom         bool
	ConfigItemValueList DescribeClusterServiceConfigConfigItemValueList
}

type DescribeClusterServiceConfigConfigItemValue struct {
	ItemName    goaliyun.String
	Value       goaliyun.String
	IsCustom    bool
	Description goaliyun.String
}

type DescribeClusterServiceConfigPropertyInfo struct {
	Name                    goaliyun.String
	Value                   goaliyun.String
	Description             goaliyun.String
	FileName                goaliyun.String
	DisplayName             goaliyun.String
	ServiceName             goaliyun.String
	Component               goaliyun.String
	PropertyTypes           DescribeClusterServiceConfigPropertyTypeList
	PropertyValueAttributes DescribeClusterServiceConfigPropertyValueAttributes
	EffectWay               DescribeClusterServiceConfigEffectWay
}

type DescribeClusterServiceConfigPropertyValueAttributes struct {
	Type          goaliyun.String
	Maximum       goaliyun.String
	Mimimum       goaliyun.String
	Unit          goaliyun.String
	ReadOnly      bool
	Hidden        bool
	IncrememtStep goaliyun.String
	Entries       DescribeClusterServiceConfigValueEntryInfoList
}

type DescribeClusterServiceConfigValueEntryInfo struct {
	Value       goaliyun.String
	Label       goaliyun.String
	Description goaliyun.String
}

type DescribeClusterServiceConfigEffectWay struct {
	EffectType        goaliyun.String
	InvokeServiceName goaliyun.String
}

type DescribeClusterServiceConfigConfigValueList []DescribeClusterServiceConfigConfigValue

func (list *DescribeClusterServiceConfigConfigValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceConfigConfigValue)
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

type DescribeClusterServiceConfigPropertyInfoList []DescribeClusterServiceConfigPropertyInfo

func (list *DescribeClusterServiceConfigPropertyInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceConfigPropertyInfo)
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

type DescribeClusterServiceConfigConfigItemValueList []DescribeClusterServiceConfigConfigItemValue

func (list *DescribeClusterServiceConfigConfigItemValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceConfigConfigItemValue)
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

type DescribeClusterServiceConfigPropertyTypeList []goaliyun.String

func (list *DescribeClusterServiceConfigPropertyTypeList) UnmarshalJSON(data []byte) error {
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

type DescribeClusterServiceConfigValueEntryInfoList []DescribeClusterServiceConfigValueEntryInfo

func (list *DescribeClusterServiceConfigValueEntryInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceConfigValueEntryInfo)
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
