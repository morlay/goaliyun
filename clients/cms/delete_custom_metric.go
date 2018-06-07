package cms

import (
	"github.com/morlay/goaliyun"
)

type DeleteCustomMetricRequest struct {
	GroupId    string `position:"Query" name:"GroupId"`
	MetricName string `position:"Query" name:"MetricName"`
	UUID       string `position:"Query" name:"UUID"`
	Md5        string `position:"Query" name:"Md.5"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *DeleteCustomMetricRequest) Invoke(client goaliyun.Client) (*DeleteCustomMetricResponse, error) {
	resp := &DeleteCustomMetricResponse{}
	err := client.Request("cms", "DeleteCustomMetric", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCustomMetricResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Result    goaliyun.String
}
