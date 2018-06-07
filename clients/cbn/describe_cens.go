package cbn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCensRequest struct {
	Filters              *DescribeCensFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                   `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                  `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                  `position:"Query" name:"OwnerAccount"`
	PageSize             int64                   `position:"Query" name:"PageSize"`
	OwnerId              int64                   `position:"Query" name:"OwnerId"`
	PageNumber           int64                   `position:"Query" name:"PageNumber"`
	RegionId             string                  `position:"Query" name:"RegionId"`
}

func (req *DescribeCensRequest) Invoke(client goaliyun.Client) (*DescribeCensResponse, error) {
	resp := &DescribeCensResponse{}
	err := client.Request("cbn", "DescribeCens", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCensFilter struct {
	Key    string                 `name:"Key"`
	Values *DescribeCensValueList `type:"Repeated" name:"Value"`
}

type DescribeCensResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Cens       DescribeCensCenList
}

type DescribeCensCen struct {
	CenId                  goaliyun.String
	Name                   goaliyun.String
	Description            goaliyun.String
	Status                 goaliyun.String
	CreationTime           goaliyun.String
	CenBandwidthPackageIds DescribeCensCenBandwidthPackageIdList
}

type DescribeCensFilterList []DescribeCensFilter

func (list *DescribeCensFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCensFilter)
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

type DescribeCensValueList []string

func (list *DescribeCensValueList) UnmarshalJSON(data []byte) error {
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

type DescribeCensCenList []DescribeCensCen

func (list *DescribeCensCenList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCensCen)
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

type DescribeCensCenBandwidthPackageIdList []goaliyun.String

func (list *DescribeCensCenBandwidthPackageIdList) UnmarshalJSON(data []byte) error {
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
