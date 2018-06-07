package ccc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetConversationListRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	TaskId     string `position:"Query" name:"TaskId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *GetConversationListRequest) Invoke(client goaliyun.Client) (*GetConversationListResponse, error) {
	resp := &GetConversationListResponse{}
	err := client.Request("ccc", "GetConversationList", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetConversationListResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.String
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	Conversations  GetConversationListConversationDetailList
}

type GetConversationListConversationDetail struct {
	Timestamp goaliyun.Integer
	Speaker   goaliyun.String
	Script    goaliyun.String
	Summary   GetConversationListSummaryItemList
}

type GetConversationListSummaryItem struct {
	Category    goaliyun.String
	SummaryName goaliyun.String
	Content     goaliyun.String
}

type GetConversationListConversationDetailList []GetConversationListConversationDetail

func (list *GetConversationListConversationDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetConversationListConversationDetail)
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

type GetConversationListSummaryItemList []GetConversationListSummaryItem

func (list *GetConversationListSummaryItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetConversationListSummaryItem)
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
