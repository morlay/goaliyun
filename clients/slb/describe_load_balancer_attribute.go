package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLoadBalancerAttributeRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeLoadBalancerAttributeRequest) Invoke(client goaliyun.Client) (*DescribeLoadBalancerAttributeResponse, error) {
	resp := &DescribeLoadBalancerAttributeResponse{}
	err := client.Request("slb", "DescribeLoadBalancerAttribute", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLoadBalancerAttributeResponse struct {
	RequestId                goaliyun.String
	LoadBalancerId           goaliyun.String
	ResourceGroupId          goaliyun.String
	LoadBalancerName         goaliyun.String
	LoadBalancerStatus       goaliyun.String
	RegionId                 goaliyun.String
	RegionIdAlias            goaliyun.String
	Address                  goaliyun.String
	AddressType              goaliyun.String
	VpcId                    goaliyun.String
	VSwitchId                goaliyun.String
	NetworkType              goaliyun.String
	InternetChargeType       goaliyun.String
	AutoReleaseTime          goaliyun.Integer
	Bandwidth                goaliyun.Integer
	LoadBalancerSpec         goaliyun.String
	CreateTime               goaliyun.String
	CreateTimeStamp          goaliyun.Integer
	EndTime                  goaliyun.String
	EndTimeStamp             goaliyun.Integer
	PayType                  goaliyun.String
	MasterZoneId             goaliyun.String
	SlaveZoneId              goaliyun.String
	ListenerPortsAndProtocal DescribeLoadBalancerAttributeListenerPortAndProtocalList
	ListenerPortsAndProtocol DescribeLoadBalancerAttributeListenerPortAndProtocolList
	BackendServers           DescribeLoadBalancerAttributeBackendServerList
	ListenerPorts            DescribeLoadBalancerAttributeListenerPortList
}

type DescribeLoadBalancerAttributeListenerPortAndProtocal struct {
	ListenerPort     goaliyun.Integer
	ListenerProtocal goaliyun.String
}

type DescribeLoadBalancerAttributeListenerPortAndProtocol struct {
	ListenerPort     goaliyun.Integer
	ListenerProtocol goaliyun.String
	ListenerForward  goaliyun.String
	ForwardPort      goaliyun.Integer
}

type DescribeLoadBalancerAttributeBackendServer struct {
	ServerId goaliyun.String
	Weight   goaliyun.Integer
	Type     goaliyun.String
	ServerIp goaliyun.String
	VpcId    goaliyun.String
}

type DescribeLoadBalancerAttributeListenerPortAndProtocalList []DescribeLoadBalancerAttributeListenerPortAndProtocal

func (list *DescribeLoadBalancerAttributeListenerPortAndProtocalList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancerAttributeListenerPortAndProtocal)
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

type DescribeLoadBalancerAttributeListenerPortAndProtocolList []DescribeLoadBalancerAttributeListenerPortAndProtocol

func (list *DescribeLoadBalancerAttributeListenerPortAndProtocolList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancerAttributeListenerPortAndProtocol)
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

type DescribeLoadBalancerAttributeBackendServerList []DescribeLoadBalancerAttributeBackendServer

func (list *DescribeLoadBalancerAttributeBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLoadBalancerAttributeBackendServer)
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

type DescribeLoadBalancerAttributeListenerPortList []goaliyun.String

func (list *DescribeLoadBalancerAttributeListenerPortList) UnmarshalJSON(data []byte) error {
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
