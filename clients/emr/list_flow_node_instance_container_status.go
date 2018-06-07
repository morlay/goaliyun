package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type ListFlowNodeInstanceContainerStatusRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	NodeInstanceId  string `position:"Query" name:"NodeInstanceId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *ListFlowNodeInstanceContainerStatusRequest) Invoke(client goaliyun.Client) (*ListFlowNodeInstanceContainerStatusResponse, error) {
	resp := &ListFlowNodeInstanceContainerStatusResponse{}
	err := client.Request("emr", "ListFlowNodeInstanceContainerStatus", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ListFlowNodeInstanceContainerStatusResponse struct {
	RequestId           goaliyun.String
	ContainerStatusList ListFlowNodeInstanceContainerStatusContainerStatusList
}

type ListFlowNodeInstanceContainerStatusContainerStatus struct {
	ApplicationId goaliyun.String
	ContainerId   goaliyun.String
	HostName      goaliyun.String
	Status        goaliyun.String
}

type ListFlowNodeInstanceContainerStatusContainerStatusList []ListFlowNodeInstanceContainerStatusContainerStatus

func (list *ListFlowNodeInstanceContainerStatusContainerStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFlowNodeInstanceContainerStatusContainerStatus)
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
