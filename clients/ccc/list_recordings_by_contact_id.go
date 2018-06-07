package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListRecordingsByContactIdRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	ContactId  string `position:"Query" name:"ContactId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ListRecordingsByContactIdRequest) Invoke(client goaliyun.Client) (*ListRecordingsByContactIdResponse, error) {
	resp := &ListRecordingsByContactIdResponse{}
	err := client.Request("ccc", "ListRecordingsByContactId", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListRecordingsByContactIdResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	Recordings     ListRecordingsByContactIdRecordingList
}

type ListRecordingsByContactIdRecording struct {
	ContactId       goaliyun.String
	ContactType     goaliyun.String
	AgentId         goaliyun.String
	AgentName       goaliyun.String
	CallingNumber   goaliyun.String
	CalledNumber    goaliyun.String
	StartTime       goaliyun.Integer
	Duration        goaliyun.Integer
	FileName        goaliyun.String
	FilePath        goaliyun.String
	FileDescription goaliyun.String
	Channel         goaliyun.String
	InstanceId      goaliyun.String
}

type ListRecordingsByContactIdRecordingList []ListRecordingsByContactIdRecording

func (list *ListRecordingsByContactIdRecordingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRecordingsByContactIdRecording)
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
