package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListContactGroupRequest struct {
	PageSize   int64  `position:"Query" name:"PageSize"`
	PageNumber int64  `position:"Query" name:"PageNumber"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListContactGroupRequest) Invoke(client goaliyun.Client) (*ListContactGroupResponse, error) {
	resp := &ListContactGroupResponse{}
	err := client.Request("cms", "ListContactGroup", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListContactGroupResponse struct {
	Success       bool
	Code          goaliyun.String
	Message       goaliyun.String
	NextToken     goaliyun.Integer
	Total         goaliyun.Integer
	RequestId     goaliyun.String
	ContactGroups ListContactGroupContactGroupList
}

type ListContactGroupContactGroupList []goaliyun.String

func (list *ListContactGroupContactGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
