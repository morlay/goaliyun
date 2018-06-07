package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListJobExecutionInstanceTrendRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListJobExecutionInstanceTrendRequest) Invoke(client goaliyun.Client) (*ListJobExecutionInstanceTrendResponse, error) {
	resp := &ListJobExecutionInstanceTrendResponse{}
	err := client.Request("emr", "ListJobExecutionInstanceTrend", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListJobExecutionInstanceTrendResponse struct {
	RequestId goaliyun.String
	Trends    ListJobExecutionInstanceTrendTrendList
}

type ListJobExecutionInstanceTrendTrend struct {
	Day    goaliyun.String
	Count  goaliyun.Integer
	Status goaliyun.String
}

type ListJobExecutionInstanceTrendTrendList []ListJobExecutionInstanceTrendTrend

func (list *ListJobExecutionInstanceTrendTrendList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobExecutionInstanceTrendTrend)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
