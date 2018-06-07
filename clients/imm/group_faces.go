package imm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GroupFacesRequest struct {
	Project   string `position:"Query" name:"Project"`
	SetId     string `position:"Query" name:"SetId"`
	Operation string `position:"Query" name:"Operation"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GroupFacesRequest) Invoke(client goaliyun.Client) (*GroupFacesResponse, error) {
	resp := &GroupFacesResponse{}
	err := client.Request("imm", "GroupFaces", "2017-09-06", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GroupFacesResponse struct {
	RequestId goaliyun.String
	SetId     goaliyun.String
	HasMore   goaliyun.Integer
	Groups    GroupFacesGroupsItemList
}

type GroupFacesGroupsItem struct {
	FaceId        goaliyun.String
	GroupId       goaliyun.String
	UnGroupReason goaliyun.String
}

type GroupFacesGroupsItemList []GroupFacesGroupsItem

func (list *GroupFacesGroupsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GroupFacesGroupsItem)
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
