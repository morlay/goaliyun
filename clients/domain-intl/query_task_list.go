package domain_intl

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryTaskListRequest struct {
	BeginCreateTime int64  `position:"Query" name:"BeginCreateTime"`
	EndCreateTime   int64  `position:"Query" name:"EndCreateTime"`
	UserClientIp    string `position:"Query" name:"UserClientIp"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	Lang            string `position:"Query" name:"Lang"`
	PageNum         int64  `position:"Query" name:"PageNum"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *QueryTaskListRequest) Invoke(client goaliyun.Client) (*QueryTaskListResponse, error) {
	resp := &QueryTaskListResponse{}
	err := client.Request("domain-intl", "QueryTaskList", "2017-12-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryTaskListResponse struct {
	RequestId      goaliyun.String
	TotalItemNum   goaliyun.Integer
	CurrentPageNum goaliyun.Integer
	TotalPageNum   goaliyun.Integer
	PageSize       goaliyun.Integer
	PrePage        bool
	NextPage       bool
	Data           QueryTaskListTaskInfoList
}

type QueryTaskListTaskInfo struct {
	TaskType            goaliyun.String
	TaskNum             goaliyun.Integer
	TaskStatus          goaliyun.String
	CreateTime          goaliyun.String
	Clientip            goaliyun.String
	TaskNo              goaliyun.String
	TaskTypeDescription goaliyun.String
	TaskStatusCode      goaliyun.Integer
}

type QueryTaskListTaskInfoList []QueryTaskListTaskInfo

func (list *QueryTaskListTaskInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTaskListTaskInfo)
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
