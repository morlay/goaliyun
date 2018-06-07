package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListFlowJobHistoryRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	Id              string `position:"Query" name:"Id"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListFlowJobHistoryRequest) Invoke(client goaliyun.Client) (*ListFlowJobHistoryResponse, error) {
	resp := &ListFlowJobHistoryResponse{}
	err := client.Request("emr", "ListFlowJobHistory", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListFlowJobHistoryResponse struct {
	RequestId     goaliyun.String
	PageNumber    goaliyun.Integer
	PageSize      goaliyun.Integer
	Total         goaliyun.Integer
	NodeInstances ListFlowJobHistoryNodeInstanceList
}

type ListFlowJobHistoryNodeInstance struct {
	Id             goaliyun.String
	GmtCreate      goaliyun.Integer
	GmtModified    goaliyun.Integer
	Type           goaliyun.String
	Status         goaliyun.String
	JobId          goaliyun.String
	JobName        goaliyun.String
	JobType        goaliyun.String
	JobParams      goaliyun.String
	FailAct        goaliyun.String
	MaxRetry       goaliyun.Integer
	RetryInterval  goaliyun.Integer
	NodeName       goaliyun.String
	ClusterId      goaliyun.String
	HostName       goaliyun.String
	ProjectId      goaliyun.String
	StartTime      goaliyun.Integer
	EndTime        goaliyun.Integer
	Pending        bool
	Retries        goaliyun.Integer
	ExternalId     goaliyun.String
	ExternalStatus goaliyun.String
	ExternalInfo   goaliyun.String
	ParamConf      goaliyun.String
	EnvConf        goaliyun.String
	RunConf        goaliyun.String
}

type ListFlowJobHistoryNodeInstanceList []ListFlowJobHistoryNodeInstance

func (list *ListFlowJobHistoryNodeInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFlowJobHistoryNodeInstance)
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
