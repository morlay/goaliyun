package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateNatGatewayRequest struct {
	ResourceOwnerId      int64                                 `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                                `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string                                `position:"Query" name:"ClientToken"`
	OwnerAccount         string                                `position:"Query" name:"OwnerAccount"`
	VpcId                string                                `position:"Query" name:"VpcId"`
	Name                 string                                `position:"Query" name:"Name"`
	Description          string                                `position:"Query" name:"Description"`
	OwnerId              int64                                 `position:"Query" name:"OwnerId"`
	BandwidthPackages    *CreateNatGatewayBandwidthPackageList `position:"Query" type:"Repeated" name:"BandwidthPackage"`
	RegionId             string                                `position:"Query" name:"RegionId"`
}

func (req *CreateNatGatewayRequest) Invoke(client goaliyun.Client) (*CreateNatGatewayResponse, error) {
	resp := &CreateNatGatewayResponse{}
	err := client.Request("ecs", "CreateNatGateway", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateNatGatewayBandwidthPackage struct {
	IpCount   int64  `name:"IpCount"`
	Bandwidth int64  `name:"Bandwidth"`
	Zone      string `name:"Zone"`
}

type CreateNatGatewayResponse struct {
	RequestId           goaliyun.String
	NatGatewayId        goaliyun.String
	ForwardTableIds     CreateNatGatewayForwardTableIdList
	BandwidthPackageIds CreateNatGatewayBandwidthPackageIdList
}

type CreateNatGatewayBandwidthPackageList []CreateNatGatewayBandwidthPackage

func (list *CreateNatGatewayBandwidthPackageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateNatGatewayBandwidthPackage)
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

type CreateNatGatewayForwardTableIdList []goaliyun.String

func (list *CreateNatGatewayForwardTableIdList) UnmarshalJSON(data []byte) error {
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

type CreateNatGatewayBandwidthPackageIdList []goaliyun.String

func (list *CreateNatGatewayBandwidthPackageIdList) UnmarshalJSON(data []byte) error {
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
