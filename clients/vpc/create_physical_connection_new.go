package vpc

import (
	"github.com/morlay/goaliyun"
)

type CreatePhysicalConnectionNewRequest struct {
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
	InterfaceName                 string `position:"Query" name:"InterfaceName"`
	Type                          string `position:"Query" name:"Type"`
	OwnerId                       int64  `position:"Query" name:"OwnerId"`
	LineOperator                  string `position:"Query" name:"LineOperator"`
	Name                          string `position:"Query" name:"Name"`
	DeviceName                    string `position:"Query" name:"DeviceName"`
	RegionId                      string `position:"Query" name:"RegionId"`
}

func (req *CreatePhysicalConnectionNewRequest) Invoke(client goaliyun.Client) (*CreatePhysicalConnectionNewResponse, error) {
	resp := &CreatePhysicalConnectionNewResponse{}
	err := client.Request("vpc", "CreatePhysicalConnectionNew", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreatePhysicalConnectionNewResponse struct {
	RequestId            goaliyun.String
	PhysicalConnectionId goaliyun.String
}
