package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryMediaWorkflowExecutionListRequest struct {
	RunIds               string `position:"Query" name:"RunIds"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *QueryMediaWorkflowExecutionListRequest) Invoke(client goaliyun.Client) (*QueryMediaWorkflowExecutionListResponse, error) {
	resp := &QueryMediaWorkflowExecutionListResponse{}
	err := client.Request("mts", "QueryMediaWorkflowExecutionList", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryMediaWorkflowExecutionListResponse struct {
	RequestId                  goaliyun.String
	MediaWorkflowExecutionList QueryMediaWorkflowExecutionListMediaWorkflowExecutionList
	NonExistRunIds             QueryMediaWorkflowExecutionListNonExistRunIdList
}

type QueryMediaWorkflowExecutionListMediaWorkflowExecution struct {
	RunId           goaliyun.String
	MediaWorkflowId goaliyun.String
	Name            goaliyun.String
	State           goaliyun.String
	MediaId         goaliyun.String
	CreationTime    goaliyun.String
	ActivityList    QueryMediaWorkflowExecutionListActivityList
	Input           QueryMediaWorkflowExecutionListInput
}

type QueryMediaWorkflowExecutionListActivity struct {
	Name             goaliyun.String
	Type             goaliyun.String
	JobId            goaliyun.String
	State            goaliyun.String
	Code             goaliyun.String
	Message          goaliyun.String
	StartTime        goaliyun.String
	EndTime          goaliyun.String
	MNSMessageResult QueryMediaWorkflowExecutionListMNSMessageResult
}

type QueryMediaWorkflowExecutionListMNSMessageResult struct {
	MessageId    goaliyun.String
	ErrorMessage goaliyun.String
	ErrorCode    goaliyun.String
}

type QueryMediaWorkflowExecutionListInput struct {
	UserData  goaliyun.String
	InputFile QueryMediaWorkflowExecutionListInputFile
}

type QueryMediaWorkflowExecutionListInputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type QueryMediaWorkflowExecutionListMediaWorkflowExecutionList []QueryMediaWorkflowExecutionListMediaWorkflowExecution

func (list *QueryMediaWorkflowExecutionListMediaWorkflowExecutionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaWorkflowExecutionListMediaWorkflowExecution)
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

type QueryMediaWorkflowExecutionListNonExistRunIdList []goaliyun.String

func (list *QueryMediaWorkflowExecutionListNonExistRunIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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

type QueryMediaWorkflowExecutionListActivityList []QueryMediaWorkflowExecutionListActivity

func (list *QueryMediaWorkflowExecutionListActivityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaWorkflowExecutionListActivity)
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
