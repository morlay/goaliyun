package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterOperationHostRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	OperationId     string `position:"Query" name:"OperationId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	Status          string `position:"Query" name:"Status"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterOperationHostRequest) Invoke(client goaliyun.Client) (*ListClusterOperationHostResponse, error) {
	resp := &ListClusterOperationHostResponse{}
	err := client.Request("emr", "ListClusterOperationHost", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterOperationHostResponse struct {
	RequestId                goaliyun.String
	TotalCount               goaliyun.Integer
	PageNumber               goaliyun.Integer
	PageSize                 goaliyun.Integer
	ClusterOperationHostList ListClusterOperationHostClusterOperationHostList
}

type ListClusterOperationHostClusterOperationHost struct {
	HostId     goaliyun.String
	HostName   goaliyun.String
	Status     goaliyun.String
	Percentage goaliyun.String
}

type ListClusterOperationHostClusterOperationHostList []ListClusterOperationHostClusterOperationHost

func (list *ListClusterOperationHostClusterOperationHostList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterOperationHostClusterOperationHost)
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
