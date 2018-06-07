package teslamaxcompute

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryResourceInventoryRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *QueryResourceInventoryRequest) Invoke(client goaliyun.Client) (*QueryResourceInventoryResponse, error) {
	resp := &QueryResourceInventoryResponse{}
	err := client.Request("teslamaxcompute", "QueryResourceInventory", "2018-01-04", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryResourceInventoryResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Data      QueryResourceInventoryData
}

type QueryResourceInventoryData struct {
	LastUpdate goaliyun.String
	Clusters   QueryResourceInventoryClusterList
}

type QueryResourceInventoryCluster struct {
	Status              goaliyun.String
	Cluster             goaliyun.String
	MachineRoom         goaliyun.String
	Region              goaliyun.String
	ResourceParameters  QueryResourceInventoryResourceParameterList
	ResourceInventories QueryResourceInventoryResourceInventoryList
}

type QueryResourceInventoryResourceParameter struct {
	ParaName  goaliyun.String
	ParaValue goaliyun.String
}

type QueryResourceInventoryResourceInventory struct {
	Total        goaliyun.Integer
	Available    goaliyun.Integer
	Used         goaliyun.Integer
	ResourceType goaliyun.String
}

type QueryResourceInventoryClusterList []QueryResourceInventoryCluster

func (list *QueryResourceInventoryClusterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryResourceInventoryCluster)
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

type QueryResourceInventoryResourceParameterList []QueryResourceInventoryResourceParameter

func (list *QueryResourceInventoryResourceParameterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryResourceInventoryResourceParameter)
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

type QueryResourceInventoryResourceInventoryList []QueryResourceInventoryResourceInventory

func (list *QueryResourceInventoryResourceInventoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryResourceInventoryResourceInventory)
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
