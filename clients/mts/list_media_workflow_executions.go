package mts

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListMediaWorkflowExecutionsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InputFileURL         string `position:"Query" name:"InputFileURL"`
	NextPageToken        string `position:"Query" name:"NextPageToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MaximumPageSize      int64  `position:"Query" name:"MaximumPageSize"`
	MediaWorkflowId      string `position:"Query" name:"MediaWorkflowId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MediaWorkflowName    string `position:"Query" name:"MediaWorkflowName"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ListMediaWorkflowExecutionsRequest) Invoke(client goaliyun.Client) (*ListMediaWorkflowExecutionsResponse, error) {
	resp := &ListMediaWorkflowExecutionsResponse{}
	err := client.Request("mts", "ListMediaWorkflowExecutions", "2014-06-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListMediaWorkflowExecutionsResponse struct {
	RequestId                  goaliyun.String
	NextPageToken              goaliyun.String
	MediaWorkflowExecutionList ListMediaWorkflowExecutionsMediaWorkflowExecutionList
}

type ListMediaWorkflowExecutionsMediaWorkflowExecution struct {
	RunId           goaliyun.String
	MediaWorkflowId goaliyun.String
	Name            goaliyun.String
	State           goaliyun.String
	MediaId         goaliyun.String
	CreationTime    goaliyun.String
	ActivityList    ListMediaWorkflowExecutionsActivityList
	Input           ListMediaWorkflowExecutionsInput
}

type ListMediaWorkflowExecutionsActivity struct {
	Name             goaliyun.String
	Type             goaliyun.String
	JobId            goaliyun.String
	State            goaliyun.String
	Code             goaliyun.String
	Message          goaliyun.String
	StartTime        goaliyun.String
	EndTime          goaliyun.String
	MNSMessageResult ListMediaWorkflowExecutionsMNSMessageResult
}

type ListMediaWorkflowExecutionsMNSMessageResult struct {
	MessageId    goaliyun.String
	ErrorMessage goaliyun.String
	ErrorCode    goaliyun.String
}

type ListMediaWorkflowExecutionsInput struct {
	UserData  goaliyun.String
	InputFile ListMediaWorkflowExecutionsInputFile
}

type ListMediaWorkflowExecutionsInputFile struct {
	Bucket   goaliyun.String
	Location goaliyun.String
	Object   goaliyun.String
}

type ListMediaWorkflowExecutionsMediaWorkflowExecutionList []ListMediaWorkflowExecutionsMediaWorkflowExecution

func (list *ListMediaWorkflowExecutionsMediaWorkflowExecutionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMediaWorkflowExecutionsMediaWorkflowExecution)
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

type ListMediaWorkflowExecutionsActivityList []ListMediaWorkflowExecutionsActivity

func (list *ListMediaWorkflowExecutionsActivityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMediaWorkflowExecutionsActivity)
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
