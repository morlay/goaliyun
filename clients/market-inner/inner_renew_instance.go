package market_inner

import (
	"github.com/morlay/goaliyun"
)

type InnerRenewInstanceRequest struct {
	PayLoad     string `position:"Query" name:"PayLoad"`
	ClientToken string `position:"Query" name:"ClientToken"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *InnerRenewInstanceRequest) Invoke(client goaliyun.Client) (*InnerRenewInstanceResponse, error) {
	resp := &InnerRenewInstanceResponse{}
	err := client.Request("market-inner", "InnerRenewInstance", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InnerRenewInstanceResponse struct {
	RequestId goaliyun.String
}
