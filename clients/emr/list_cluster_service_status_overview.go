package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterServiceStatusOverviewRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterServiceStatusOverviewRequest) Invoke(client goaliyun.Client) (*ListClusterServiceStatusOverviewResponse, error) {
	resp := &ListClusterServiceStatusOverviewResponse{}
	err := client.Request("emr", "ListClusterServiceStatusOverview", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterServiceStatusOverviewResponse struct {
	RequestId                goaliyun.String
	ClusterServiceStatusList ListClusterServiceStatusOverviewClusterServiceStatusOverviewList
}

type ListClusterServiceStatusOverviewClusterServiceStatusOverview struct {
	ClusterId   goaliyun.String
	ServiceName goaliyun.String
	Status      goaliyun.String
}

type ListClusterServiceStatusOverviewClusterServiceStatusOverviewList []ListClusterServiceStatusOverviewClusterServiceStatusOverview

func (list *ListClusterServiceStatusOverviewClusterServiceStatusOverviewList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceStatusOverviewClusterServiceStatusOverview)
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
