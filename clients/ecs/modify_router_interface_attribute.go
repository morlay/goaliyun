package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyRouterInterfaceAttributeRequest struct {
	OppositeRouterId         string `position:"Query" name:"OppositeRouterId"`
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	Description              string `position:"Query" name:"Description"`
	HealthCheckTargetIp      string `position:"Query" name:"HealthCheckTargetIp"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
	RouterInterfaceId        string `position:"Query" name:"RouterInterfaceId"`
	OppositeInterfaceOwnerId int64  `position:"Query" name:"OppositeInterfaceOwnerId"`
	HealthCheckSourceIp      string `position:"Query" name:"HealthCheckSourceIp"`
	Name                     string `position:"Query" name:"Name"`
	OppositeRouterType       string `position:"Query" name:"OppositeRouterType"`
	OppositeInterfaceId      string `position:"Query" name:"OppositeInterfaceId"`
	RegionId                 string `position:"Query" name:"RegionId"`
}

func (req *ModifyRouterInterfaceAttributeRequest) Invoke(client goaliyun.Client) (*ModifyRouterInterfaceAttributeResponse, error) {
	resp := &ModifyRouterInterfaceAttributeResponse{}
	err := client.Request("ecs", "ModifyRouterInterfaceAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyRouterInterfaceAttributeResponse struct {
	RequestId goaliyun.String
}
