package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListOpsOperationRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListOpsOperationRequest) Invoke(client goaliyun.Client) (*ListOpsOperationResponse, error) {
	resp := &ListOpsOperationResponse{}
	err := client.Request("emr", "ListOpsOperation", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListOpsOperationResponse struct {
	RequestId        goaliyun.String
	Total            goaliyun.Integer
	PageNumber       goaliyun.Integer
	PageSize         goaliyun.Integer
	OpsOperationList ListOpsOperationOpsOperationListItemList
}

type ListOpsOperationOpsOperationListItem struct {
	Id              goaliyun.Integer
	StartTime       goaliyun.String
	EndTime         goaliyun.String
	OpsCommandId    goaliyun.Integer
	OpsCommandName  goaliyun.String
	Status          goaliyun.String
	TotalTaskNum    goaliyun.Integer
	FailedTaskNum   goaliyun.Integer
	FinishedTaskNum goaliyun.Integer
	ClusterId       goaliyun.String
	RegionId        goaliyun.String
	Params          goaliyun.String
	Remark          goaliyun.String
	RunningTime     goaliyun.Integer
	Category        goaliyun.String
}

type ListOpsOperationOpsOperationListItemList []ListOpsOperationOpsOperationListItem

func (list *ListOpsOperationOpsOperationListItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListOpsOperationOpsOperationListItem)
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
