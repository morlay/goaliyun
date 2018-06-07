package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListExecutionPlansRequest struct {
	JobId           string                            `position:"Query" name:"JobId"`
	ResourceOwnerId int64                             `position:"Query" name:"ResourceOwnerId"`
	StatusLists     *ListExecutionPlansStatusListList `position:"Query" type:"Repeated" name:"StatusList"`
	PageSize        int64                             `position:"Query" name:"PageSize"`
	QueryString     string                            `position:"Query" name:"QueryString"`
	ClusterId       string                            `position:"Query" name:"ClusterId"`
	IsDesc          string                            `position:"Query" name:"IsDesc"`
	Strategy        string                            `position:"Query" name:"Strategy"`
	PageNumber      int64                             `position:"Query" name:"PageNumber"`
	QueryType       string                            `position:"Query" name:"QueryType"`
	RegionId        string                            `position:"Query" name:"RegionId"`
}

func (req *ListExecutionPlansRequest) Invoke(client goaliyun.Client) (*ListExecutionPlansResponse, error) {
	resp := &ListExecutionPlansResponse{}
	err := client.Request("emr", "ListExecutionPlans", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListExecutionPlansResponse struct {
	RequestId      goaliyun.String
	TotalCount     goaliyun.Integer
	PageNumber     goaliyun.Integer
	PageSize       goaliyun.Integer
	ExecutionPlans ListExecutionPlansExecutionPlanInfoList
}

type ListExecutionPlansExecutionPlanInfo struct {
	Id                    goaliyun.String
	Name                  goaliyun.String
	CreateClusterOnDemand bool
	Stragety              goaliyun.String
	Status                goaliyun.String
	TimeInterval          goaliyun.Integer
	StartTime             goaliyun.Integer
	TimeUnit              goaliyun.String
}

type ListExecutionPlansStatusListList []string

func (list *ListExecutionPlansStatusListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type ListExecutionPlansExecutionPlanInfoList []ListExecutionPlansExecutionPlanInfo

func (list *ListExecutionPlansExecutionPlanInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListExecutionPlansExecutionPlanInfo)
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
