package bss

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCouponDetailRequest struct {
	CouponNumber string `position:"Query" name:"CouponNumber"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeCouponDetailRequest) Invoke(client goaliyun.Client) (*DescribeCouponDetailResponse, error) {
	resp := &DescribeCouponDetailResponse{}
	err := client.Request("bss", "DescribeCouponDetail", "2014-07-14", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCouponDetailResponse struct {
	RequestId        goaliyun.String
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
	ProductCodes     DescribeCouponDetailProductCodeList
	TradeTypes       DescribeCouponDetailTradeTypeList
}

type DescribeCouponDetailProductCodeList []goaliyun.String

func (list *DescribeCouponDetailProductCodeList) UnmarshalJSON(data []byte) error {
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

type DescribeCouponDetailTradeTypeList []goaliyun.String

func (list *DescribeCouponDetailTradeTypeList) UnmarshalJSON(data []byte) error {
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
