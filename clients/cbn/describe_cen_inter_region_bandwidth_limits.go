package cbn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCenInterRegionBandwidthLimitsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string `position:"Query" name:"CenId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeCenInterRegionBandwidthLimitsRequest) Invoke(client goaliyun.Client) (*DescribeCenInterRegionBandwidthLimitsResponse, error) {
	resp := &DescribeCenInterRegionBandwidthLimitsResponse{}
	err := client.Request("cbn", "DescribeCenInterRegionBandwidthLimits", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCenInterRegionBandwidthLimitsResponse struct {
	RequestId                     goaliyun.String
	TotalCount                    goaliyun.Integer
	PageNumber                    goaliyun.Integer
	PageSize                      goaliyun.Integer
	CenInterRegionBandwidthLimits DescribeCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimitList
}

type DescribeCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit struct {
	CenId            goaliyun.String
	LocalRegionId    goaliyun.String
	OppositeRegionId goaliyun.String
	BandwidthLimit   goaliyun.Integer
	Status           goaliyun.String
}

type DescribeCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimitList []DescribeCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit

func (list *DescribeCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimitList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCenInterRegionBandwidthLimitsCenInterRegionBandwidthLimit)
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
