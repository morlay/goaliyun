package domain

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryTaskDetailHistoryRequest struct {
	TaskStatus         int64  `position:"Query" name:"TaskStatus"`
	UserClientIp       string `position:"Query" name:"UserClientIp"`
	TaskNo             string `position:"Query" name:"TaskNo"`
	DomainName         string `position:"Query" name:"DomainName"`
	PageSize           int64  `position:"Query" name:"PageSize"`
	TaskDetailNoCursor string `position:"Query" name:"TaskDetailNoCursor"`
	Lang               string `position:"Query" name:"Lang"`
	DomainNameCursor   string `position:"Query" name:"DomainNameCursor"`
	RegionId           string `position:"Query" name:"RegionId"`
}

func (req *QueryTaskDetailHistoryRequest) Invoke(client goaliyun.Client) (*QueryTaskDetailHistoryResponse, error) {
	resp := &QueryTaskDetailHistoryResponse{}
	err := client.Request("domain", "QueryTaskDetailHistory", "2018-01-29", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryTaskDetailHistoryResponse struct {
	RequestId         goaliyun.String
	PageSize          goaliyun.Integer
	Objects           QueryTaskDetailHistoryTaskDetailHistoryList
	CurrentPageCursor QueryTaskDetailHistoryCurrentPageCursor
	NextPageCursor    QueryTaskDetailHistoryNextPageCursor
	PrePageCursor     QueryTaskDetailHistoryPrePageCursor
}

type QueryTaskDetailHistoryTaskDetailHistory struct {
	TaskNo              goaliyun.String
	TaskDetailNo        goaliyun.String
	TaskType            goaliyun.String
	InstanceId          goaliyun.String
	DomainName          goaliyun.String
	TaskStatus          goaliyun.String
	UpdateTime          goaliyun.String
	CreateTime          goaliyun.String
	TryCount            goaliyun.Integer
	ErrorMsg            goaliyun.String
	TaskStatusCode      goaliyun.Integer
	TaskTypeDescription goaliyun.String
}

type QueryTaskDetailHistoryCurrentPageCursor struct {
	TaskNo              goaliyun.String
	TaskDetailNo        goaliyun.String
	TaskType            goaliyun.String
	InstanceId          goaliyun.String
	DomainName          goaliyun.String
	TaskStatus          goaliyun.String
	UpdateTime          goaliyun.String
	CreateTime          goaliyun.String
	TryCount            goaliyun.Integer
	ErrorMsg            goaliyun.String
	TaskStatusCode      goaliyun.Integer
	TaskTypeDescription goaliyun.String
}

type QueryTaskDetailHistoryNextPageCursor struct {
	TaskNo              goaliyun.String
	TaskDetailNo        goaliyun.String
	TaskType            goaliyun.String
	InstanceId          goaliyun.String
	DomainName          goaliyun.String
	TaskStatus          goaliyun.String
	UpdateTime          goaliyun.String
	CreateTime          goaliyun.String
	TryCount            goaliyun.Integer
	ErrorMsg            goaliyun.String
	TaskStatusCode      goaliyun.Integer
	TaskTypeDescription goaliyun.String
}

type QueryTaskDetailHistoryPrePageCursor struct {
	TaskNo              goaliyun.String
	TaskDetailNo        goaliyun.String
	TaskType            goaliyun.String
	InstanceId          goaliyun.String
	DomainName          goaliyun.String
	TaskStatus          goaliyun.String
	UpdateTime          goaliyun.String
	CreateTime          goaliyun.String
	TryCount            goaliyun.Integer
	ErrorMsg            goaliyun.String
	TaskStatusCode      goaliyun.Integer
	TaskTypeDescription goaliyun.String
}

type QueryTaskDetailHistoryTaskDetailHistoryList []QueryTaskDetailHistoryTaskDetailHistory

func (list *QueryTaskDetailHistoryTaskDetailHistoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTaskDetailHistoryTaskDetailHistory)
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
