package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRenewalPriceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	Period               int64  `position:"Query" name:"Period"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PriceUnit            string `position:"Query" name:"PriceUnit"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceType         string `position:"Query" name:"ResourceType"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeRenewalPriceRequest) Invoke(client goaliyun.Client) (*DescribeRenewalPriceResponse, error) {
	resp := &DescribeRenewalPriceResponse{}
	err := client.Request("ecs", "DescribeRenewalPrice", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRenewalPriceResponse struct {
	RequestId goaliyun.String
	PriceInfo DescribeRenewalPricePriceInfo
}

type DescribeRenewalPricePriceInfo struct {
	Rules DescribeRenewalPriceRuleList
	Price DescribeRenewalPricePrice
}

type DescribeRenewalPriceRule struct {
	RuleId      goaliyun.Integer
	Description goaliyun.String
}

type DescribeRenewalPricePrice struct {
	OriginalPrice goaliyun.Float
	DiscountPrice goaliyun.Float
	TradePrice    goaliyun.Float
	Currency      goaliyun.String
}

type DescribeRenewalPriceRuleList []DescribeRenewalPriceRule

func (list *DescribeRenewalPriceRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRenewalPriceRule)
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
