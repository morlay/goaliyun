package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DeleteMyGroupsRequest struct {
	GroupId  int64  `position:"Query" name:"GroupId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DeleteMyGroupsRequest) Invoke(client goaliyun.Client) (*DeleteMyGroupsResponse, error) {
	resp := &DeleteMyGroupsResponse{}
	err := client.Request("cms", "DeleteMyGroups", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteMyGroupsResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
	Group        DeleteMyGroupsGroup
}

type DeleteMyGroupsGroup struct {
	GroupId       goaliyun.Integer
	GroupName     goaliyun.String
	ServiceId     goaliyun.String
	BindUrls      goaliyun.String
	Type          goaliyun.String
	ContactGroups DeleteMyGroupsContactGroupList
}

type DeleteMyGroupsContactGroup struct {
	Name goaliyun.String
}

type DeleteMyGroupsContactGroupList []DeleteMyGroupsContactGroup

func (list *DeleteMyGroupsContactGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteMyGroupsContactGroup)
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
