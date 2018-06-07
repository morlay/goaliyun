package cms

import (
	"github.com/morlay/goaliyun"
)

type QuerySystemEventHistogramRequest struct {
	QueryJson string `position:"Query" name:"QueryJson"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *QuerySystemEventHistogramRequest) Invoke(client goaliyun.Client) (*QuerySystemEventHistogramResponse, error) {
	resp := &QuerySystemEventHistogramResponse{}
	err := client.Request("cms", "QuerySystemEventHistogram", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QuerySystemEventHistogramResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
	RequestId goaliyun.String
	Success   goaliyun.String
}
