package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListConvertOfficeFormatTasksRequest struct {
	MaxKeys  int64  `position:"Query" name:"MaxKeys"`
	Marker   string `position:"Query" name:"Marker"`
	Project  string `position:"Query" name:"Project"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListConvertOfficeFormatTasksRequest) Invoke(client goaliyun.Client) (*ListConvertOfficeFormatTasksResponse, error) {
	resp := &ListConvertOfficeFormatTasksResponse{}
	err := client.Request("imm", "ListConvertOfficeFormatTasks", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListConvertOfficeFormatTasksResponse struct {
	RequestId  goaliyun.String
	NextMarker goaliyun.String
	Tasks      ListConvertOfficeFormatTasksTasksItemList
}

type ListConvertOfficeFormatTasksTasksItem struct {
	TaskId          goaliyun.String
	Status          goaliyun.String
	Percent         goaliyun.Integer
	PageCount       goaliyun.Integer
	SrcUri          goaliyun.String
	TgtType         goaliyun.String
	TgtUri          goaliyun.String
	ImageSpec       goaliyun.String
	NotifyTopicName goaliyun.String
	NotifyEndpoint  goaliyun.String
	ExternalID      goaliyun.String
	CreateTime      goaliyun.String
	FinishTime      goaliyun.String
}

type ListConvertOfficeFormatTasksTasksItemList []ListConvertOfficeFormatTasksTasksItem

func (list *ListConvertOfficeFormatTasksTasksItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListConvertOfficeFormatTasksTasksItem)
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
