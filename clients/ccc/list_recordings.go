package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListRecordingsRequest struct {
	AgentId     string `position:"Query" name:"AgentId"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	Criteria    string `position:"Query" name:"Criteria"`
	PhoneNumber string `position:"Query" name:"PhoneNumber"`
	PageSize    int64  `position:"Query" name:"PageSize"`
	StopTime    int64  `position:"Query" name:"StopTime"`
	StartTime   int64  `position:"Query" name:"StartTime"`
	PageNumber  int64  `position:"Query" name:"PageNumber"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *ListRecordingsRequest) Invoke(client goaliyun.Client) (*ListRecordingsResponse, error) {
	resp := &ListRecordingsResponse{}
	err := client.Request("ccc", "ListRecordings", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListRecordingsResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	Recordings     ListRecordingsRecordings
}

type ListRecordingsRecordings struct {
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	List       ListRecordingsRecordingList
}

type ListRecordingsRecording struct {
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

type ListRecordingsRecordingList []ListRecordingsRecording

func (list *ListRecordingsRecordingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRecordingsRecording)
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
