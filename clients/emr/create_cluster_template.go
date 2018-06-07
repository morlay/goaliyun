package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateClusterTemplateRequest struct {
	ResourceOwnerId        int64                                        `position:"Query" name:"ResourceOwnerId"`
	LogPath                string                                       `position:"Query" name:"LogPath"`
	MasterPwd              string                                       `position:"Query" name:"MasterPwd"`
	Configurations         string                                       `position:"Query" name:"Configurations"`
	IoOptimized            string                                       `position:"Query" name:"IoOptimized"`
	SecurityGroupId        string                                       `position:"Query" name:"SecurityGroupId"`
	SshEnable              string                                       `position:"Query" name:"SshEnable"`
	EasEnable              string                                       `position:"Query" name:"EasEnable"`
	SecurityGroupName      string                                       `position:"Query" name:"SecurityGroupName"`
	BootstrapActions       *CreateClusterTemplateBootstrapActionList    `position:"Query" type:"Repeated" name:"BootstrapAction"`
	UseLocalMetaDb         string                                       `position:"Query" name:"UseLocalMetaDb"`
	EmrVer                 string                                       `position:"Query" name:"EmrVer"`
	TemplateName           string                                       `position:"Query" name:"TemplateName"`
	UserDefinedEmrEcsRole  string                                       `position:"Query" name:"UserDefinedEmrEcsRole"`
	IsOpenPublicIp         string                                       `position:"Query" name:"IsOpenPublicIp"`
	Period                 int64                                        `position:"Query" name:"Period"`
	InstanceGeneration     string                                       `position:"Query" name:"InstanceGeneration"`
	VSwitchId              string                                       `position:"Query" name:"VSwitchId"`
	ClusterType            string                                       `position:"Query" name:"ClusterType"`
	AutoRenew              string                                       `position:"Query" name:"AutoRenew"`
	OptionSoftWareLists    *CreateClusterTemplateOptionSoftWareListList `position:"Query" type:"Repeated" name:"OptionSoftWareList"`
	VpcId                  string                                       `position:"Query" name:"VpcId"`
	NetType                string                                       `position:"Query" name:"NetType"`
	HostGroups             *CreateClusterTemplateHostGroupList          `position:"Query" type:"Repeated" name:"HostGroup"`
	ZoneId                 string                                       `position:"Query" name:"ZoneId"`
	HighAvailabilityEnable string                                       `position:"Query" name:"HighAvailabilityEnable"`
	RegionId               string                                       `position:"Query" name:"RegionId"`
}

func (req *CreateClusterTemplateRequest) Invoke(client goaliyun.Client) (*CreateClusterTemplateResponse, error) {
	resp := &CreateClusterTemplateResponse{}
	err := client.Request("emr", "CreateClusterTemplate", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateClusterTemplateBootstrapAction struct {
	Name string `name:"Name"`
	Path string `name:"Path"`
	Arg  string `name:"Arg"`
}

type CreateClusterTemplateHostGroup struct {
	HostGroupId     string `name:"HostGroupId"`
	HostGroupName   string `name:"HostGroupName"`
	HostGroupType   string `name:"HostGroupType"`
	ClusterId       string `name:"ClusterId"`
	Comment         string `name:"Comment"`
	CreateType      string `name:"CreateType"`
	ChargeType      string `name:"ChargeType"`
	Period          int64  `name:"Period"`
	NodeCount       int64  `name:"NodeCount"`
	InstanceType    string `name:"InstanceType"`
	DiskType        string `name:"DiskType"`
	DiskCapacity    int64  `name:"DiskCapacity"`
	DiskCount       int64  `name:"DiskCount"`
	SysDiskCapacity int64  `name:"SysDiskCapacity"`
	AutoRenew       string `name:"AutoRenew"`
	VSwitchId       string `name:"VSwitchId"`
}

type CreateClusterTemplateResponse struct {
	RequestId         goaliyun.String
	Success           goaliyun.String
	ErrCode           goaliyun.String
	ErrMsg            goaliyun.String
	ClusterTemplateId goaliyun.String
}

type CreateClusterTemplateBootstrapActionList []CreateClusterTemplateBootstrapAction

func (list *CreateClusterTemplateBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateClusterTemplateBootstrapAction)
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

type CreateClusterTemplateOptionSoftWareListList []string

func (list *CreateClusterTemplateOptionSoftWareListList) UnmarshalJSON(data []byte) error {
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

type CreateClusterTemplateHostGroupList []CreateClusterTemplateHostGroup

func (list *CreateClusterTemplateHostGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateClusterTemplateHostGroup)
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
