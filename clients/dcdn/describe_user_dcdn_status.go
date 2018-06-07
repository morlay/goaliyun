package dcdn

import (
	"github.com/morlay/goaliyun"
)

type DescribeUserDcdnStatusRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeUserDcdnStatusRequest) Invoke(client goaliyun.Client) (*DescribeUserDcdnStatusResponse, error) {
	resp := &DescribeUserDcdnStatusResponse{}
	err := client.Request("dcdn", "DescribeUserDcdnStatus", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeUserDcdnStatusResponse struct {
	RequestId     goaliyun.String
	Enabled       bool
	OnService     bool
	InDebt        bool
	InDebtOverdue bool
}
