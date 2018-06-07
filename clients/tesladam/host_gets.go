package tesladam

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type HostGetsRequest struct {
	Query     string `position:"Query" name:"Query"`
	EndTime   int64  `position:"Query" name:"EndTime"`
	StartTime int64  `position:"Query" name:"StartTime"`
	QueryType string `position:"Query" name:"QueryType"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *HostGetsRequest) Invoke(client goaliyun.Client) (*HostGetsResponse, error) {
	resp := &HostGetsResponse{}
	err := client.Request("tesladam", "HostGets", "2018-01-18", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type HostGetsResponse struct {
	Status  bool
	Message goaliyun.String
	Data    HostGetsDataItemList
}

type HostGetsDataItem struct {
	Hostname         goaliyun.String
	Ip               goaliyun.String
	AppCode          goaliyun.String
	ClusterCode      goaliyun.String
	SshStatus        goaliyun.Integer
	HeartStatus      goaliyun.Integer
	HealthScoreLast  goaliyun.Integer
	HealthReasonLast goaliyun.String
}

type HostGetsDataItemList []HostGetsDataItem

func (list *HostGetsDataItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]HostGetsDataItem)
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
