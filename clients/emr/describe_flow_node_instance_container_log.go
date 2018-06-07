package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeFlowNodeInstanceContainerLogRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Offset          int64  `position:"Query" name:"Offset"`
	AppId           string `position:"Query" name:"AppId"`
	Length          int64  `position:"Query" name:"Length"`
	ContainerId     string `position:"Query" name:"ContainerId"`
	NodeInstanceId  string `position:"Query" name:"NodeInstanceId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeFlowNodeInstanceContainerLogRequest) Invoke(client goaliyun.Client) (*DescribeFlowNodeInstanceContainerLogResponse, error) {
	resp := &DescribeFlowNodeInstanceContainerLogResponse{}
	err := client.Request("emr", "DescribeFlowNodeInstanceContainerLog", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeFlowNodeInstanceContainerLogResponse struct {
	RequestId goaliyun.String
	Logs      DescribeFlowNodeInstanceContainerLogLogList
}

type DescribeFlowNodeInstanceContainerLogLog struct {
	LogEntry DescribeFlowNodeInstanceContainerLogLogEntry
}

type DescribeFlowNodeInstanceContainerLogLogEntry struct {
	Content goaliyun.String
}

type DescribeFlowNodeInstanceContainerLogLogList []DescribeFlowNodeInstanceContainerLogLog

func (list *DescribeFlowNodeInstanceContainerLogLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeFlowNodeInstanceContainerLogLog)
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
