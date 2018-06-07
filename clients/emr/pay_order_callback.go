package emr

import (
	"github.com/morlay/goaliyun"
)

type PayOrderCallbackRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Data            string `position:"Query" name:"Data"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *PayOrderCallbackRequest) Invoke(client goaliyun.Client) (*PayOrderCallbackResponse, error) {
	resp := &PayOrderCallbackResponse{}
	err := client.Request("emr", "PayOrderCallback", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PayOrderCallbackResponse struct {
	RequestId goaliyun.String
	Success   bool
	Code      goaliyun.String
	Message   goaliyun.String
}
