package domain_intl

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryTaskInfoHistoryRequest struct {
	BeginCreateTime  int64  `position:"Query" name:"BeginCreateTime"`
	EndCreateTime    int64  `position:"Query" name:"EndCreateTime"`
	TaskNoCursor     string `position:"Query" name:"TaskNoCursor"`
	UserClientIp     string `position:"Query" name:"UserClientIp"`
	PageSize         int64  `position:"Query" name:"PageSize"`
	Lang             string `position:"Query" name:"Lang"`
	CreateTimeCursor int64  `position:"Query" name:"CreateTimeCursor"`
	RegionId         string `position:"Query" name:"RegionId"`
}

func (req *QueryTaskInfoHistoryRequest) Invoke(client goaliyun.Client) (*QueryTaskInfoHistoryResponse, error) {
	resp := &QueryTaskInfoHistoryResponse{}
	err := client.Request("domain-intl", "QueryTaskInfoHistory", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryTaskInfoHistoryResponse struct {
	RequestId         goaliyun.String
	PageSize          goaliyun.Integer
	Objects           QueryTaskInfoHistoryTaskInfoHistoryList
	CurrentPageCursor QueryTaskInfoHistoryCurrentPageCursor
	NextPageCursor    QueryTaskInfoHistoryNextPageCursor
	PrePageCursor     QueryTaskInfoHistoryPrePageCursor
}

type QueryTaskInfoHistoryTaskInfoHistory struct {
	TaskType            goaliyun.String
	TaskNum             goaliyun.Integer
	TaskStatus          goaliyun.String
	CreateTime          goaliyun.String
	Clientip            goaliyun.String
	TaskNo              goaliyun.String
	CreateTimeLong      goaliyun.Integer
	TaskStatusCode      goaliyun.Integer
	TaskTypeDescription goaliyun.String
}

type QueryTaskInfoHistoryCurrentPageCursor struct {
	TaskType            goaliyun.String
	TaskNum             goaliyun.Integer
	TaskStatus          goaliyun.String
	CreateTime          goaliyun.String
	Clientip            goaliyun.String
	TaskNo              goaliyun.String
	CreateTimeLong      goaliyun.Integer
	TaskStatusCode      goaliyun.Integer
	TaskTypeDescription goaliyun.String
}

type QueryTaskInfoHistoryNextPageCursor struct {
	TaskType            goaliyun.String
	TaskNum             goaliyun.Integer
	TaskStatus          goaliyun.String
	CreateTime          goaliyun.String
	Clientip            goaliyun.String
	TaskNo              goaliyun.String
	CreateTimeLong      goaliyun.Integer
	TaskStatusCode      goaliyun.Integer
	TaskTypeDescription goaliyun.String
}

type QueryTaskInfoHistoryPrePageCursor struct {
	TaskType            goaliyun.String
	TaskNum             goaliyun.Integer
	TaskStatus          goaliyun.String
	CreateTime          goaliyun.String
	Clientip            goaliyun.String
	TaskNo              goaliyun.String
	CreateTimeLong      goaliyun.Integer
	TaskStatusCode      goaliyun.Integer
	TaskTypeDescription goaliyun.String
}

type QueryTaskInfoHistoryTaskInfoHistoryList []QueryTaskInfoHistoryTaskInfoHistory

func (list *QueryTaskInfoHistoryTaskInfoHistoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTaskInfoHistoryTaskInfoHistory)
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
