package domain_intl

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryTaskDetailListRequest struct {
	TaskStatus   int64  `position:"Query" name:"TaskStatus"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	TaskNo       string `position:"Query" name:"TaskNo"`
	DomainName   string `position:"Query" name:"DomainName"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	Lang         string `position:"Query" name:"Lang"`
	PageNum      int64  `position:"Query" name:"PageNum"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *QueryTaskDetailListRequest) Invoke(client goaliyun.Client) (*QueryTaskDetailListResponse, error) {
	resp := &QueryTaskDetailListResponse{}
	err := client.Request("domain-intl", "QueryTaskDetailList", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryTaskDetailListResponse struct {
	RequestId      goaliyun.String
	TotalItemNum   goaliyun.Integer
	CurrentPageNum goaliyun.Integer
	TotalPageNum   goaliyun.Integer
	PageSize       goaliyun.Integer
	PrePage        bool
	NextPage       bool
	Data           QueryTaskDetailListTaskDetailList
}

type QueryTaskDetailListTaskDetail struct {
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
	TaskResult          goaliyun.String
	TaskTypeDescription goaliyun.String
}

type QueryTaskDetailListTaskDetailList []QueryTaskDetailListTaskDetail

func (list *QueryTaskDetailListTaskDetailList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTaskDetailListTaskDetail)
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
