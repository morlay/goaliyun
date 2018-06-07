package phoenixsp_inner

import (
	"github.com/morlay/goaliyun"
)

type FindResRenewalStartTimeAsBatchRequest struct {
	Caller         string `position:"Query" name:"Caller"`
	InstanceIdList string `position:"Query" name:"InstanceIdList"`
	Aliuid         int64  `position:"Query" name:"Aliuid"`
	Bid            string `position:"Query" name:"Bid"`
	ResourceType   string `position:"Query" name:"ResourceType"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *FindResRenewalStartTimeAsBatchRequest) Invoke(client goaliyun.Client) (*FindResRenewalStartTimeAsBatchResponse, error) {
	resp := &FindResRenewalStartTimeAsBatchResponse{}
	err := client.Request("phoenixsp-inner", "FindResRenewalStartTimeAsBatch", "2016-08-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type FindResRenewalStartTimeAsBatchResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
	Success   bool
	Data      goaliyun.Integer
}
