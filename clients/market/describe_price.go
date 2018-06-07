package market

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribePriceRequest struct {
	Commodity string `position:"Query" name:"Commodity"`
	OrderType string `position:"Query" name:"OrderType"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *DescribePriceRequest) Invoke(client goaliyun.Client) (*DescribePriceResponse, error) {
	resp := &DescribePriceResponse{}
	err := client.Request("market", "DescribePrice", "2015-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribePriceResponse struct {
	ProductCode    goaliyun.String
	OriginalPrice  goaliyun.Float
	TradePrice     goaliyun.Float
	DiscountPrice  goaliyun.Float
	PromotionRules DescribePricePromotionRuleList
}

type DescribePricePromotionRule struct {
	RuleId goaliyun.String
	Name   goaliyun.String
	Title  goaliyun.String
}

type DescribePricePromotionRuleList []DescribePricePromotionRule

func (list *DescribePricePromotionRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePricePromotionRule)
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
