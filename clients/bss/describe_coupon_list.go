package bss

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCouponListRequest struct {
	Status            string `position:"Query" name:"Status"`
	StartDeliveryTime string `position:"Query" name:"StartDeliveryTime"`
	EndDeliveryTime   string `position:"Query" name:"EndDeliveryTime"`
	PageSize          int64  `position:"Query" name:"PageSize"`
	PageNum           int64  `position:"Query" name:"PageNum"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *DescribeCouponListRequest) Invoke(client goaliyun.Client) (*DescribeCouponListResponse, error) {
	resp := &DescribeCouponListResponse{}
	err := client.Request("bss", "DescribeCouponList", "2014-07-14", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCouponListResponse struct {
	RequestId goaliyun.String
	Coupons   DescribeCouponListCouponList
}

type DescribeCouponListCoupon struct {
	CouponTemplateId goaliyun.Integer
	TotalAmount      goaliyun.String
	BalanceAmount    goaliyun.String
	FrozenAmount     goaliyun.String
	ExpiredAmount    goaliyun.String
	DeliveryTime     goaliyun.String
	ExpiredTime      goaliyun.String
	CouponNumber     goaliyun.String
	Status           goaliyun.String
	Description      goaliyun.String
	CreationTime     goaliyun.String
	ModificationTime goaliyun.String
	PriceLimit       goaliyun.String
	Application      goaliyun.String
	ProductCodes     DescribeCouponListProductCodeList
	TradeTypes       DescribeCouponListTradeTypeList
}

type DescribeCouponListCouponList []DescribeCouponListCoupon

func (list *DescribeCouponListCouponList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCouponListCoupon)
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

type DescribeCouponListProductCodeList []goaliyun.String

func (list *DescribeCouponListProductCodeList) UnmarshalJSON(data []byte) error {
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

type DescribeCouponListTradeTypeList []goaliyun.String

func (list *DescribeCouponListTradeTypeList) UnmarshalJSON(data []byte) error {
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
