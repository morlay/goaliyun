package cdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRefreshTasksRequest struct {
	ObjectPath      string `position:"Query" name:"ObjectPath"`
	DomainName      string `position:"Query" name:"DomainName"`
	EndTime         string `position:"Query" name:"EndTime"`
	StartTime       string `position:"Query" name:"StartTime"`
	OwnerId         int64  `position:"Query" name:"OwnerId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ObjectType      string `position:"Query" name:"ObjectType"`
	TaskId          string `position:"Query" name:"TaskId"`
	Status          string `position:"Query" name:"Status"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeRefreshTasksRequest) Invoke(client goaliyun.Client) (*DescribeRefreshTasksResponse, error) {
	resp := &DescribeRefreshTasksResponse{}
	err := client.Request("cdn", "DescribeRefreshTasks", "2014-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRefreshTasksResponse struct {
	RequestId  goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalCount goaliyun.Integer
	Tasks      DescribeRefreshTasksCDNTaskList
}

type DescribeRefreshTasksCDNTask struct {
	TaskId       goaliyun.String
	ObjectPath   goaliyun.String
	Process      goaliyun.String
	Status       goaliyun.String
	CreationTime goaliyun.String
	Description  goaliyun.String
	ObjectType   goaliyun.String
}

type DescribeRefreshTasksCDNTaskList []DescribeRefreshTasksCDNTask

func (list *DescribeRefreshTasksCDNTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRefreshTasksCDNTask)
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
