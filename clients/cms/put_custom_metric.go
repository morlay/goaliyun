package cms

import (
	"github.com/morlay/goaliyun"
)

type PutCustomMetricRequest struct {
	MetricList string `position:"Query" name:"MetricList"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *PutCustomMetricRequest) Invoke(client goaliyun.Client) (*PutCustomMetricResponse, error) {
	resp := &PutCustomMetricResponse{}
	err := client.Request("cms", "PutCustomMetric", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type PutCustomMetricResponse struct {
	RequestId goaliyun.String
	Code      goaliyun.String
	Message   goaliyun.String
	Data      goaliyun.String
}
