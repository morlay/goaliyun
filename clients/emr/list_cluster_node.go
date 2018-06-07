package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterNodeRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterNodeRequest) Invoke(client goaliyun.Client) (*ListClusterNodeResponse, error) {
	resp := &ListClusterNodeResponse{}
	err := client.Request("emr", "ListClusterNode", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterNodeResponse struct {
	RequestId       goaliyun.String
	ClusterNodeList ListClusterNodeClusterNodeList
}

type ListClusterNodeClusterNode struct {
	NodeId goaliyun.String
	NodeIp goaliyun.String
	Status goaliyun.String
}

type ListClusterNodeClusterNodeList []ListClusterNodeClusterNode

func (list *ListClusterNodeClusterNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterNodeClusterNode)
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
