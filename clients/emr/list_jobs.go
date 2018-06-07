package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListJobsRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	QueryString     string `position:"Query" name:"QueryString"`
	IsDesc          string `position:"Query" name:"IsDesc"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	QueryType       string `position:"Query" name:"QueryType"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListJobsRequest) Invoke(client goaliyun.Client) (*ListJobsResponse, error) {
	resp := &ListJobsResponse{}
	err := client.Request("emr", "ListJobs", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListJobsResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Jobs       ListJobsJobInfoList
}

type ListJobsJobInfo struct {
	Id            goaliyun.String
	Name          goaliyun.String
	Type          goaliyun.String
	RunParameter  goaliyun.String
	FailAct       goaliyun.String
	MaxRetry      goaliyun.Integer
	RetryInterval goaliyun.Integer
}

type ListJobsJobInfoList []ListJobsJobInfo

func (list *ListJobsJobInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobsJobInfo)
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
