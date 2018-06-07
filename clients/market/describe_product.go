package market

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeProductRequest struct {
	Code       string `position:"Query" name:"Code"`
	QueryDraft string `position:"Query" name:"QueryDraft"`
	AliUid     string `position:"Query" name:"AliUid"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DescribeProductRequest) Invoke(client goaliyun.Client) (*DescribeProductResponse, error) {
	resp := &DescribeProductResponse{}
	err := client.Request("market", "DescribeProduct", "2015-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeProductResponse struct {
	Code             goaliyun.String
	Name             goaliyun.String
	Type             goaliyun.String
	PicUrl           goaliyun.String
	Description      goaliyun.String
	ShortDescription goaliyun.String
	UseCount         goaliyun.Integer
	Score            goaliyun.Float
	Status           goaliyun.String
	AuditStatus      goaliyun.String
	AuditFailMsg     goaliyun.String
	AuditTime        goaliyun.Integer
	GmtCreated       goaliyun.Integer
	GmtModified      goaliyun.Integer
	ProductSkus      DescribeProductProductSkuList
	ProductExtras    DescribeProductProductExtraList
	ShopInfo         DescribeProductShopInfo
}

type DescribeProductProductSku struct {
	Name         goaliyun.String
	Code         goaliyun.String
	ChargeType   goaliyun.String
	Constraints  goaliyun.String
	Hidden       bool
	OrderPeriods DescribeProductOrderPeriodList
	Modules      DescribeProductModuleList
}

type DescribeProductOrderPeriod struct {
	Name       goaliyun.String
	PeriodType goaliyun.String
}

type DescribeProductModule struct {
	Id         goaliyun.String
	Name       goaliyun.String
	Code       goaliyun.String
	Properties DescribeProductPropertyList
}

type DescribeProductProperty struct {
	Name           goaliyun.String
	Key            goaliyun.String
	ShowType       goaliyun.String
	DisplayUnit    goaliyun.String
	PropertyValues DescribeProductPropertyValueList
}

type DescribeProductPropertyValue struct {
	Value       goaliyun.String
	DisplayName goaliyun.String
	Type        goaliyun.String
	Min         goaliyun.String
	Max         goaliyun.String
	Step        goaliyun.String
	Remark      goaliyun.String
}

type DescribeProductProductExtra struct {
	Key    goaliyun.String
	Values goaliyun.String
	Label  goaliyun.String
	Order  goaliyun.Integer
	Type   goaliyun.String
}

type DescribeProductShopInfo struct {
	Id         goaliyun.Integer
	Name       goaliyun.String
	Emails     goaliyun.String
	WangWangs  DescribeProductWangWangList
	Telephones DescribeProductTelephoneList
}

type DescribeProductWangWang struct {
	UserName goaliyun.String
	Remark   goaliyun.String
}

type DescribeProductProductSkuList []DescribeProductProductSku

func (list *DescribeProductProductSkuList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductProductSku)
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

type DescribeProductProductExtraList []DescribeProductProductExtra

func (list *DescribeProductProductExtraList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductProductExtra)
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

type DescribeProductOrderPeriodList []DescribeProductOrderPeriod

func (list *DescribeProductOrderPeriodList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductOrderPeriod)
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

type DescribeProductModuleList []DescribeProductModule

func (list *DescribeProductModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductModule)
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

type DescribeProductPropertyList []DescribeProductProperty

func (list *DescribeProductPropertyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductProperty)
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

type DescribeProductPropertyValueList []DescribeProductPropertyValue

func (list *DescribeProductPropertyValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductPropertyValue)
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

type DescribeProductWangWangList []DescribeProductWangWang

func (list *DescribeProductWangWangList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductWangWang)
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

type DescribeProductTelephoneList []goaliyun.String

func (list *DescribeProductTelephoneList) UnmarshalJSON(data []byte) error {
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
