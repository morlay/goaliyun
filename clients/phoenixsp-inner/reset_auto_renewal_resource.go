package phoenixsp_inner

import (
	"github.com/morlay/goaliyun"
)

type ResetAutoRenewalResourceRequest struct {
	InstanceNames         string `position:"Query" name:"InstanceNames"`
	Caller                string `position:"Query" name:"Caller"`
	RenewalDuration       int64  `position:"Query" name:"RenewalDuration"`
	AutoRenewal           string `position:"Query" name:"AutoRenewal"`
	RenewalStatus         string `position:"Query" name:"RenewalStatus"`
	SaleCycle             string `position:"Query" name:"SaleCycle"`
	AutoRenewalOffsetDays string `position:"Query" name:"AutoRenewalOffsetDays"`
	AliUID                int64  `position:"Query" name:"AliUID"`
	RenewalCycUnit        int64  `position:"Query" name:"RenewalCycUnit"`
	Bid                   string `position:"Query" name:"Bid"`
	ResourceType          string `position:"Query" name:"ResourceType"`
	Operator              string `position:"Query" name:"Operator"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *ResetAutoRenewalResourceRequest) Invoke(client goaliyun.Client) (*ResetAutoRenewalResourceResponse, error) {
	resp := &ResetAutoRenewalResourceResponse{}
	err := client.Request("phoenixsp-inner", "ResetAutoRenewalResource", "2016-08-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ResetAutoRenewalResourceResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
}
