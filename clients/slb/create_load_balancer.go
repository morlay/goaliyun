package slb

import (
	"github.com/morlay/goaliyun"
)

type CreateLoadBalancerRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	LoadBalancerSpec     string `position:"Query" name:"LoadBalancerSpec"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            int64  `position:"Query" name:"Bandwidth"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MasterZoneId         string `position:"Query" name:"MasterZoneId"`
	Tags                 string `position:"Query" name:"Tags"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	Duration             int64  `position:"Query" name:"Duration"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	LoadBalancerName     string `position:"Query" name:"LoadBalancerName"`
	EnableVpcVipFlow     string `position:"Query" name:"EnableVpcVipFlow"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
	VpcId                string `position:"Query" name:"VpcId"`
	AddressType          string `position:"Query" name:"AddressType"`
	SlaveZoneId          string `position:"Query" name:"SlaveZoneId"`
	PayType              string `position:"Query" name:"PayType"`
	PricingCycle         string `position:"Query" name:"PricingCycle"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateLoadBalancerRequest) Invoke(client goaliyun.Client) (*CreateLoadBalancerResponse, error) {
	resp := &CreateLoadBalancerResponse{}
	err := client.Request("slb", "CreateLoadBalancer", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateLoadBalancerResponse struct {
	RequestId        goaliyun.String
	LoadBalancerId   goaliyun.String
	ResourceGroupId  goaliyun.String
	Address          goaliyun.String
	LoadBalancerName goaliyun.String
	VpcId            goaliyun.String
	VSwitchId        goaliyun.String
	NetworkType      goaliyun.String
	OrderId          goaliyun.Integer
}
