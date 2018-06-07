package market_inner

import (
	"github.com/morlay/goaliyun"
)

type InnerReleaseInstanceRequest struct {
	PayLoad     string `position:"Query" name:"PayLoad"`
	ClientToken string `position:"Query" name:"ClientToken"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *InnerReleaseInstanceRequest) Invoke(client goaliyun.Client) (*InnerReleaseInstanceResponse, error) {
	resp := &InnerReleaseInstanceResponse{}
	err := client.Request("market-inner", "InnerReleaseInstance", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type InnerReleaseInstanceResponse struct {
	RequestId goaliyun.String
}
