package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListTagSetsRequest struct {
	MaxKeys  int64  `position:"Query" name:"MaxKeys"`
	Marker   string `position:"Query" name:"Marker"`
	Project  string `position:"Query" name:"Project"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListTagSetsRequest) Invoke(client goaliyun.Client) (*ListTagSetsResponse, error) {
	resp := &ListTagSetsResponse{}
	err := client.Request("imm", "ListTagSets", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListTagSetsResponse struct {
	RequestId  goaliyun.String
	NextMarker goaliyun.String
	Sets       ListTagSetsSetsItemList
}

type ListTagSetsSetsItem struct {
	SetId      goaliyun.String
	Status     goaliyun.String
	Photos     goaliyun.Integer
	CreateTime goaliyun.String
	ModifyTime goaliyun.String
}

type ListTagSetsSetsItemList []ListTagSetsSetsItem

func (list *ListTagSetsSetsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTagSetsSetsItem)
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
