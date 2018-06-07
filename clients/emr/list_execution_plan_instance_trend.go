package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListExecutionPlanInstanceTrendRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListExecutionPlanInstanceTrendRequest) Invoke(client goaliyun.Client) (*ListExecutionPlanInstanceTrendResponse, error) {
	resp := &ListExecutionPlanInstanceTrendResponse{}
	err := client.Request("emr", "ListExecutionPlanInstanceTrend", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListExecutionPlanInstanceTrendResponse struct {
	RequestId goaliyun.String
	Trends    ListExecutionPlanInstanceTrendTrendList
}

type ListExecutionPlanInstanceTrendTrend struct {
	Day    goaliyun.String
	Count  goaliyun.Integer
	Status goaliyun.String
}

type ListExecutionPlanInstanceTrendTrendList []ListExecutionPlanInstanceTrendTrend

func (list *ListExecutionPlanInstanceTrendTrendList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListExecutionPlanInstanceTrendTrend)
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
