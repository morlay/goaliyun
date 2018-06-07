package cms

import (
	"github.com/morlay/goaliyun"
)

type QueryMetricTopRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Period          string `position:"Query" name:"Period"`
	Metric          string `position:"Query" name:"Metric"`
	Length          string `position:"Query" name:"Length"`
	Project         string `position:"Query" name:"Project"`
	EndTime         string `position:"Query" name:"EndTime"`
	Orderby         string `position:"Query" name:"Orderby"`
	Express         string `position:"Query" name:"Express"`
	StartTime       string `position:"Query" name:"StartTime"`
	Dimensions      string `position:"Query" name:"Dimensions"`
	OrderDesc       string `position:"Query" name:"OrderDesc"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *QueryMetricTopRequest) Invoke(client goaliyun.Client) (*QueryMetricTopResponse, error) {
	resp := &QueryMetricTopResponse{}
	err := client.Request("cms", "QueryMetricTop", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryMetricTopResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	RequestId  goaliyun.String
	Datapoints goaliyun.String
	Period     goaliyun.String
}
