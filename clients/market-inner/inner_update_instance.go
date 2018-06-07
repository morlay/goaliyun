package market_inner

import (
	"github.com/morlay/goaliyun"
)

type InnerUpdateInstanceRequest struct {
	PayLoad     string `position:"Query" name:"PayLoad"`
	ClientToken string `position:"Query" name:"ClientToken"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *InnerUpdateInstanceRequest) Invoke(client goaliyun.Client) (*InnerUpdateInstanceResponse, error) {
	resp := &InnerUpdateInstanceResponse{}
	err := client.Request("market-inner", "InnerUpdateInstance", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InnerUpdateInstanceResponse struct {
	RequestId goaliyun.String
}
