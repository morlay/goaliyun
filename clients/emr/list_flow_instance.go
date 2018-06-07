package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListFlowInstanceRequest struct {
	ResourceOwnerId int64                           `position:"Query" name:"ResourceOwnerId"`
	StatusLists     *ListFlowInstanceStatusListList `position:"Query" type:"Repeated" name:"StatusList"`
	PageSize        int64                           `position:"Query" name:"PageSize"`
	FlowName        string                          `position:"Query" name:"FlowName"`
	Id              string                          `position:"Query" name:"Id"`
	FlowId          string                          `position:"Query" name:"FlowId"`
	ProjectId       string                          `position:"Query" name:"ProjectId"`
	PageNumber      int64                           `position:"Query" name:"PageNumber"`
	RegionId        string                          `position:"Query" name:"RegionId"`
}

func (req *ListFlowInstanceRequest) Invoke(client goaliyun.Client) (*ListFlowInstanceResponse, error) {
	resp := &ListFlowInstanceResponse{}
	err := client.Request("emr", "ListFlowInstance", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListFlowInstanceResponse struct {
	RequestId     goaliyun.String
	PageNumber    goaliyun.Integer
	PageSize      goaliyun.Integer
	Total         goaliyun.Integer
	FlowInstances ListFlowInstanceFlowInstanceList
}

type ListFlowInstanceFlowInstance struct {
	Id          goaliyun.String
	GmtCreate   goaliyun.Integer
	GmtModified goaliyun.Integer
	FlowId      goaliyun.String
	FlowName    goaliyun.String
	ProjectId   goaliyun.String
	Status      goaliyun.String
	ClusterId   goaliyun.String
	StartTime   goaliyun.Integer
	EndTime     goaliyun.Integer
}

type ListFlowInstanceStatusListList []string

func (list *ListFlowInstanceStatusListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type ListFlowInstanceFlowInstanceList []ListFlowInstanceFlowInstance

func (list *ListFlowInstanceFlowInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFlowInstanceFlowInstance)
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
