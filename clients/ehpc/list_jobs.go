package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListJobsRequest struct {
	Owner      string `position:"Query" name:"Owner"`
	PageSize   int64  `position:"Query" name:"PageSize"`
	ClusterId  string `position:"Query" name:"ClusterId"`
	State      string `position:"Query" name:"State"`
	Rerunable  string `position:"Query" name:"Rerunable"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListJobsRequest) Invoke(client goaliyun.Client) (*ListJobsResponse, error) {
	resp := &ListJobsResponse{}
	err := client.Request("ehpc", "ListJobs", "2018-04-12", req.RegionId, "", "").Do(req, resp)
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
	Id             goaliyun.String
	Name           goaliyun.String
	Owner          goaliyun.String
	Priority       goaliyun.Integer
	State          goaliyun.String
	SubmitTime     goaliyun.String
	StartTime      goaliyun.String
	LastModifyTime goaliyun.String
	Stdout         goaliyun.String
	Stderr         goaliyun.String
	ShellPath      goaliyun.String
	Comment        goaliyun.String
	ArrayRequest   goaliyun.String
	Resources      ListJobsResources
}

type ListJobsResources struct {
	Nodes goaliyun.Integer
	Cores goaliyun.Integer
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
