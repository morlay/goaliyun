package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribePriceRequest struct {
	DataDisk3Size           int64  `position:"Query" name:"DataDisk.3.Size"`
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId                 string `position:"Query" name:"ImageId"`
	DataDisk3Category       string `position:"Query" name:"DataDisk.3.Category"`
	IoOptimized             string `position:"Query" name:"IoOptimized"`
	InternetMaxBandwidthOut int64  `position:"Query" name:"InternetMaxBandwidthOut"`
	SystemDiskCategory      string `position:"Query" name:"SystemDiskCategory"`
	DataDisk4Category       string `position:"Query" name:"DataDisk.4.Category"`
	DataDisk4Size           int64  `position:"Query" name:"DataDisk.4.Size"`
	PriceUnit               string `position:"Query" name:"PriceUnit"`
	InstanceType            string `position:"Query" name:"InstanceType"`
	DataDisk2Category       string `position:"Query" name:"DataDisk.2.Category"`
	DataDisk1Size           int64  `position:"Query" name:"DataDisk.1.Size"`
	Period                  int64  `position:"Query" name:"Period"`
	Amount                  int64  `position:"Query" name:"Amount"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	DataDisk2Size           int64  `position:"Query" name:"DataDisk.2.Size"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	ResourceType            string `position:"Query" name:"ResourceType"`
	DataDisk1Category       string `position:"Query" name:"DataDisk.1.Category"`
	SystemDiskSize          int64  `position:"Query" name:"SystemDiskSize"`
	InternetChargeType      string `position:"Query" name:"InternetChargeType"`
	InstanceNetworkType     string `position:"Query" name:"InstanceNetworkType"`
	RegionId                string `position:"Query" name:"RegionId"`
}

func (req *DescribePriceRequest) Invoke(client goaliyun.Client) (*DescribePriceResponse, error) {
	resp := &DescribePriceResponse{}
	err := client.Request("ecs", "DescribePrice", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribePriceResponse struct {
	RequestId goaliyun.String
	PriceInfo DescribePricePriceInfo
}

type DescribePricePriceInfo struct {
	Rules DescribePriceRuleList
	Price DescribePricePrice
}

type DescribePriceRule struct {
	RuleId      goaliyun.Integer
	Description goaliyun.String
}

type DescribePricePrice struct {
	OriginalPrice goaliyun.Float
	DiscountPrice goaliyun.Float
	TradePrice    goaliyun.Float
	Currency      goaliyun.String
}

type DescribePriceRuleList []DescribePriceRule

func (list *DescribePriceRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePriceRule)
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
