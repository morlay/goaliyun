package cbn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCenBandwidthPackagesRequest struct {
	Filters              *DescribeCenBandwidthPackagesFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                                   `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                                  `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                  `position:"Query" name:"OwnerAccount"`
	PageSize             int64                                   `position:"Query" name:"PageSize"`
	OwnerId              int64                                   `position:"Query" name:"OwnerId"`
	PageNumber           int64                                   `position:"Query" name:"PageNumber"`
	IsOrKey              string                                  `position:"Query" name:"IsOrKey"`
	RegionId             string                                  `position:"Query" name:"RegionId"`
}

func (req *DescribeCenBandwidthPackagesRequest) Invoke(client goaliyun.Client) (*DescribeCenBandwidthPackagesResponse, error) {
	resp := &DescribeCenBandwidthPackagesResponse{}
	err := client.Request("cbn", "DescribeCenBandwidthPackages", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCenBandwidthPackagesFilter struct {
	Key    string                                 `name:"Key"`
	Values *DescribeCenBandwidthPackagesValueList `type:"Repeated" name:"Value"`
}

type DescribeCenBandwidthPackagesResponse struct {
	RequestId            goaliyun.String
	TotalCount           goaliyun.Integer
	PageNumber           goaliyun.Integer
	PageSize             goaliyun.Integer
	CenBandwidthPackages DescribeCenBandwidthPackagesCenBandwidthPackageList
}

type DescribeCenBandwidthPackagesCenBandwidthPackage struct {
	CenBandwidthPackageId      goaliyun.String
	Name                       goaliyun.String
	Description                goaliyun.String
	Bandwidth                  goaliyun.Integer
	BandwidthPackageChargeType goaliyun.String
	GeographicRegionAId        goaliyun.String
	GeographicRegionBId        goaliyun.String
	BusinessStatus             goaliyun.String
	CreationTime               goaliyun.String
	ExpiredTime                goaliyun.String
	Status                     goaliyun.String
	CenIds                     DescribeCenBandwidthPackagesCenIdList
}

type DescribeCenBandwidthPackagesFilterList []DescribeCenBandwidthPackagesFilter

func (list *DescribeCenBandwidthPackagesFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCenBandwidthPackagesFilter)
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

type DescribeCenBandwidthPackagesValueList []string

func (list *DescribeCenBandwidthPackagesValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeCenBandwidthPackagesCenBandwidthPackageList []DescribeCenBandwidthPackagesCenBandwidthPackage

func (list *DescribeCenBandwidthPackagesCenBandwidthPackageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCenBandwidthPackagesCenBandwidthPackage)
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

type DescribeCenBandwidthPackagesCenIdList []goaliyun.String

func (list *DescribeCenBandwidthPackagesCenIdList) UnmarshalJSON(data []byte) error {
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
