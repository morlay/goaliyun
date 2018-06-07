package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterOperationForAdminRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	UserId          string `position:"Query" name:"UserId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	Status          string `position:"Query" name:"Status"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterOperationForAdminRequest) Invoke(client goaliyun.Client) (*ListClusterOperationForAdminResponse, error) {
	resp := &ListClusterOperationForAdminResponse{}
	err := client.Request("emr", "ListClusterOperationForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterOperationForAdminResponse struct {
	RequestId            goaliyun.String
	TotalCount           goaliyun.Integer
	PageNumber           goaliyun.Integer
	PageSize             goaliyun.Integer
	ClusterOperationList ListClusterOperationForAdminClusterOperationList
}

type ListClusterOperationForAdminClusterOperation struct {
	OperationId   goaliyun.String
	OperationName goaliyun.String
	StartTime     goaliyun.String
	Duration      goaliyun.String
	Status        goaliyun.String
	Percentage    goaliyun.String
	Comment       goaliyun.String
}

type ListClusterOperationForAdminClusterOperationList []ListClusterOperationForAdminClusterOperation

func (list *ListClusterOperationForAdminClusterOperationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterOperationForAdminClusterOperation)
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
