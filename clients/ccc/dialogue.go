package ccc

import (
	"github.com/morlay/goaliyun"
)

type DialogueRequest struct {
	CallId        string `position:"Query" name:"CallId"`
	CallingNumber string `position:"Query" name:"CallingNumber"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	CalledNumber  string `position:"Query" name:"CalledNumber"`
	ActionParams  string `position:"Query" name:"ActionParams"`
	CallType      string `position:"Query" name:"CallType"`
	ScenarioId    string `position:"Query" name:"ScenarioId"`
	TaskId        string `position:"Query" name:"TaskId"`
	Utterance     string `position:"Query" name:"Utterance"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DialogueRequest) Invoke(client goaliyun.Client) (*DialogueResponse, error) {
	resp := &DialogueResponse{}
	err := client.Request("ccc", "Dialogue", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DialogueResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	Feedback       DialogueFeedback
}

type DialogueFeedback struct {
	Content      goaliyun.String
	Action       goaliyun.String
	ActionParams goaliyun.String
}
