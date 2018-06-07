package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterOperationRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	Status          string `position:"Query" name:"Status"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterOperationRequest) Invoke(client goaliyun.Client) (*ListClusterOperationResponse, error) {
	resp := &ListClusterOperationResponse{}
	err := client.Request("emr", "ListClusterOperation", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterOperationResponse struct {
	RequestId            goaliyun.String
	TotalCount           goaliyun.Integer
	PageNumber           goaliyun.Integer
	PageSize             goaliyun.Integer
	ClusterOperationList ListClusterOperationClusterOperationList
}

type ListClusterOperationClusterOperation struct {
	OperationId   goaliyun.String
	OperationName goaliyun.String
	StartTime     goaliyun.String
	Duration      goaliyun.String
	Status        goaliyun.String
	Percentage    goaliyun.String
	Comment       goaliyun.String
}

type ListClusterOperationClusterOperationList []ListClusterOperationClusterOperation

func (list *ListClusterOperationClusterOperationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterOperationClusterOperation)
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
