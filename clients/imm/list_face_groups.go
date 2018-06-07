package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListFaceGroupsRequest struct {
	MaxKeys  int64  `position:"Query" name:"MaxKeys"`
	Marker   int64  `position:"Query" name:"Marker"`
	Project  string `position:"Query" name:"Project"`
	SetId    string `position:"Query" name:"SetId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListFaceGroupsRequest) Invoke(client goaliyun.Client) (*ListFaceGroupsResponse, error) {
	resp := &ListFaceGroupsResponse{}
	err := client.Request("imm", "ListFaceGroups", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListFaceGroupsResponse struct {
	RequestId  goaliyun.String
	NextMarker goaliyun.Integer
	Groups     ListFaceGroupsGroupsItemList
}

type ListFaceGroupsGroupsItem struct {
	GroupId goaliyun.Integer
	FaceNum goaliyun.Integer
}

type ListFaceGroupsGroupsItemList []ListFaceGroupsGroupsItem

func (list *ListFaceGroupsGroupsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFaceGroupsGroupsItem)
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
