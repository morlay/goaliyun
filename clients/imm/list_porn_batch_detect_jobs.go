package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListPornBatchDetectJobsRequest struct {
	MaxKeys  int64  `position:"Query" name:"MaxKeys"`
	Marker   string `position:"Query" name:"Marker"`
	Project  string `position:"Query" name:"Project"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListPornBatchDetectJobsRequest) Invoke(client goaliyun.Client) (*ListPornBatchDetectJobsResponse, error) {
	resp := &ListPornBatchDetectJobsResponse{}
	err := client.Request("imm", "ListPornBatchDetectJobs", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListPornBatchDetectJobsResponse struct {
	RequestId  goaliyun.String
	NextMarker goaliyun.String
	Jobs       ListPornBatchDetectJobsJobsItemList
}

type ListPornBatchDetectJobsJobsItem struct {
	JobId           goaliyun.String
	SrcUri          goaliyun.String
	Status          goaliyun.String
	TgtUri          goaliyun.String
	NotifyTopicName goaliyun.Integer
	NotifyEndpoint  goaliyun.String
	ExternalID      goaliyun.String
	CreateTime      goaliyun.String
	FinishTime      goaliyun.String
	Percent         goaliyun.Integer
}

type ListPornBatchDetectJobsJobsItemList []ListPornBatchDetectJobsJobsItem

func (list *ListPornBatchDetectJobsJobsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListPornBatchDetectJobsJobsItem)
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
