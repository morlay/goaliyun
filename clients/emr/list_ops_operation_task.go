package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListOpsOperationTaskRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	OperationId     int64  `position:"Query" name:"OperationId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListOpsOperationTaskRequest) Invoke(client goaliyun.Client) (*ListOpsOperationTaskResponse, error) {
	resp := &ListOpsOperationTaskResponse{}
	err := client.Request("emr", "ListOpsOperationTask", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListOpsOperationTaskResponse struct {
	RequestId  goaliyun.String
	Total      goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TaskList   ListOpsOperationTaskTaskListItemList
}

type ListOpsOperationTaskTaskListItem struct {
	Id                goaliyun.Integer
	OpsOperationId    goaliyun.Integer
	TaskId            goaliyun.Integer
	Status            goaliyun.String
	RegionId          goaliyun.String
	UserId            goaliyun.String
	ClusterId         goaliyun.Integer
	ExternalClusterId goaliyun.String
	HostId            goaliyun.Integer
	StartTime         goaliyun.String
	EndTime           goaliyun.String
	CommandName       goaliyun.String
	HostName          goaliyun.String
}

type ListOpsOperationTaskTaskListItemList []ListOpsOperationTaskTaskListItem

func (list *ListOpsOperationTaskTaskListItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListOpsOperationTaskTaskListItem)
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
