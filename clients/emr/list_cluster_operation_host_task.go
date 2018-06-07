package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterOperationHostTaskRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	OperationId     string `position:"Query" name:"OperationId"`
	HostId          string `position:"Query" name:"HostId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	Status          string `position:"Query" name:"Status"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterOperationHostTaskRequest) Invoke(client goaliyun.Client) (*ListClusterOperationHostTaskResponse, error) {
	resp := &ListClusterOperationHostTaskResponse{}
	err := client.Request("emr", "ListClusterOperationHostTask", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterOperationHostTaskResponse struct {
	RequestId                    goaliyun.String
	TotalCount                   goaliyun.Integer
	PageNumber                   goaliyun.Integer
	PageSize                     goaliyun.Integer
	ClusterOperationHostTaskList ListClusterOperationHostTaskClusterOperationHostTaskList
}

type ListClusterOperationHostTaskClusterOperationHostTask struct {
	TaskId     goaliyun.String
	TaskName   goaliyun.String
	Status     goaliyun.String
	Percentage goaliyun.String
}

type ListClusterOperationHostTaskClusterOperationHostTaskList []ListClusterOperationHostTaskClusterOperationHostTask

func (list *ListClusterOperationHostTaskClusterOperationHostTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterOperationHostTaskClusterOperationHostTask)
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
