package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSpotPriceHistoryRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	IoOptimized          string `position:"Query" name:"IoOptimized"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	StartTime            string `position:"Query" name:"StartTime"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	Offset               int64  `position:"Query" name:"Offset"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	OSType               string `position:"Query" name:"OSType"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSpotPriceHistoryRequest) Invoke(client goaliyun.Client) (*DescribeSpotPriceHistoryResponse, error) {
	resp := &DescribeSpotPriceHistoryResponse{}
	err := client.Request("ecs", "DescribeSpotPriceHistory", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSpotPriceHistoryResponse struct {
	RequestId  goaliyun.String
	NextOffset goaliyun.Integer
	Currency   goaliyun.String
	SpotPrices DescribeSpotPriceHistorySpotPriceTypeList
}

type DescribeSpotPriceHistorySpotPriceType struct {
	ZoneId       goaliyun.String
	InstanceType goaliyun.String
	IoOptimized  goaliyun.String
	Timestamp    goaliyun.String
	NetworkType  goaliyun.String
	SpotPrice    goaliyun.Float
	OriginPrice  goaliyun.Float
}

type DescribeSpotPriceHistorySpotPriceTypeList []DescribeSpotPriceHistorySpotPriceType

func (list *DescribeSpotPriceHistorySpotPriceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSpotPriceHistorySpotPriceType)
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
