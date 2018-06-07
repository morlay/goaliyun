package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCommonBandwidthPackagesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeCommonBandwidthPackagesRequest) Invoke(client goaliyun.Client) (*DescribeCommonBandwidthPackagesResponse, error) {
	resp := &DescribeCommonBandwidthPackagesResponse{}
	err := client.Request("vpc", "DescribeCommonBandwidthPackages", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCommonBandwidthPackagesResponse struct {
	RequestId               goaliyun.String
	TotalCount              goaliyun.Integer
	PageNumber              goaliyun.Integer
	PageSize                goaliyun.Integer
	CommonBandwidthPackages DescribeCommonBandwidthPackagesCommonBandwidthPackageList
}

type DescribeCommonBandwidthPackagesCommonBandwidthPackage struct {
	BandwidthPackageId goaliyun.String
	RegionId           goaliyun.String
	Name               goaliyun.String
	Description        goaliyun.String
	Bandwidth          goaliyun.String
	InstanceChargeType goaliyun.String
	InternetChargeType goaliyun.String
	BusinessStatus     goaliyun.String
	CreationTime       goaliyun.String
	ExpiredTime        goaliyun.String
	Status             goaliyun.String
	Ratio              goaliyun.Integer
	PublicIpAddresses  DescribeCommonBandwidthPackagesPublicIpAddresseList
}

type DescribeCommonBandwidthPackagesPublicIpAddresse struct {
	AllocationId goaliyun.String
	IpAddress    goaliyun.String
}

type DescribeCommonBandwidthPackagesCommonBandwidthPackageList []DescribeCommonBandwidthPackagesCommonBandwidthPackage

func (list *DescribeCommonBandwidthPackagesCommonBandwidthPackageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCommonBandwidthPackagesCommonBandwidthPackage)
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

type DescribeCommonBandwidthPackagesPublicIpAddresseList []DescribeCommonBandwidthPackagesPublicIpAddresse

func (list *DescribeCommonBandwidthPackagesPublicIpAddresseList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCommonBandwidthPackagesPublicIpAddresse)
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
