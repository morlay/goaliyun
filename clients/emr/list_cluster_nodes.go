package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListClusterNodesRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListClusterNodesRequest) Invoke(client goaliyun.Client) (*ListClusterNodesResponse, error) {
	resp := &ListClusterNodesResponse{}
	err := client.Request("emr", "ListClusterNodes", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListClusterNodesResponse struct {
	RequestId    goaliyun.String
	ClusterNodes ListClusterNodesClusterNodeList
}

type ListClusterNodesClusterNode struct {
	NodeId goaliyun.String
	NodeIp goaliyun.String
	Status goaliyun.String
}

type ListClusterNodesClusterNodeList []ListClusterNodesClusterNode

func (list *ListClusterNodesClusterNodeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterNodesClusterNode)
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
