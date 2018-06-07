package dcdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDcdnRefreshTasksRequest struct {
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

func (req *DescribeDcdnRefreshTasksRequest) Invoke(client goaliyun.Client) (*DescribeDcdnRefreshTasksResponse, error) {
	resp := &DescribeDcdnRefreshTasksResponse{}
	err := client.Request("dcdn", "DescribeDcdnRefreshTasks", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDcdnRefreshTasksResponse struct {
	RequestId  goaliyun.String
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	TotalCount goaliyun.Integer
	Tasks      DescribeDcdnRefreshTasksTaskList
}

type DescribeDcdnRefreshTasksTask struct {
	TaskId       goaliyun.String
	ObjectPath   goaliyun.String
	Process      goaliyun.String
	Status       goaliyun.String
	CreationTime goaliyun.String
	Description  goaliyun.String
	ObjectType   goaliyun.String
}

type DescribeDcdnRefreshTasksTaskList []DescribeDcdnRefreshTasksTask

func (list *DescribeDcdnRefreshTasksTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnRefreshTasksTask)
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
