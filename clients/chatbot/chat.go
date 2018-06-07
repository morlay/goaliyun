package chatbot

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ChatRequest struct {
	KnowledgeId string `position:"Query" name:"KnowledgeId"`
	SenderId    string `position:"Query" name:"SenderId"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	SenderNick  string `position:"Query" name:"SenderNick"`
	SessionId   string `position:"Query" name:"SessionId"`
	Tag         string `position:"Query" name:"Tag"`
	Utterance   string `position:"Query" name:"Utterance"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *ChatRequest) Invoke(client goaliyun.Client) (*ChatResponse, error) {
	resp := &ChatResponse{}
	err := client.Request("chatbot", "Chat", "2017-10-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ChatResponse struct {
	RequestId goaliyun.String
	SessionId goaliyun.String
	MessageId goaliyun.String
	Tag       goaliyun.String
	Messages  ChatMessageList
}

type ChatMessage struct {
	Type       goaliyun.String
	Recommends ChatRecommendList
	Text       ChatText
	Knowledge  ChatKnowledge
}

type ChatRecommend struct {
	KnowledgeId  goaliyun.String
	Title        goaliyun.String
	AnswerSource goaliyun.String
}

type ChatText struct {
	Content      goaliyun.String
	AnswerSource goaliyun.String
}

type ChatKnowledge struct {
	Id           goaliyun.String
	Title        goaliyun.String
	Summary      goaliyun.String
	Content      goaliyun.String
	AnswerSource goaliyun.String
}

type ChatMessageList []ChatMessage

func (list *ChatMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ChatMessage)
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

type ChatRecommendList []ChatRecommend

func (list *ChatRecommendList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ChatRecommend)
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
