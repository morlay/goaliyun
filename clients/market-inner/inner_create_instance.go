package market_inner

import (
	"github.com/morlay/goaliyun"
)

type InnerCreateInstanceRequest struct {
	PayLoad     string `position:"Query" name:"PayLoad"`
	ClientToken string `position:"Query" name:"ClientToken"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *InnerCreateInstanceRequest) Invoke(client goaliyun.Client) (*InnerCreateInstanceResponse, error) {
	resp := &InnerCreateInstanceResponse{}
	err := client.Request("market-inner", "InnerCreateInstance", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InnerCreateInstanceResponse struct {
	RequestId  goaliyun.String
	InstanceId goaliyun.String
}
