package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListPhotoProcessTasksRequest struct {
	MaxKeys  int64  `position:"Query" name:"MaxKeys"`
	Marker   string `position:"Query" name:"Marker"`
	Project  string `position:"Query" name:"Project"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListPhotoProcessTasksRequest) Invoke(client goaliyun.Client) (*ListPhotoProcessTasksResponse, error) {
	resp := &ListPhotoProcessTasksResponse{}
	err := client.Request("imm", "ListPhotoProcessTasks", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListPhotoProcessTasksResponse struct {
	RequestId  goaliyun.String
	NextMarker goaliyun.String
	Tasks      ListPhotoProcessTasksTasksItemList
}

type ListPhotoProcessTasksTasksItem struct {
	TaskId          goaliyun.String
	Status          goaliyun.String
	SrcUri          goaliyun.String
	TgtUri          goaliyun.String
	Style           goaliyun.String
	NotifyTopicName goaliyun.String
	NotifyEndpoint  goaliyun.String
	ExternalID      goaliyun.String
	CreateTime      goaliyun.String
	FinishTime      goaliyun.String
	Percent         goaliyun.Integer
}

type ListPhotoProcessTasksTasksItemList []ListPhotoProcessTasksTasksItem

func (list *ListPhotoProcessTasksTasksItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPhotoProcessTasksTasksItem)
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
