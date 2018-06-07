package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribePriceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage    int64  `position:"Query" name:"DBInstanceStorage"`
	Quantity             int64  `position:"Query" name:"Quantity"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CommodityCode        string `position:"Query" name:"CommodityCode"`
	EngineVersion        string `position:"Query" name:"EngineVersion"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	UsedTime             string `position:"Query" name:"UsedTime"`
	DBInstanceClass      string `position:"Query" name:"DBInstanceClass"`
	InstanceUsedType     int64  `position:"Query" name:"InstanceUsedType"`
	Engine               string `position:"Query" name:"Engine"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	TimeType             string `position:"Query" name:"TimeType"`
	PayType              string `position:"Query" name:"PayType"`
	OrderType            string `position:"Query" name:"OrderType"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribePriceRequest) Invoke(client goaliyun.Client) (*DescribePriceResponse, error) {
	resp := &DescribePriceResponse{}
	err := client.Request("rds", "DescribePrice", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribePriceResponse struct {
	RequestId goaliyun.String
	Rules     DescribePriceRuleList
	PriceInfo DescribePricePriceInfo
}

type DescribePriceRule struct {
	RuleId      goaliyun.Integer
	Name        goaliyun.String
	Description goaliyun.String
}

type DescribePricePriceInfo struct {
	Currency      goaliyun.String
	OriginalPrice goaliyun.Float
	TradePrice    goaliyun.Float
	DiscountPrice goaliyun.Float
	Coupons       DescribePriceCouponList
	RuleIds       DescribePriceRuleIdList
	ActivityInfo  DescribePriceActivityInfo
}

type DescribePriceCoupon struct {
	CouponNo    goaliyun.String
	Name        goaliyun.String
	Description goaliyun.String
	IsSelected  goaliyun.String
}

type DescribePriceActivityInfo struct {
	CheckErrMsg goaliyun.String
	ErrorCode   goaliyun.String
	Success     goaliyun.String
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

type DescribePriceCouponList []DescribePriceCoupon

func (list *DescribePriceCouponList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePriceCoupon)
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

type DescribePriceRuleIdList []goaliyun.String

func (list *DescribePriceRuleIdList) UnmarshalJSON(data []byte) error {
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
