package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRenewalPriceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Quantity             int64  `position:"Query" name:"Quantity"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CommodityCode        string `position:"Query" name:"CommodityCode"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	UsedTime             string `position:"Query" name:"UsedTime"`
	DBInstanceClass      string `position:"Query" name:"DBInstanceClass"`
	PromotionCode        string `position:"Query" name:"PromotionCode"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	TimeType             string `position:"Query" name:"TimeType"`
	PayType              string `position:"Query" name:"PayType"`
	BusinessInfo         string `position:"Query" name:"BusinessInfo"`
	OrderType            string `position:"Query" name:"OrderType"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeRenewalPriceRequest) Invoke(client goaliyun.Client) (*DescribeRenewalPriceResponse, error) {
	resp := &DescribeRenewalPriceResponse{}
	err := client.Request("rds", "DescribeRenewalPrice", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRenewalPriceResponse struct {
	RequestId goaliyun.String
	Rules     DescribeRenewalPriceRuleList
	PriceInfo DescribeRenewalPricePriceInfo
}

type DescribeRenewalPriceRule struct {
	RuleId      goaliyun.Integer
	Name        goaliyun.String
	Description goaliyun.String
}

type DescribeRenewalPricePriceInfo struct {
	Currency      goaliyun.String
	OriginalPrice goaliyun.Float
	TradePrice    goaliyun.Float
	DiscountPrice goaliyun.Float
	Coupons       DescribeRenewalPriceCouponList
	RuleIds       DescribeRenewalPriceRuleIdList
	ActivityInfo  DescribeRenewalPriceActivityInfo
}

type DescribeRenewalPriceCoupon struct {
	CouponNo    goaliyun.String
	Name        goaliyun.String
	Description goaliyun.String
	IsSelected  goaliyun.String
}

type DescribeRenewalPriceActivityInfo struct {
	CheckErrMsg goaliyun.String
	ErrorCode   goaliyun.String
	Success     goaliyun.String
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

type DescribeRenewalPriceCouponList []DescribeRenewalPriceCoupon

func (list *DescribeRenewalPriceCouponList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRenewalPriceCoupon)
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

type DescribeRenewalPriceRuleIdList []goaliyun.String

func (list *DescribeRenewalPriceRuleIdList) UnmarshalJSON(data []byte) error {
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
