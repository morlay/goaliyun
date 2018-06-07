package cms

import (
	"github.com/morlay/goaliyun"
)

type QueryMetricDataRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Period          string `position:"Query" name:"Period"`
	Metric          string `position:"Query" name:"Metric"`
	Length          string `position:"Query" name:"Length"`
	Project         string `position:"Query" name:"Project"`
	EndTime         string `position:"Query" name:"EndTime"`
	Express         string `position:"Query" name:"Express"`
	StartTime       string `position:"Query" name:"StartTime"`
	Dimensions      string `position:"Query" name:"Dimensions"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *QueryMetricDataRequest) Invoke(client goaliyun.Client) (*QueryMetricDataResponse, error) {
	resp := &QueryMetricDataResponse{}
	err := client.Request("cms", "QueryMetricData", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryMetricDataResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	RequestId  goaliyun.String
	Datapoints goaliyun.String
	Period     goaliyun.String
}
