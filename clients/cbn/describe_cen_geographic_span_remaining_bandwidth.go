package cbn

import (
	"github.com/morlay/goaliyun"
)

type DescribeCenGeographicSpanRemainingBandwidthRequest struct {
	GeographicRegionBId  string `position:"Query" name:"GeographicRegionBId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	GeographicRegionAId  string `position:"Query" name:"GeographicRegionAId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string `position:"Query" name:"CenId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeCenGeographicSpanRemainingBandwidthRequest) Invoke(client goaliyun.Client) (*DescribeCenGeographicSpanRemainingBandwidthResponse, error) {
	resp := &DescribeCenGeographicSpanRemainingBandwidthResponse{}
	err := client.Request("cbn", "DescribeCenGeographicSpanRemainingBandwidth", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCenGeographicSpanRemainingBandwidthResponse struct {
	RequestId          goaliyun.String
	RemainingBandwidth goaliyun.Integer
}
