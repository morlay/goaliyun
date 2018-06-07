package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type DoPhysicalDeleteResourceRequest struct {
	Country        string `position:"Query" name:"Country"`
	Hid            int64  `position:"Query" name:"Hid"`
	Success        string `position:"Query" name:"Success"`
	Interrupt      string `position:"Query" name:"Interrupt"`
	GmtWakeup      string `position:"Query" name:"GmtWakeup"`
	Pk             string `position:"Query" name:"Pk"`
	Bid            string `position:"Query" name:"Bid"`
	TaskExtraData  string `position:"Query" name:"TaskExtraData"`
	TaskIdentifier string `position:"Query" name:"TaskIdentifier"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DoPhysicalDeleteResourceRequest) Invoke(client goaliyun.Client) (*DoPhysicalDeleteResourceResponse, error) {
	resp := &DoPhysicalDeleteResourceResponse{}
	err := client.Request("qualitycheck", "DoPhysicalDeleteResource", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DoPhysicalDeleteResourceResponse struct {
	Interrupt       goaliyun.String
	Invoker         goaliyun.Integer
	Pk              goaliyun.String
	Bid             goaliyun.String
	Hid             goaliyun.Integer
	Country         goaliyun.String
	TaskIdentifier  goaliyun.String
	TaskIdentifier1 goaliyun.String
	GmtWakeup       goaliyun.String
	Success         bool
	Message         goaliyun.String
	Level           goaliyun.Integer
	Prompt          goaliyun.String
}
