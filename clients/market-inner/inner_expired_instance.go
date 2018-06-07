package market_inner

import (
	"github.com/morlay/goaliyun"
)

type InnerExpiredInstanceRequest struct {
	PayLoad     string `position:"Query" name:"PayLoad"`
	ClientToken string `position:"Query" name:"ClientToken"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *InnerExpiredInstanceRequest) Invoke(client goaliyun.Client) (*InnerExpiredInstanceResponse, error) {
	resp := &InnerExpiredInstanceResponse{}
	err := client.Request("market-inner", "InnerExpiredInstance", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InnerExpiredInstanceResponse struct {
	RequestId goaliyun.String
}
