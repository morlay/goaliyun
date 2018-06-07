package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListJobExecutionPlanHierarchyRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	CurrentId       int64  `position:"Query" name:"CurrentId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListJobExecutionPlanHierarchyRequest) Invoke(client goaliyun.Client) (*ListJobExecutionPlanHierarchyResponse, error) {
	resp := &ListJobExecutionPlanHierarchyResponse{}
	err := client.Request("emr", "ListJobExecutionPlanHierarchy", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListJobExecutionPlanHierarchyResponse struct {
	RequestId      goaliyun.String
	Success        goaliyun.String
	ErrCode        goaliyun.String
	ErrMsg         goaliyun.String
	TotalCount     goaliyun.Integer
	PageSize       goaliyun.Integer
	PageNumber     goaliyun.Integer
	HierarchyInfos ListJobExecutionPlanHierarchyHierarchyInfoList
}

type ListJobExecutionPlanHierarchyHierarchyInfo struct {
	Id                   goaliyun.Integer
	NodeBizType          goaliyun.String
	NodeType             goaliyun.String
	RelateId             goaliyun.String
	Name                 goaliyun.String
	ParentId             goaliyun.String
	ResourceOwnerId      goaliyun.String
	UtcCreateTimestamp   goaliyun.Integer
	UtcModifiedTimestamp goaliyun.Integer
	NodeStatus           goaliyun.Integer
	ExecutionPlan        ListJobExecutionPlanHierarchyExecutionPlan
	Job                  ListJobExecutionPlanHierarchyJob
}

type ListJobExecutionPlanHierarchyExecutionPlan struct {
	BizId                goaliyun.String
	Name                 goaliyun.String
	Status               goaliyun.Integer
	LastExeStatus        goaliyun.Integer
	IsCreateCluster      bool
	IsInterruptWhenError bool
	IsCycle              bool
	ScheduleCycle        goaliyun.String
	RegionId             goaliyun.String
	CronExpr             goaliyun.String
	UtcCreateTimestamp   goaliyun.Integer
	UtcModifiedTimestamp goaliyun.Integer
	Version              goaliyun.String
	ClusterTemplateId    goaliyun.String
}

type ListJobExecutionPlanHierarchyJob struct {
	BizId         goaliyun.String
	Name          goaliyun.String
	JobFailAct    goaliyun.Integer
	JobType       goaliyun.Integer
	EnvParam      goaliyun.String
	MaxRetry      goaliyun.Integer
	RetryInterval goaliyun.Integer
}

type ListJobExecutionPlanHierarchyHierarchyInfoList []ListJobExecutionPlanHierarchyHierarchyInfo

func (list *ListJobExecutionPlanHierarchyHierarchyInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobExecutionPlanHierarchyHierarchyInfo)
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
