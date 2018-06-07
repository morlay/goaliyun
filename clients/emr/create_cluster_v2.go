package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateClusterV2Request struct {
	ResourceOwnerId        int64                                  `position:"Query" name:"ResourceOwnerId"`
	LogPath                string                                 `position:"Query" name:"LogPath"`
	MasterPwd              string                                 `position:"Query" name:"MasterPwd"`
	Configurations         string                                 `position:"Query" name:"Configurations"`
	IoOptimized            string                                 `position:"Query" name:"IoOptimized"`
	SecurityGroupId        string                                 `position:"Query" name:"SecurityGroupId"`
	SshEnable              string                                 `position:"Query" name:"SshEnable"`
	EasEnable              string                                 `position:"Query" name:"EasEnable"`
	SecurityGroupName      string                                 `position:"Query" name:"SecurityGroupName"`
	DepositType            string                                 `position:"Query" name:"DepositType"`
	MachineType            string                                 `position:"Query" name:"MachineType"`
	BootstrapActions       *CreateClusterV2BootstrapActionList    `position:"Query" type:"Repeated" name:"BootstrapAction"`
	UseLocalMetaDb         string                                 `position:"Query" name:"UseLocalMetaDb"`
	EmrVer                 string                                 `position:"Query" name:"EmrVer"`
	UserDefinedEmrEcsRole  string                                 `position:"Query" name:"UserDefinedEmrEcsRole"`
	IsOpenPublicIp         string                                 `position:"Query" name:"IsOpenPublicIp"`
	Period                 int64                                  `position:"Query" name:"Period"`
	RelatedClusterId       string                                 `position:"Query" name:"RelatedClusterId"`
	InstanceGeneration     string                                 `position:"Query" name:"InstanceGeneration"`
	VSwitchId              string                                 `position:"Query" name:"VSwitchId"`
	ClusterType            string                                 `position:"Query" name:"ClusterType"`
	AutoRenew              string                                 `position:"Query" name:"AutoRenew"`
	OptionSoftWareLists    *CreateClusterV2OptionSoftWareListList `position:"Query" type:"Repeated" name:"OptionSoftWareList"`
	VpcId                  string                                 `position:"Query" name:"VpcId"`
	NetType                string                                 `position:"Query" name:"NetType"`
	Name                   string                                 `position:"Query" name:"Name"`
	HostGroups             *CreateClusterV2HostGroupList          `position:"Query" type:"Repeated" name:"HostGroup"`
	ZoneId                 string                                 `position:"Query" name:"ZoneId"`
	ChargeType             string                                 `position:"Query" name:"ChargeType"`
	HighAvailabilityEnable string                                 `position:"Query" name:"HighAvailabilityEnable"`
	RegionId               string                                 `position:"Query" name:"RegionId"`
}

func (req *CreateClusterV2Request) Invoke(client goaliyun.Client) (*CreateClusterV2Response, error) {
	resp := &CreateClusterV2Response{}
	err := client.Request("emr", "CreateClusterV2", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateClusterV2BootstrapAction struct {
	Name string `name:"Name"`
	Path string `name:"Path"`
	Arg  string `name:"Arg"`
}

type CreateClusterV2HostGroup struct {
	ClusterId       string `name:"ClusterId"`
	HostGroupId     string `name:"HostGroupId"`
	HostGroupName   string `name:"HostGroupName"`
	HostGroupType   string `name:"HostGroupType"`
	Comment         string `name:"Comment"`
	CreateType      string `name:"CreateType"`
	ChargeType      string `name:"ChargeType"`
	Period          int64  `name:"Period"`
	NodeCount       int64  `name:"NodeCount"`
	InstanceType    string `name:"InstanceType"`
	DiskType        string `name:"DiskType"`
	DiskCapacity    int64  `name:"DiskCapacity"`
	DiskCount       int64  `name:"DiskCount"`
	SysDiskType     string `name:"SysDiskType"`
	SysDiskCapacity int64  `name:"SysDiskCapacity"`
	AutoRenew       string `name:"AutoRenew"`
	VSwitchId       string `name:"VSwitchId"`
}

type CreateClusterV2Response struct {
	RequestId     goaliyun.String
	ClusterId     goaliyun.String
	EmrOrderId    goaliyun.String
	MasterOrderId goaliyun.String
	CoreOrderId   goaliyun.String
}

type CreateClusterV2BootstrapActionList []CreateClusterV2BootstrapAction

func (list *CreateClusterV2BootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateClusterV2BootstrapAction)
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

type CreateClusterV2OptionSoftWareListList []string

func (list *CreateClusterV2OptionSoftWareListList) UnmarshalJSON(data []byte) error {
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

type CreateClusterV2HostGroupList []CreateClusterV2HostGroup

func (list *CreateClusterV2HostGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateClusterV2HostGroup)
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
