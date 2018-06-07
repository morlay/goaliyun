package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyPhysicalConnectionAttributeRequest struct {
	RedundantPhysicalConnectionId string `position:"Query" name:"RedundantPhysicalConnectionId"`
	PeerLocation                  string `position:"Query" name:"PeerLocation"`
	ResourceOwnerId               int64  `position:"Query" name:"ResourceOwnerId"`
	PortType                      string `position:"Query" name:"PortType"`
	CircuitCode                   string `position:"Query" name:"CircuitCode"`
	Bandwidth                     int64  `position:"Query" name:"Bandwidth"`
	ClientToken                   string `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount          string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                  string `position:"Query" name:"OwnerAccount"`
	Description                   string `position:"Query" name:"Description"`
	OwnerId                       int64  `position:"Query" name:"OwnerId"`
	LineOperator                  string `position:"Query" name:"LineOperator"`
	PhysicalConnectionId          string `position:"Query" name:"PhysicalConnectionId"`
	Name                          string `position:"Query" name:"Name"`
	UserCidr                      string `position:"Query" name:"UserCidr"`
	RegionId                      string `position:"Query" name:"RegionId"`
}

func (req *ModifyPhysicalConnectionAttributeRequest) Invoke(client goaliyun.Client) (*ModifyPhysicalConnectionAttributeResponse, error) {
	resp := &ModifyPhysicalConnectionAttributeResponse{}
	err := client.Request("ecs", "ModifyPhysicalConnectionAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyPhysicalConnectionAttributeResponse struct {
	RequestId goaliyun.String
}
