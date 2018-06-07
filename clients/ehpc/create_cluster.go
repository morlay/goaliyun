package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateClusterRequest struct {
	SccClusterId                string                              `position:"Query" name:"SccClusterId"`
	ImageId                     string                              `position:"Query" name:"ImageId"`
	EcsOrderManagerInstanceType string                              `position:"Query" name:"EcsOrderManagerInstanceType"`
	EhpcVersion                 string                              `position:"Query" name:"EhpcVersion"`
	AccountType                 string                              `position:"Query" name:"AccountType"`
	SecurityGroupId             string                              `position:"Query" name:"SecurityGroupId"`
	Description                 string                              `position:"Query" name:"Description"`
	KeyPairName                 string                              `position:"Query" name:"KeyPairName"`
	SecurityGroupName           string                              `position:"Query" name:"SecurityGroupName"`
	EcsOrderComputeInstanceType string                              `position:"Query" name:"EcsOrderComputeInstanceType"`
	ImageOwnerAlias             string                              `position:"Query" name:"ImageOwnerAlias"`
	VolumeType                  string                              `position:"Query" name:"VolumeType"`
	DeployMode                  string                              `position:"Query" name:"DeployMode"`
	EcsOrderManagerCount        int64                               `position:"Query" name:"EcsOrderManagerCount"`
	Password                    string                              `position:"Query" name:"Password"`
	EcsOrderLoginCount          int64                               `position:"Query" name:"EcsOrderLoginCount"`
	ComputeSpotPriceLimit       string                              `position:"Query" name:"ComputeSpotPriceLimit"`
	AutoRenewPeriod             int64                               `position:"Query" name:"AutoRenewPeriod"`
	Period                      int64                               `position:"Query" name:"Period"`
	VolumeProtocol              string                              `position:"Query" name:"VolumeProtocol"`
	OsTag                       string                              `position:"Query" name:"OsTag"`
	RemoteDirectory             string                              `position:"Query" name:"RemoteDirectory"`
	EcsOrderComputeCount        int64                               `position:"Query" name:"EcsOrderComputeCount"`
	ComputeSpotStrategy         string                              `position:"Query" name:"ComputeSpotStrategy"`
	PostInstallScripts          *CreateClusterPostInstallScriptList `position:"Query" type:"Repeated" name:"PostInstallScript"`
	VSwitchId                   string                              `position:"Query" name:"VSwitchId"`
	PeriodUnit                  string                              `position:"Query" name:"PeriodUnit"`
	Applications                *CreateClusterApplicationList       `position:"Query" type:"Repeated" name:"Application"`
	AutoRenew                   string                              `position:"Query" name:"AutoRenew"`
	EcsChargeType               string                              `position:"Query" name:"EcsChargeType"`
	VpcId                       string                              `position:"Query" name:"VpcId"`
	HaEnable                    string                              `position:"Query" name:"HaEnable"`
	Name                        string                              `position:"Query" name:"Name"`
	SchedulerType               string                              `position:"Query" name:"SchedulerType"`
	VolumeId                    string                              `position:"Query" name:"VolumeId"`
	VolumeMountpoint            string                              `position:"Query" name:"VolumeMountpoint"`
	EcsOrderLoginInstanceType   string                              `position:"Query" name:"EcsOrderLoginInstanceType"`
	ZoneId                      string                              `position:"Query" name:"ZoneId"`
	RegionId                    string                              `position:"Query" name:"RegionId"`
}

func (req *CreateClusterRequest) Invoke(client goaliyun.Client) (*CreateClusterResponse, error) {
	resp := &CreateClusterResponse{}
	err := client.Request("ehpc", "CreateCluster", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateClusterPostInstallScript struct {
	Url  string `name:"Url"`
	Args string `name:"Args"`
}

type CreateClusterApplication struct {
	Tag string `name:"Tag"`
}

type CreateClusterResponse struct {
	RequestId goaliyun.String
	ClusterId goaliyun.String
}

type CreateClusterPostInstallScriptList []CreateClusterPostInstallScript

func (list *CreateClusterPostInstallScriptList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateClusterPostInstallScript)
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

type CreateClusterApplicationList []CreateClusterApplication

func (list *CreateClusterApplicationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateClusterApplication)
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
