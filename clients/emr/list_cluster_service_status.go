package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterServiceStatusRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterServiceStatusRequest) Invoke(client goaliyun.Client) (*ListClusterServiceStatusResponse, error) {
	resp := &ListClusterServiceStatusResponse{}
	err := client.Request("emr", "ListClusterServiceStatus", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterServiceStatusResponse struct {
	RequestId                goaliyun.String
	ClusterServiceStatusList ListClusterServiceStatusClusterServiceStatusList
}

type ListClusterServiceStatusClusterServiceStatus struct {
	NodeIp goaliyun.String
	Status goaliyun.String
}

type ListClusterServiceStatusClusterServiceStatusList []ListClusterServiceStatusClusterServiceStatus

func (list *ListClusterServiceStatusClusterServiceStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceStatusClusterServiceStatus)
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
