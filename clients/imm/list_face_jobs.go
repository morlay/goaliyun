package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListFaceJobsRequest struct {
	Condition string `position:"Query" name:"Condition"`
	MaxKeys   int64  `position:"Query" name:"MaxKeys"`
	Marker    string `position:"Query" name:"Marker"`
	Project   string `position:"Query" name:"Project"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ListFaceJobsRequest) Invoke(client goaliyun.Client) (*ListFaceJobsResponse, error) {
	resp := &ListFaceJobsResponse{}
	err := client.Request("imm", "ListFaceJobs", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListFaceJobsResponse struct {
	RequestId  goaliyun.String
	NextMarker goaliyun.String
	Jobs       ListFaceJobsJobsItemList
}

type ListFaceJobsJobsItem struct {
	JobId      goaliyun.String
	SetId      goaliyun.String
	SrcUri     goaliyun.String
	Status     goaliyun.String
	Percent    goaliyun.Integer
	CreateTime goaliyun.String
	FinishTime goaliyun.String
}

type ListFaceJobsJobsItemList []ListFaceJobsJobsItem

func (list *ListFaceJobsJobsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFaceJobsJobsItem)
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
