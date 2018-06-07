package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListOfficeConversionTaskRequest struct {
	MaxKeys  int64  `position:"Query" name:"MaxKeys"`
	Marker   string `position:"Query" name:"Marker"`
	Project  string `position:"Query" name:"Project"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListOfficeConversionTaskRequest) Invoke(client goaliyun.Client) (*ListOfficeConversionTaskResponse, error) {
	resp := &ListOfficeConversionTaskResponse{}
	err := client.Request("imm", "ListOfficeConversionTask", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListOfficeConversionTaskResponse struct {
	RequestId  goaliyun.String
	NextMarker goaliyun.String
	Tasks      ListOfficeConversionTaskTasksItemList
}

type ListOfficeConversionTaskTasksItem struct {
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

type ListOfficeConversionTaskTasksItemList []ListOfficeConversionTaskTasksItem

func (list *ListOfficeConversionTaskTasksItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListOfficeConversionTaskTasksItem)
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
