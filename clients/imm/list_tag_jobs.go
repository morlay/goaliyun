package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListTagJobsRequest struct {
	Condition string `position:"Query" name:"Condition"`
	MaxKeys   int64  `position:"Query" name:"MaxKeys"`
	Marker    string `position:"Query" name:"Marker"`
	Project   string `position:"Query" name:"Project"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *ListTagJobsRequest) Invoke(client goaliyun.Client) (*ListTagJobsResponse, error) {
	resp := &ListTagJobsResponse{}
	err := client.Request("imm", "ListTagJobs", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListTagJobsResponse struct {
	RequestId  goaliyun.String
	NextMarker goaliyun.String
	Jobs       ListTagJobsJobsItemList
}

type ListTagJobsJobsItem struct {
	JobId      goaliyun.String
	SetId      goaliyun.String
	SrcUri     goaliyun.String
	Status     goaliyun.String
	Percent    goaliyun.Integer
	CreateTime goaliyun.String
	FinishTime goaliyun.String
}

type ListTagJobsJobsItemList []ListTagJobsJobsItem

func (list *ListTagJobsJobsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTagJobsJobsItem)
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
