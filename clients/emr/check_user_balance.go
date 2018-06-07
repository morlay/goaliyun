package emr

import (
	"github.com/morlay/goaliyun"
)

type CheckUserBalanceRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *CheckUserBalanceRequest) Invoke(client goaliyun.Client) (*CheckUserBalanceResponse, error) {
	resp := &CheckUserBalanceResponse{}
	err := client.Request("emr", "CheckUserBalance", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CheckUserBalanceResponse struct {
	RequestId goaliyun.String
	Balance   goaliyun.String
	Enough    goaliyun.String
}
