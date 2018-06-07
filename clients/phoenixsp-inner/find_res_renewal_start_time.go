package phoenixsp_inner

import (
	"github.com/morlay/goaliyun"
)

type FindResRenewalStartTimeRequest struct {
	Caller         string `position:"Query" name:"Caller"`
	InstanceIdList string `position:"Query" name:"InstanceIdList"`
	Aliuid         int64  `position:"Query" name:"Aliuid"`
	Bid            string `position:"Query" name:"Bid"`
	ResourceType   string `position:"Query" name:"ResourceType"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *FindResRenewalStartTimeRequest) Invoke(client goaliyun.Client) (*FindResRenewalStartTimeResponse, error) {
	resp := &FindResRenewalStartTimeResponse{}
	err := client.Request("phoenixsp-inner", "FindResRenewalStartTime", "2016-08-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FindResRenewalStartTimeResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
	Success   bool
	Data      goaliyun.Integer
}
