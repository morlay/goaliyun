package bss

import (
	"github.com/morlay/goaliyun"
)

type DescribeCashDetailRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DescribeCashDetailRequest) Invoke(client goaliyun.Client) (*DescribeCashDetailResponse, error) {
	resp := &DescribeCashDetailResponse{}
	err := client.Request("bss", "DescribeCashDetail", "2014-07-14", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCashDetailResponse struct {
	RequestId            goaliyun.String
	BalanceAmount        goaliyun.String
	AmountOwed           goaliyun.String
	EnableThresholdAlert goaliyun.String
	MiniAlertThreshold   goaliyun.Integer
	FrozenAmount         goaliyun.String
	CreditCardAmount     goaliyun.String
	RemmitanceAmount     goaliyun.String
	CreditLimit          goaliyun.String
	AvailableCredit      goaliyun.String
}
