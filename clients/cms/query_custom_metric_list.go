package cms

import (
	"github.com/morlay/goaliyun"
)

type QueryCustomMetricListRequest struct {
	Size       string `position:"Query" name:"Size"`
	GroupId    string `position:"Query" name:"GroupId"`
	Page       string `position:"Query" name:"Page"`
	MetricName string `position:"Query" name:"MetricName"`
	Dimension  string `position:"Query" name:"Dimension"`
	Md5        string `position:"Query" name:"Md.5"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *QueryCustomMetricListRequest) Invoke(client goaliyun.Client) (*QueryCustomMetricListResponse, error) {
	resp := &QueryCustomMetricListResponse{}
	err := client.Request("cms", "QueryCustomMetricList", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryCustomMetricListResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Result    goaliyun.String
}
