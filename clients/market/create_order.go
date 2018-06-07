package market

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CreateOrderRequest struct {
	OrderSouce  string `position:"Query" name:"OrderSouce"`
	Commodity   string `position:"Query" name:"Commodity"`
	ClientToken string `position:"Query" name:"ClientToken"`
	OwnerId     string `position:"Query" name:"OwnerId"`
	PaymentType string `position:"Query" name:"PaymentType"`
	OrderType   string `position:"Query" name:"OrderType"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *CreateOrderRequest) Invoke(client goaliyun.Client) (*CreateOrderResponse, error) {
	resp := &CreateOrderResponse{}
	err := client.Request("market", "CreateOrder", "2015-11-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateOrderResponse struct {
	RequestId   goaliyun.String
	OrderId     goaliyun.String
	InstanceIds CreateOrderInstanceIdList
}

type CreateOrderInstanceIdList []goaliyun.String

func (list *CreateOrderInstanceIdList) UnmarshalJSON(data []byte) error {
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
