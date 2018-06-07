package yundun

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DdosFlowGraphRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DdosFlowGraphRequest) Invoke(client goaliyun.Client) (*DdosFlowGraphResponse, error) {
	resp := &DdosFlowGraphResponse{}
	err := client.Request("yundun", "DdosFlowGraph", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DdosFlowGraphResponse struct {
	RequestId   goaliyun.String
	NormalFlows DdosFlowGraphNormalFlowList
	TotalFlows  DdosFlowGraphTotalFlowList
}

type DdosFlowGraphNormalFlow struct {
	Time    goaliyun.Integer
	BitRecv goaliyun.Integer
	BitSend goaliyun.Integer
	PktRecv goaliyun.Integer
	PktSend goaliyun.Integer
}

type DdosFlowGraphTotalFlow struct {
	Time    goaliyun.Integer
	BitRecv goaliyun.Integer
	PktRecv goaliyun.Integer
}

type DdosFlowGraphNormalFlowList []DdosFlowGraphNormalFlow

func (list *DdosFlowGraphNormalFlowList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DdosFlowGraphNormalFlow)
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

type DdosFlowGraphTotalFlowList []DdosFlowGraphTotalFlow

func (list *DdosFlowGraphTotalFlowList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DdosFlowGraphTotalFlow)
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
