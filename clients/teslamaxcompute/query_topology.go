package teslamaxcompute

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type QueryTopologyRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *QueryTopologyRequest) Invoke(client goaliyun.Client) (*QueryTopologyResponse, error) {
	resp := &QueryTopologyResponse{}
	err := client.Request("teslamaxcompute", "QueryTopology", "2018-01-04", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type QueryTopologyResponse struct {
	Code      goaliyun.Integer
	Message   goaliyun.String
	RequestId goaliyun.String
	Result    QueryTopologyResultItemList
}

type QueryTopologyResultItem struct {
	LastUpdate goaliyun.String
	Regions    QueryTopologyRegionItemList
}

type QueryTopologyRegionItem struct {
	Region       goaliyun.String
	RegionEnName goaliyun.String
	RegionCnName goaliyun.String
	Clusters     QueryTopologyClusterItemList
}

type QueryTopologyClusterItem struct {
	Cluster      goaliyun.String
	ProductLine  goaliyun.String
	ProductClass goaliyun.String
	NetCode      goaliyun.String
	Business     goaliyun.String
	MachineRoom  goaliyun.String
	NetArch      goaliyun.String
}

type QueryTopologyResultItemList []QueryTopologyResultItem

func (list *QueryTopologyResultItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTopologyResultItem)
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

type QueryTopologyRegionItemList []QueryTopologyRegionItem

func (list *QueryTopologyRegionItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTopologyRegionItem)
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

type QueryTopologyClusterItemList []QueryTopologyClusterItem

func (list *QueryTopologyClusterItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTopologyClusterItem)
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
