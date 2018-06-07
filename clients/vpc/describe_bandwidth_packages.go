package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeBandwidthPackagesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	NatGatewayId         string `position:"Query" name:"NatGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeBandwidthPackagesRequest) Invoke(client goaliyun.Client) (*DescribeBandwidthPackagesResponse, error) {
	resp := &DescribeBandwidthPackagesResponse{}
	err := client.Request("vpc", "DescribeBandwidthPackages", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeBandwidthPackagesResponse struct {
	RequestId         goaliyun.String
	TotalCount        goaliyun.Integer
	PageNumber        goaliyun.Integer
	PageSize          goaliyun.Integer
	BandwidthPackages DescribeBandwidthPackagesBandwidthPackageList
}

type DescribeBandwidthPackagesBandwidthPackage struct {
	BandwidthPackageId goaliyun.String
	RegionId           goaliyun.String
	Name               goaliyun.String
	Description        goaliyun.String
	ZoneId             goaliyun.String
	NatGatewayId       goaliyun.String
	Bandwidth          goaliyun.String
	InstanceChargeType goaliyun.String
	InternetChargeType goaliyun.String
	BusinessStatus     goaliyun.String
	IpCount            goaliyun.String
	CreationTime       goaliyun.String
	Status             goaliyun.String
	ISP                goaliyun.String
	PublicIpAddresses  DescribeBandwidthPackagesPublicIpAddresseList
}

type DescribeBandwidthPackagesPublicIpAddresse struct {
	AllocationId    goaliyun.String
	IpAddress       goaliyun.String
	UsingStatus     goaliyun.String
	ApAccessEnabled bool
}

type DescribeBandwidthPackagesBandwidthPackageList []DescribeBandwidthPackagesBandwidthPackage

func (list *DescribeBandwidthPackagesBandwidthPackageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBandwidthPackagesBandwidthPackage)
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

type DescribeBandwidthPackagesPublicIpAddresseList []DescribeBandwidthPackagesPublicIpAddresse

func (list *DescribeBandwidthPackagesPublicIpAddresseList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBandwidthPackagesPublicIpAddresse)
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
