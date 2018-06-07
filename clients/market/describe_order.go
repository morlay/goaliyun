package market

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeOrderRequest struct {
	OrderId  string `position:"Query" name:"OrderId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DescribeOrderRequest) Invoke(client goaliyun.Client) (*DescribeOrderResponse, error) {
	resp := &DescribeOrderResponse{}
	err := client.Request("market", "DescribeOrder", "2015-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeOrderResponse struct {
	OrderId             goaliyun.Integer
	AliUid              goaliyun.Integer
	SupplierCompanyName goaliyun.String
	ProductCode         goaliyun.String
	ProductSkuCode      goaliyun.String
	ProductName         goaliyun.String
	PeriodType          goaliyun.String
	Quantity            goaliyun.Integer
	AccountQuantity     goaliyun.Integer
	OrderType           goaliyun.String
	OrderStatus         goaliyun.String
	PayStatus           goaliyun.String
	PaidOn              goaliyun.Integer
	CreatedOn           goaliyun.Integer
	OriginalPrice       goaliyun.Float
	TotalPrice          goaliyun.Float
	PaymentPrice        goaliyun.Float
	CouponPrice         goaliyun.Float
	SupplierTelephones  DescribeOrderSupplierTelephoneList
	InstanceIds         DescribeOrderInstanceIdList
}

type DescribeOrderSupplierTelephoneList []goaliyun.String

func (list *DescribeOrderSupplierTelephoneList) UnmarshalJSON(data []byte) error {
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

type DescribeOrderInstanceIdList []goaliyun.String

func (list *DescribeOrderInstanceIdList) UnmarshalJSON(data []byte) error {
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
