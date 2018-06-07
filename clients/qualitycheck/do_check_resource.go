package qualitycheck

import (
	"github.com/morlay/goaliyun"
)

type DoCheckResourceRequest struct {
	Country        string `position:"Query" name:"Country"`
	Hid            int64  `position:"Query" name:"Hid"`
	Level          int64  `position:"Query" name:"Level"`
	Success        string `position:"Query" name:"Success"`
	Interrupt      string `position:"Query" name:"Interrupt"`
	GmtWakeup      string `position:"Query" name:"GmtWakeup"`
	Pk             string `position:"Query" name:"Pk"`
	Bid            string `position:"Query" name:"Bid"`
	Message        string `position:"Query" name:"Message"`
	Prompt         string `position:"Query" name:"Prompt"`
	TaskExtraData  string `position:"Query" name:"TaskExtraData"`
	TaskIdentifier string `position:"Query" name:"TaskIdentifier"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DoCheckResourceRequest) Invoke(client goaliyun.Client) (*DoCheckResourceResponse, error) {
	resp := &DoCheckResourceResponse{}
	err := client.Request("qualitycheck", "DoCheckResource", "2016-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DoCheckResourceResponse struct {
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
	Url             goaliyun.String
	Prompt          goaliyun.String
}
