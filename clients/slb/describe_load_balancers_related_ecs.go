package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLoadBalancersRelatedEcsRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeLoadBalancersRelatedEcsRequest) Invoke(client goaliyun.Client) (*DescribeLoadBalancersRelatedEcsResponse, error) {
	resp := &DescribeLoadBalancersRelatedEcsResponse{}
	err := client.Request("slb", "DescribeLoadBalancersRelatedEcs", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLoadBalancersRelatedEcsResponse struct {
	Message       goaliyun.String
	Success       bool
	RequestId     goaliyun.String
	LoadBalancers DescribeLoadBalancersRelatedEcsLoadBalancerList
}

type DescribeLoadBalancersRelatedEcsLoadBalancer struct {
	LoadBalancerId           goaliyun.String
	Count                    goaliyun.Integer
	MasterSlaveVServerGroups DescribeLoadBalancersRelatedEcsMasterSlaveVServerGroupList
	VServerGroups            DescribeLoadBalancersRelatedEcsVServerGroupList
	BackendServers           DescribeLoadBalancersRelatedEcsBackendServer4List
}

type DescribeLoadBalancersRelatedEcsMasterSlaveVServerGroup struct {
	GroupId         goaliyun.String
	GroupName       goaliyun.String
	BackendServers1 DescribeLoadBalancersRelatedEcsBackendServerList
}

type DescribeLoadBalancersRelatedEcsBackendServer struct {
	VmName      goaliyun.String
	Weight      goaliyun.Integer
	Port        goaliyun.Integer
	NetworkType goaliyun.String
}

type DescribeLoadBalancersRelatedEcsVServerGroup struct {
	GroupId         goaliyun.String
	GroupName       goaliyun.String
	BackendServers2 DescribeLoadBalancersRelatedEcsBackendServer3List
}

type DescribeLoadBalancersRelatedEcsBackendServer3 struct {
	VmName      goaliyun.String
	Weight      goaliyun.Integer
	Port        goaliyun.Integer
	NetworkType goaliyun.String
}

type DescribeLoadBalancersRelatedEcsBackendServer4 struct {
	VmName      goaliyun.String
	Weight      goaliyun.Integer
	Port        goaliyun.Integer
	NetworkType goaliyun.String
}

type DescribeLoadBalancersRelatedEcsLoadBalancerList []DescribeLoadBalancersRelatedEcsLoadBalancer

func (list *DescribeLoadBalancersRelatedEcsLoadBalancerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersRelatedEcsLoadBalancer)
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

type DescribeLoadBalancersRelatedEcsMasterSlaveVServerGroupList []DescribeLoadBalancersRelatedEcsMasterSlaveVServerGroup

func (list *DescribeLoadBalancersRelatedEcsMasterSlaveVServerGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersRelatedEcsMasterSlaveVServerGroup)
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

type DescribeLoadBalancersRelatedEcsVServerGroupList []DescribeLoadBalancersRelatedEcsVServerGroup

func (list *DescribeLoadBalancersRelatedEcsVServerGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersRelatedEcsVServerGroup)
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

type DescribeLoadBalancersRelatedEcsBackendServer4List []DescribeLoadBalancersRelatedEcsBackendServer4

func (list *DescribeLoadBalancersRelatedEcsBackendServer4List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersRelatedEcsBackendServer4)
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

type DescribeLoadBalancersRelatedEcsBackendServerList []DescribeLoadBalancersRelatedEcsBackendServer

func (list *DescribeLoadBalancersRelatedEcsBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersRelatedEcsBackendServer)
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

type DescribeLoadBalancersRelatedEcsBackendServer3List []DescribeLoadBalancersRelatedEcsBackendServer3

func (list *DescribeLoadBalancersRelatedEcsBackendServer3List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancersRelatedEcsBackendServer3)
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
