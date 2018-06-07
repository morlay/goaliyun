package phoenixsp_inner

import (
	"github.com/morlay/goaliyun"
)

type RiskControlRequest struct {
	Reason        string `position:"Query" name:"Reason"`
	RiskCondition string `position:"Query" name:"RiskCondition"`
	Aliuid        int64  `position:"Query" name:"Aliuid"`
	Operation     string `position:"Query" name:"Operation"`
	Token         string `position:"Query" name:"Token"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *RiskControlRequest) Invoke(client goaliyun.Client) (*RiskControlResponse, error) {
	resp := &RiskControlResponse{}
	err := client.Request("phoenixsp-inner", "RiskControl", "2016-08-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RiskControlResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
	Success   bool
}
