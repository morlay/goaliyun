package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterOperationHostTaskForAdminRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	OperationId     string `position:"Query" name:"OperationId"`
	HostId          string `position:"Query" name:"HostId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	UserId          string `position:"Query" name:"UserId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	Status          string `position:"Query" name:"Status"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterOperationHostTaskForAdminRequest) Invoke(client goaliyun.Client) (*ListClusterOperationHostTaskForAdminResponse, error) {
	resp := &ListClusterOperationHostTaskForAdminResponse{}
	err := client.Request("emr", "ListClusterOperationHostTaskForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterOperationHostTaskForAdminResponse struct {
	RequestId                    goaliyun.String
	TotalCount                   goaliyun.Integer
	PageNumber                   goaliyun.Integer
	PageSize                     goaliyun.Integer
	ClusterOperationHostTaskList ListClusterOperationHostTaskForAdminClusterOperationHostTaskList
}

type ListClusterOperationHostTaskForAdminClusterOperationHostTask struct {
	TaskId     goaliyun.String
	TaskName   goaliyun.String
	Status     goaliyun.String
	Percentage goaliyun.String
}

type ListClusterOperationHostTaskForAdminClusterOperationHostTaskList []ListClusterOperationHostTaskForAdminClusterOperationHostTask

func (list *ListClusterOperationHostTaskForAdminClusterOperationHostTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterOperationHostTaskForAdminClusterOperationHostTask)
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
