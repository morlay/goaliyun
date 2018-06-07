package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListFaceSetsRequest struct {
	Marker   string `position:"Query" name:"Marker"`
	Project  string `position:"Query" name:"Project"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *ListFaceSetsRequest) Invoke(client goaliyun.Client) (*ListFaceSetsResponse, error) {
	resp := &ListFaceSetsResponse{}
	err := client.Request("imm", "ListFaceSets", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListFaceSetsResponse struct {
	RequestId  goaliyun.String
	NextMarker goaliyun.String
	Sets       ListFaceSetsSetsItemList
}

type ListFaceSetsSetsItem struct {
	SetId      goaliyun.String
	Status     goaliyun.String
	Photos     goaliyun.Integer
	CreateTime goaliyun.String
	ModifyTime goaliyun.String
}

type ListFaceSetsSetsItemList []ListFaceSetsSetsItem

func (list *ListFaceSetsSetsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFaceSetsSetsItem)
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
