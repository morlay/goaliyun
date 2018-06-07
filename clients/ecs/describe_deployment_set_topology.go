package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDeploymentSetTopologyRequest struct {
	DeploymentSetId      string `position:"Query" name:"DeploymentSetId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Granularity          string `position:"Query" name:"Granularity"`
	Domain               string `position:"Query" name:"Domain"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	DeploymentSetName    string `position:"Query" name:"DeploymentSetName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Strategy             string `position:"Query" name:"Strategy"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDeploymentSetTopologyRequest) Invoke(client goaliyun.Client) (*DescribeDeploymentSetTopologyResponse, error) {
	resp := &DescribeDeploymentSetTopologyResponse{}
	err := client.Request("ecs", "DescribeDeploymentSetTopology", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDeploymentSetTopologyResponse struct {
	RequestId goaliyun.String
	Switchs   DescribeDeploymentSetTopology_SwitchList
	Racks     DescribeDeploymentSetTopologyRackList
}

type DescribeDeploymentSetTopology_Switch struct {
	SwitchId goaliyun.String
	Hosts    DescribeDeploymentSetTopologyHostList
}

type DescribeDeploymentSetTopologyHost struct {
	HostId      goaliyun.String
	InstanceIds DescribeDeploymentSetTopologyInstanceIdList
}

type DescribeDeploymentSetTopologyRack struct {
	RackId goaliyun.String
	Hosts1 DescribeDeploymentSetTopologyHost2List
}

type DescribeDeploymentSetTopologyHost2 struct {
	HostId       goaliyun.String
	InstanceIds3 DescribeDeploymentSetTopologyInstanceIds3List
}

type DescribeDeploymentSetTopology_SwitchList []DescribeDeploymentSetTopology_Switch

func (list *DescribeDeploymentSetTopology_SwitchList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeploymentSetTopology_Switch)
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

type DescribeDeploymentSetTopologyRackList []DescribeDeploymentSetTopologyRack

func (list *DescribeDeploymentSetTopologyRackList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeploymentSetTopologyRack)
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

type DescribeDeploymentSetTopologyHostList []DescribeDeploymentSetTopologyHost

func (list *DescribeDeploymentSetTopologyHostList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeploymentSetTopologyHost)
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

type DescribeDeploymentSetTopologyInstanceIdList []goaliyun.String

func (list *DescribeDeploymentSetTopologyInstanceIdList) UnmarshalJSON(data []byte) error {
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

type DescribeDeploymentSetTopologyHost2List []DescribeDeploymentSetTopologyHost2

func (list *DescribeDeploymentSetTopologyHost2List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeploymentSetTopologyHost2)
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

type DescribeDeploymentSetTopologyInstanceIds3List []goaliyun.String

func (list *DescribeDeploymentSetTopologyInstanceIds3List) UnmarshalJSON(data []byte) error {
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
