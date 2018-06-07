package vpc

import (
	"github.com/morlay/goaliyun"
)

type CreateCommonBandwidthPackageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	Bandwidth            int64  `position:"Query" name:"Bandwidth"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Ratio                int64  `position:"Query" name:"Ratio"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateCommonBandwidthPackageRequest) Invoke(client goaliyun.Client) (*CreateCommonBandwidthPackageResponse, error) {
	resp := &CreateCommonBandwidthPackageResponse{}
	err := client.Request("vpc", "CreateCommonBandwidthPackage", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateCommonBandwidthPackageResponse struct {
	RequestId          goaliyun.String
	BandwidthPackageId goaliyun.String
}
