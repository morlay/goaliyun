package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateClusterRequest struct {
	ResourceOwnerId        int64                                `position:"Query" name:"ResourceOwnerId"`
	LogPath                string                               `position:"Query" name:"LogPath"`
	MasterPwd              string                               `position:"Query" name:"MasterPwd"`
	Configurations         string                               `position:"Query" name:"Configurations"`
	IoOptimized            string                               `position:"Query" name:"IoOptimized"`
	SecurityGroupId        string                               `position:"Query" name:"SecurityGroupId"`
	SshEnable              string                               `position:"Query" name:"SshEnable"`
	EasEnable              string                               `position:"Query" name:"EasEnable"`
	SecurityGroupName      string                               `position:"Query" name:"SecurityGroupName"`
	DepositType            string                               `position:"Query" name:"DepositType"`
	MachineType            string                               `position:"Query" name:"MachineType"`
	BootstrapActions       *CreateClusterBootstrapActionList    `position:"Query" type:"Repeated" name:"BootstrapAction"`
	UseLocalMetaDb         string                               `position:"Query" name:"UseLocalMetaDb"`
	EmrVer                 string                               `position:"Query" name:"EmrVer"`
	UserDefinedEmrEcsRole  string                               `position:"Query" name:"UserDefinedEmrEcsRole"`
	IsOpenPublicIp         string                               `position:"Query" name:"IsOpenPublicIp"`
	Period                 int64                                `position:"Query" name:"Period"`
	RelatedClusterId       string                               `position:"Query" name:"RelatedClusterId"`
	InstanceGeneration     string                               `position:"Query" name:"InstanceGeneration"`
	VSwitchId              string                               `position:"Query" name:"VSwitchId"`
	ClusterType            string                               `position:"Query" name:"ClusterType"`
	AutoRenew              string                               `position:"Query" name:"AutoRenew"`
	OptionSoftWareLists    *CreateClusterOptionSoftWareListList `position:"Query" type:"Repeated" name:"OptionSoftWareList"`
	VpcId                  string                               `position:"Query" name:"VpcId"`
	NetType                string                               `position:"Query" name:"NetType"`
	EcsOrders              *CreateClusterEcsOrderList           `position:"Query" type:"Repeated" name:"EcsOrder"`
	Name                   string                               `position:"Query" name:"Name"`
	ZoneId                 string                               `position:"Query" name:"ZoneId"`
	ChargeType             string                               `position:"Query" name:"ChargeType"`
	HighAvailabilityEnable string                               `position:"Query" name:"HighAvailabilityEnable"`
	MasterPwdEnable        string                               `position:"Query" name:"MasterPwdEnable"`
	LogEnable              string                               `position:"Query" name:"LogEnable"`
	RegionId               string                               `position:"Query" name:"RegionId"`
}

func (req *CreateClusterRequest) Invoke(client goaliyun.Client) (*CreateClusterResponse, error) {
	resp := &CreateClusterResponse{}
	err := client.Request("emr", "CreateCluster", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateClusterBootstrapAction struct {
	Name string `name:"Name"`
	Path string `name:"Path"`
	Arg  string `name:"Arg"`
}

type CreateClusterEcsOrder struct {
	Index        int64  `name:"Index"`
	NodeCount    int64  `name:"NodeCount"`
	NodeType     string `name:"NodeType"`
	InstanceType string `name:"InstanceType"`
	DiskType     string `name:"DiskType"`
	DiskCapacity int64  `name:"DiskCapacity"`
	DiskCount    int64  `name:"DiskCount"`
}

type CreateClusterResponse struct {
	RequestId     goaliyun.String
	ClusterId     goaliyun.String
	EmrOrderId    goaliyun.String
	MasterOrderId goaliyun.String
	CoreOrderId   goaliyun.String
}

type CreateClusterBootstrapActionList []CreateClusterBootstrapAction

func (list *CreateClusterBootstrapActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateClusterBootstrapAction)
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

type CreateClusterOptionSoftWareListList []string

func (list *CreateClusterOptionSoftWareListList) UnmarshalJSON(data []byte) error {
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

type CreateClusterEcsOrderList []CreateClusterEcsOrder

func (list *CreateClusterEcsOrderList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateClusterEcsOrder)
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
