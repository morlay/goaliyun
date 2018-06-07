package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetMyGroupsRequest struct {
	SelectContactGroups string `position:"Query" name:"SelectContactGroups"`
	InstanceId          string `position:"Query" name:"InstanceId"`
	GroupId             int64  `position:"Query" name:"GroupId"`
	Type                string `position:"Query" name:"Type"`
	GroupName           string `position:"Query" name:"GroupName"`
	BindUrl             string `position:"Query" name:"BindUrl"`
	RegionId            string `position:"Query" name:"RegionId"`
}

func (req *GetMyGroupsRequest) Invoke(client goaliyun.Client) (*GetMyGroupsResponse, error) {
	resp := &GetMyGroupsResponse{}
	err := client.Request("cms", "GetMyGroups", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetMyGroupsResponse struct {
	RequestId    goaliyun.String
	Success      bool
	ErrorCode    goaliyun.Integer
	ErrorMessage goaliyun.String
	Group        GetMyGroupsGroup
}

type GetMyGroupsGroup struct {
	GroupId       goaliyun.Integer
	GroupName     goaliyun.String
	ServiceId     goaliyun.Integer
	BindUrl       goaliyun.String
	Type          goaliyun.String
	ContactGroups GetMyGroupsContactGroupList
}

type GetMyGroupsContactGroup struct {
	Name goaliyun.String
}

type GetMyGroupsContactGroupList []GetMyGroupsContactGroup

func (list *GetMyGroupsContactGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetMyGroupsContactGroup)
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
