package vpc

import (
	"github.com/morlay/goaliyun"
)

type CreateBandwidthPackageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	Bandwidth            int64  `position:"Query" name:"Bandwidth"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ISP                  string `position:"Query" name:"ISP"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Zone                 string `position:"Query" name:"Zone"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
	Name                 string `position:"Query" name:"Name"`
	NatGatewayId         string `position:"Query" name:"NatGatewayId"`
	IpCount              int64  `position:"Query" name:"IpCount"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateBandwidthPackageRequest) Invoke(client goaliyun.Client) (*CreateBandwidthPackageResponse, error) {
	resp := &CreateBandwidthPackageResponse{}
	err := client.Request("vpc", "CreateBandwidthPackage", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateBandwidthPackageResponse struct {
	RequestId          goaliyun.String
	BandwidthPackageId goaliyun.String
}
