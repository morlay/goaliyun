package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListNotesRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListNotesRequest) Invoke(client goaliyun.Client) (*ListNotesResponse, error) {
	resp := &ListNotesResponse{}
	err := client.Request("emr", "ListNotes", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListNotesResponse struct {
	RequestId goaliyun.String
	Notes     ListNotesNoteInfoList
}

type ListNotesNoteInfo struct {
	Id   goaliyun.String
	Name goaliyun.String
	Type goaliyun.String
}

type ListNotesNoteInfoList []ListNotesNoteInfo

func (list *ListNotesNoteInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListNotesNoteInfo)
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
