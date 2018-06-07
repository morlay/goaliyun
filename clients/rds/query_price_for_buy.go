package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryPriceForBuyRequest struct {
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

func (req *QueryPriceForBuyRequest) Invoke(client goaliyun.Client) (*QueryPriceForBuyResponse, error) {
	resp := &QueryPriceForBuyResponse{}
	err := client.Request("rds", "QueryPriceForBuy", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryPriceForBuyResponse struct {
	RequestId goaliyun.String
	Rules     QueryPriceForBuyRuleList
	PriceInfo QueryPriceForBuyPriceInfo
}

type QueryPriceForBuyRule struct {
	RuleId      goaliyun.Integer
	Name        goaliyun.String
	Description goaliyun.String
}

type QueryPriceForBuyPriceInfo struct {
	Currency      goaliyun.String
	OriginalPrice goaliyun.Float
	TradePrice    goaliyun.Float
	DiscountPrice goaliyun.Float
	Coupons       QueryPriceForBuyCouponList
	RuleIds       QueryPriceForBuyRuleIdList
	ActivityInfo  QueryPriceForBuyActivityInfo
}

type QueryPriceForBuyCoupon struct {
	CouponNo    goaliyun.String
	Name        goaliyun.String
	Description goaliyun.String
	IsSelected  goaliyun.String
}

type QueryPriceForBuyActivityInfo struct {
	CheckErrMsg goaliyun.String
	ErrorCode   goaliyun.String
	Success     goaliyun.String
}

type QueryPriceForBuyRuleList []QueryPriceForBuyRule

func (list *QueryPriceForBuyRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPriceForBuyRule)
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

type QueryPriceForBuyCouponList []QueryPriceForBuyCoupon

func (list *QueryPriceForBuyCouponList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPriceForBuyCoupon)
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

type QueryPriceForBuyRuleIdList []goaliyun.String

func (list *QueryPriceForBuyRuleIdList) UnmarshalJSON(data []byte) error {
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
