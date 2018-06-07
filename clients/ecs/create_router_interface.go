package ecs

import (
	"github.com/morlay/goaliyun"
)

type CreateRouterInterfaceRequest struct {
	AccessPointId            string `position:"Query" name:"AccessPointId"`
	OppositeRouterId         string `position:"Query" name:"OppositeRouterId"`
	OppositeAccessPointId    string `position:"Query" name:"OppositeAccessPointId"`
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	Role                     string `position:"Query" name:"Role"`
	ClientToken              string `position:"Query" name:"ClientToken"`
	HealthCheckTargetIp      string `position:"Query" name:"HealthCheckTargetIp"`
	Description              string `position:"Query" name:"Description"`
	Spec                     string `position:"Query" name:"Spec"`
	UserCidr                 string `position:"Query" name:"UserCidr"`
	OppositeInterfaceId      string `position:"Query" name:"OppositeInterfaceId"`
	InstanceChargeType       string `position:"Query" name:"InstanceChargeType"`
	Period                   int64  `position:"Query" name:"Period"`
	AutoPay                  string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	OppositeRegionId         string `position:"Query" name:"OppositeRegionId"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
	OppositeInterfaceOwnerId string `position:"Query" name:"OppositeInterfaceOwnerId"`
	RouterType               string `position:"Query" name:"RouterType"`
	HealthCheckSourceIp      string `position:"Query" name:"HealthCheckSourceIp"`
	RouterId                 string `position:"Query" name:"RouterId"`
	OppositeRouterType       string `position:"Query" name:"OppositeRouterType"`
	Name                     string `position:"Query" name:"Name"`
	PricingCycle             string `position:"Query" name:"PricingCycle"`
	RegionId                 string `position:"Query" name:"RegionId"`
}

func (req *CreateRouterInterfaceRequest) Invoke(client goaliyun.Client) (*CreateRouterInterfaceResponse, error) {
	resp := &CreateRouterInterfaceResponse{}
	err := client.Request("ecs", "CreateRouterInterface", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateRouterInterfaceResponse struct {
	RequestId         goaliyun.String
	RouterInterfaceId goaliyun.String
	OrderId           goaliyun.Integer
}
