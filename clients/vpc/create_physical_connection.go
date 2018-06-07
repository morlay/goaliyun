package vpc

import (
	"github.com/morlay/goaliyun"
)

type CreatePhysicalConnectionRequest struct {
	AccessPointId                 string `position:"Query" name:"AccessPointId"`
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
	Type                          string `position:"Query" name:"Type"`
	OwnerId                       int64  `position:"Query" name:"OwnerId"`
	LineOperator                  string `position:"Query" name:"LineOperator"`
	Name                          string `position:"Query" name:"Name"`
	RegionId                      string `position:"Query" name:"RegionId"`
}

func (req *CreatePhysicalConnectionRequest) Invoke(client goaliyun.Client) (*CreatePhysicalConnectionResponse, error) {
	resp := &CreatePhysicalConnectionResponse{}
	err := client.Request("vpc", "CreatePhysicalConnection", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreatePhysicalConnectionResponse struct {
	RequestId            goaliyun.String
	PhysicalConnectionId goaliyun.String
}
