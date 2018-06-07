package cms

import (
	"github.com/morlay/goaliyun"
)

type QueryMetricLastRequest struct {
	Cursor          string `position:"Query" name:"Cursor"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Period          string `position:"Query" name:"Period"`
	Metric          string `position:"Query" name:"Metric"`
	Length          string `position:"Query" name:"Length"`
	Project         string `position:"Query" name:"Project"`
	EndTime         string `position:"Query" name:"EndTime"`
	Express         string `position:"Query" name:"Express"`
	StartTime       string `position:"Query" name:"StartTime"`
	Page            string `position:"Query" name:"Page"`
	Dimensions      string `position:"Query" name:"Dimensions"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *QueryMetricLastRequest) Invoke(client goaliyun.Client) (*QueryMetricLastResponse, error) {
	resp := &QueryMetricLastResponse{}
	err := client.Request("cms", "QueryMetricLast", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryMetricLastResponse struct {
	Code       goaliyun.String
	Message    goaliyun.String
	RequestId  goaliyun.String
	Cursor     goaliyun.String
	Datapoints goaliyun.String
	Period     goaliyun.String
}
